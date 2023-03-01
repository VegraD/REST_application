package handler

import (
	"Assignment-1/structs"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func UInfoHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 5 {
		//ERROR, NEED ONE MORE ARGUMENT
	}

	value := parts[4]
	if len(value) == 0 {
		//ERROR, VALUE IS NULL AND NOT A VALID COUNTRY
	}

	//if()

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
