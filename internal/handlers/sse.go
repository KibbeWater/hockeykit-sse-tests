package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"hockeykit-tester/internal/models"

	"github.com/gorilla/mux"
)

func StreamSSE(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scenarioID := vars["scenario"]

	// Read the scenario file
	data, err := ioutil.ReadFile(filepath.Join("tests/scenarios", scenarioID+".json"))
	if err != nil {
		http.Error(w, "Scenario not found", http.StatusNotFound)
		return
	}

	var scenario models.Scenario
	if err := json.Unmarshal(data, &scenario); err != nil {
		http.Error(w, "Invalid scenario format", http.StatusInternalServerError)
		return
	}

	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Stream each package
	for _, pkg := range scenario.Packages {
		// Wait for the specified delay
		time.Sleep(time.Duration(pkg.Delay) * time.Millisecond)

		// Write the raw data
		_, err := fmt.Fprint(w, pkg.Data)
		if err != nil {
			return
		}

		// If this is the final package, write the SSE event separator
		if pkg.IsFinal {
			_, err := fmt.Fprint(w, "\n\n")
			if err != nil {
				return
			}
		}

		// Flush the response
		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}
	}
}
