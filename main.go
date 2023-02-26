package main

import (
	"Assignment-1/handler"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, world!")

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Assigning port, Default: 8080")
		port = "8080"
	}

	http.HandleFunc(handler.DEFAULT_PATH, handler.DefaultHandler)       //"Front page"
	http.HandleFunc(handler.UINFO_PATH, handler.UInfoHandler)           //Acquiring uni information
	http.HandleFunc(handler.UNEIGHBOUR_PATH, handler.UNeighbourHandler) //Acquiring neighbour unis
	http.HandleFunc(handler.DIAG_PATH, handler.DiagHandler)             //Diagnostics handler

	log.Println("Server initiating on port " + port + " .")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
