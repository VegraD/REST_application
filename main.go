package main

import (
	"Assignment-1/constants"
	"Assignment-1/handler"
	"Assignment-1/uptime"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Assigning port, Default: 8080")
		port = "8080"
	}

	http.HandleFunc(constants.PATH_ON_LAUNCH, handler.DefaultHandler)     //In case of start at root path.
	http.HandleFunc(constants.DEFAULT_PATH, handler.DefaultHandler)       //"Front page"
	http.HandleFunc(constants.UINFO_PATH, handler.UInfoHandler)           //Acquiring uni information
	http.HandleFunc(constants.UNEIGHBOUR_PATH, handler.UNeighbourHandler) //Acquiring neighbour unis
	http.HandleFunc(constants.DIAG_PATH, handler.DiagHandler)             //Diagnostics handler

	log.Println("Server initiating on port " + port + " .")
	uptime.Init()
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
