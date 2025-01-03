package dailylanggame/game

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type LanguageInfo struct {
	Language string `json:"language"`
	Level    string `json:"level"`
	//Focus string `json:"focus"`
}

type DailyConfig struct {
	ID             string         `json:"id"`
	Difficulty     string         `json:"difficulty"`
	LanguageConfig []LanguageInfo `json:"languageconfig"`
}

var configs = []DailyConfig{
	{
		ID: "1", Difficulty: "Easy",
		LanguageConfig: []LanguageInfo{
			{Language: "English", Level: "4"},
			{Language: "Chinese", Level: "2"},
		},
	},
	{
		ID: "2", Difficulty: "Hard",
		LanguageConfig: []LanguageInfo{
			{Language: "French", Level: "3"},
			{Language: "Spanish", Level: "1"},
		},
	},
}

// get all items
func getConfigs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(configs)
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, config := range configs {
		if config.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(config)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

// new config for new user
func createConfig(w http.ResponseWriter, r *http.Request) {
	var newConfig DailyConfig
	err := json.NewDecoder(r.Body).Decode(&newConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	configs = append(configs, newConfig)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newConfig)
}

// update config
func updateConfig(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedConfig DailyConfig
	err := json.NewDecoder(r.Body).Decode(&updatedConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, config := range configs {
		if config.ID == params["id"] {
			configs[i] = updatedConfig
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedConfig)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func deleteConfig(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, config := range configs {
		if config.ID == params["id"] {
			configs = append(configs[:i], configs[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}
