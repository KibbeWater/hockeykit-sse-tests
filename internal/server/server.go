package server

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"hockeykit-tester/internal/handlers"
	"hockeykit-tester/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

// validateScenarios checks if all JSON files in tests/scenarios are valid
func validateScenarios() error {
	files, err := os.ReadDir("tests/scenarios")
	if err != nil {
		return fmt.Errorf("failed to read scenarios directory: %v", err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		data, err := os.ReadFile(filepath.Join("tests/scenarios", file.Name()))
		if err != nil {
			return fmt.Errorf("failed to read scenario file %s: %v", file.Name(), err)
		}

		var scenario models.Scenario
		if err := json.Unmarshal(data, &scenario); err != nil {
			return fmt.Errorf("invalid scenario format in %s: %v", file.Name(), err)
		}

		// Validate required fields
		if scenario.ID == "" {
			return fmt.Errorf("missing ID in scenario file: %s", file.Name())
		}
		if scenario.Name == "" {
			return fmt.Errorf("missing Name in scenario file: %s", file.Name())
		}
		if len(scenario.Packages) == 0 {
			return fmt.Errorf("no packages defined in scenario file: %s", file.Name())
		}
	}

	return nil
}

// StartServer initializes the HTTP server and sets up the routes.
func StartServer() error {
	// Validate scenarios before starting the server
	if err := validateScenarios(); err != nil {
		return fmt.Errorf("scenario validation failed: %v", err)
	}

	r := mux.NewRouter()

	// Define routes for handling SSE events
	r.HandleFunc("/tests/scenario/{scenario}", handlers.StreamSSE).Methods("GET")
	r.HandleFunc("/tests/scenario", handlers.GetScenario).Methods("GET")

	// Start the server
	log.Println("All scenarios validated successfully. Starting server on :8080")
	http.Handle("/", r)
	return http.ListenAndServe(":8080", nil)
}
