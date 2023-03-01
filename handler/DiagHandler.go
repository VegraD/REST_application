package handler

import (
	"Assignment-1/constants"
	"fmt"
	"net/http"
)

func DiagHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "text/html")

	output := "Service does not provide any functionality on root path. Please use the" +
		"following paths: " + "<a href=\"v1\"" + constants.UINFO_PATH + "\">" + constants.UINFO_PATH + "</a> " +
		"to view university information. Or <a href=\"v1\"" + constants.UNEIGHBOUR_PATH + "\">" +
		constants.UNEIGHBOUR_PATH + "</a> to find neighbour universities!"

	_, err := fmt.Fprintf(w, "%v", output)

	if err != nil {
		http.Error(w, "Error while returning output", http.StatusInternalServerError)
	}

}
