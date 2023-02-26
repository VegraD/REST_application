package handler

import (
	"fmt"
	"net/http"
)

func UNeighbourHandler(w http.ResponseWriter, r *http.Request) {

	//TODO: MAKE DIS
	w.Header().Set("content-type", "text/html")

	output := "Service does not provide any functionality on root path. Please use the" +
		"following paths: " + "<a href=\"v1\"" + UINFO_PATH + "\">" + UINFO_PATH + "</a> " +
		"to view university information. Or <a href=\"v1\"" + UNEIGHBOUR_PATH + "\">" +
		UNEIGHBOUR_PATH + "</a> to find neighbour universities!"

	_, err := fmt.Fprintf(w, "%v", output)

	if err != nil {
		http.Error(w, "Error while returning output", http.StatusInternalServerError)
	}

}
