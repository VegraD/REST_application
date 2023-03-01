package handler

import (
	"Assignment-1/structs"
	"encoding/json"
	"net/http"
)

func UNeighbourHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetRequestN(w, r)
	default:
		http.Error(w, "REST method '"+r.Method+"' not currently supported. At this moment"+
			"only '"+http.MethodGet+"' are supported.", http.StatusNotImplemented)
		return
	}

}

func handleGetRequestN(w http.ResponseWriter, r *http.Request) {
	university := structs.UniAndCountry{
		Name:     "NTNU",
		Country:  "Norway",
		Isocode:  "NO",
		Webpages: []string{"http://www.ntnu.no/", "http://www.google.no/"},
		Languages: map[string]string{"nno": "Norwegian Nynorsk",
			"nob": "Norwegian Bokmål",
			"smi": "Sami"},
		Map: "https://www.openstreetmap.org/relation/2978650",
	}

	w.Header().Add("content-type", "application/json")

	err := json.NewEncoder(w).Encode(university)

	if err != nil {
		http.Error(w, "Error during encoding "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusNoContent)
}
