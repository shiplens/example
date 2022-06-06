package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

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

	mux.HandleFunc("/json", handleJSON)

	return mux
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
