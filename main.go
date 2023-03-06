package main

import (
	"Assignment-1/application/handlers"
	"Assignment-1/constants"
	"Assignment-1/uptime"
	"log"
	"net/http"
	"os"
)

/*
Main function of the application.
*/
func main() {

	// Chooses and sets port.
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Assigning port, Default: 8080")
		port = "8080"
	}
	// Set as default-handlers in case of start at root path to avoid 404.
	http.HandleFunc(constants.PATH_ON_LAUNCH, handlers.DefaultHandler)

	// DefaultHandler handles functionality at root path.
	http.HandleFunc(constants.DEFAULT_PATH, handlers.DefaultHandler)

	// UInfoHandler functionality for displaying university information.
	http.HandleFunc(constants.UINFO_PATH, handlers.UInfoHandler)

	// UNeighbourHandler functionality for displaying neighbouring universities
	http.HandleFunc(constants.UNEIGHBOUR_PATH, handlers.UNeighbourHandler)

	// DiagHandler for displaying service diagnostics
	http.HandleFunc(constants.DIAG_PATH, handlers.DiagHandler)

	log.Println("Server initiating on port " + port + " .")

	// Initializes uptime counter
	uptime.Init()

	// Service start.
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
