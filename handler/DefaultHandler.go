package handler

import (
	"Assignment-1/constants"
	"fmt"
	"net/http"
)

/*
The default handler of the application. Returns recommended actions for users to take.
Parameters:

	w: ResponseWriter (user of application) to write error message to.
	r: A request pointer given by the user
*/
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "text/html")

	output := "Service does not provide any functionality on root path. Please use the " +
		"following paths: " + "<a href=" + constants.UINFO_PATH + ">" + constants.UINFO_PATH + "</a> " +
		"to view university information. Or <a href=" + constants.UNEIGHBOUR_PATH + ">" +
		constants.UNEIGHBOUR_PATH + "</a> to find neighbour universities!"

	_, err := fmt.Fprintf(w, "%v", output)

	if err != nil {
		http.Error(w, "Error while returning output", http.StatusInternalServerError)
	}

}
