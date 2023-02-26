package handler

import (
	"encoding/json"
	"net/http"
)

func UInfoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetRequest(w, r)
	default:
		http.Error(w, "REST method '"+r.Method+"' not currently supported. At this moment"+
			"only '"+http.MethodGet+"' are supported.", http.StatusNotImplemented)
		return
	}

}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {

	//TODO: Write struct for university information
	university := University{
		Name:     "NTNU",
		Country:  "Norway",
		Isocode:  "NO",
		Webpages: []string{"http://www.ntnu.no/", "http://www.google.no/"},
		Languages: []string{"'nno': 'Norwegian Nynorsk'",
			"'nob': 'Norwegian Bokmål'",
			"'smi': 'Sami'"},
		Map: "https://www.openstreetmap.org/relation/2978650",
	}

	w.Header().Add("content-type", "application/json")

	err := json.NewEncoder(w).Encode(university)
	if err != nil {
		http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
