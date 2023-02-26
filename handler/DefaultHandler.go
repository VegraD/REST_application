package handler

import (
	"fmt"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "text/html")

	output := "Service does not provide any functionality on root path. Please use the" +
		"following paths: " + "<a href=" + UINFO_PATH + "\">" + UINFO_PATH + "</a> " +
		"to view university information. Or <a href=" + UNEIGHBOUR_PATH + "\">" +
		UNEIGHBOUR_PATH + "</a> to find neighbour universities!"

	_, err := fmt.Fprintf(w, "%v", output)

	if err != nil {
		http.Error(w, "Error while returning output", http.StatusInternalServerError)
	}

}
