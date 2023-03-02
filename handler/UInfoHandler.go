package handler

import (
	"Assignment-1/DB"
	"encoding/json"
	"net/http"
	"strings"
)

func UInfoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetRequest(w, r)
	default:
		http.Error(w, "REST method '"+r.Method+"' not currently supported. At this moment "+
			"only '"+http.MethodGet+"' are supported.", http.StatusNotImplemented)
		return
	}

}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 5 {
		http.Error(w, "Too many arguments, please enter input as such: 'uniinfo/{country}'", http.StatusBadRequest)
		return
		//ERROR, NEED ONE MORE ARGUMENT
	}
	value := parts[4]
	if len(value) == 0 {
		http.Error(w, "Kindly provide a valid country name!", http.StatusBadRequest)
		return
		//ERROR, VALUE IS NULL AND NOT A VALID COUNTRY
	}

	for _, s := range DB.Db {
		if value == s.Country {
			err := json.NewEncoder(w).Encode(s)
			if err != nil {
				http.Error(w, "Error while returning output", http.StatusInternalServerError)
				return
			}
		}
	}

	http.Error(w, "", http.StatusNoContent)
}
