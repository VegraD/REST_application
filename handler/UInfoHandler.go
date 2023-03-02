package handler

import (
	"Assignment-1/DB"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func UInfoHandler(w http.ResponseWriter, r *http.Request) {

	stdOutput := "Hello! Welcome to page!"

	if r.Method != http.MethodGet {
		http.Error(w, "Method "+r.Method+"is not supported. At this current time, only "+http.MethodGet+
			"are supported", http.StatusNotImplemented)
		return
	} else {
		_, err := fmt.Fprintf(w, stdOutput)
		if err != nil {
			http.Error(w, "Error while returning output", http.StatusInternalServerError)
		}
		handleGetRequestN(w, r)
	}

}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 5 {
		//ERROR, NEED ONE MORE ARGUMENT
	}

	value := parts[4]
	if len(value) == 0 {
		//ERROR, VALUE IS NULL AND NOT A VALID COUNTRY
	}

	for _, s := range DB.Db {
		if value == s.Country {
			err := json.NewEncoder(w).Encode(s)
			if err != nil {
				http.Error(w, "Error while returning output", http.StatusInternalServerError)
			}
		}
	}

	http.Error(w, "", http.StatusNoContent)
}
