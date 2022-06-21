package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

//go:embed templates/*
var content embed.FS

var (
	port = "8080"

	gitSHA    = "development"
	timestamp = "development"
)

func init() {
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
}

func main() {
	readBuildInfo()
	mux := serveMux()
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}

func serveMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", middleware(handleRoot))
	mux.HandleFunc("/json", middleware(handleJSON))

	return mux
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", fmt.Sprintf("example/%s", gitSHA))

		next.ServeHTTP(w, r)
	})
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(content, "templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	host := strings.SplitN(r.Host, ":", 2)[0]

	err = tmpl.Execute(w, struct {
		Host    string
		GitSHA  string
		ISO8601 string
	}{
		Host:    host,
		GitSHA:  gitSHA,
		ISO8601: timestamp,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]map[string]string{
			"git": {
				"sha": gitSHA,
			},
			"time": {
				"iso8601": timestamp,
			},
		},
	)
}

func readBuildInfo() {
	build, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("Could not read build info")
	}
	for _, v := range build.Settings {
		if v.Key == "vcs.revision" {
			gitSHA = v.Value
		}

		if v.Key == "vcs.time" {
			timestamp = v.Value
		}
	}

	log.Printf("Git SHA: %s", gitSHA)
	log.Printf("Timestamp: %s", timestamp)
}
