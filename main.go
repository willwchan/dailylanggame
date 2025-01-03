package dailylanggame

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/willwchan/dailylanggame/game"
	"log"
	"net/http"
)

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/daily", getConfigs).Methods("GET")
	router.HandleFunc("/daily/{id}", getConfig).Methods("GET")
	router.HandleFunc("/daily", createConfig).Methods("POST")
	router.HandleFunc("/daily/{id}", updateConfig).Methods("PUT")
	router.HandleFunc("/daily/{id}", deleteConfig).Methods("DELETE")

	// Start the server
	fmt.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

//handlerfunc is more common than making custom http handlers
