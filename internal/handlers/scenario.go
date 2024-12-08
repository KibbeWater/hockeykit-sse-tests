package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"hockeykit-tester/internal/models"
)

func GetScenario(w http.ResponseWriter, r *http.Request) {
	// List all scenarios from the directory
	scenarios := []PublicScenario{}
	files, err := os.ReadDir("tests/scenarios")
	if err != nil {
		http.Error(w, "Error reading scenarios", http.StatusInternalServerError)
		return
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			data, err := os.ReadFile(filepath.Join("tests/scenarios", file.Name()))
			if err != nil {
				continue
			}

			var scenario models.Scenario
			if err := json.Unmarshal(data, &scenario); err != nil {
				continue
			}
			scenarios = append(scenarios, PublicScenario{
				ID:          scenario.ID,
				Name:        scenario.Name,
				Description: scenario.Description,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scenarios)
}

type PublicScenario struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
