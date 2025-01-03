package main

import (
	"fmt"
	"github.com/gorilla/mux"
	game "github.com/willwchan/dailylanggame/game"
	"log"
	"net/http"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/daily", game.GetConfigs).Methods("GET")
	router.HandleFunc("/daily/{id}", game.GetConfig).Methods("GET")
	router.HandleFunc("/daily", game.CreateConfig).Methods("POST")
	router.HandleFunc("/daily/{id}", game.UpdateConfig).Methods("PUT")
	router.HandleFunc("/daily/{id}", game.DeleteConfig).Methods("DELETE")

	// Start the server
	fmt.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

//handlerfunc is more common than making custom http handlers
