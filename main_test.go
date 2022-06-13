package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	gitSHA = "abc123"
	timestamp = "2022-01-01T00:00:00Z"
}

func TestHandleJSON(t *testing.T) {
	h := serveMux()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/json", nil)
	h.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, fmt.Sprintf("example/%s", gitSHA), w.Header().Get("Server"))
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	expected := map[string]map[string]string{
		"git": {
			"sha": gitSHA,
		},
		"time": {
			"iso8601": timestamp,
		},
	}

	b, err := json.Marshal(expected)
	require.NoError(t, err)
	assert.JSONEq(t, string(b), w.Body.String())
}
