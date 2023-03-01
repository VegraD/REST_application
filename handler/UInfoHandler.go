package handler

import (
	"Assignment-1/structs"
	"encoding/json"
	"log"
	"net/http"
)

func UInfoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method "+r.Method+"is not supported. At this current time, only "+http.MethodGet+
			"are supported", http.StatusNotImplemented)
		return
	} else {
		handleGetRequestN(w, r)
	}

}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
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
		http.Error(w, "Error during encoding: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusNoContent)
}