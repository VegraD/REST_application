package handler

import (
	"Assignment-1/constants"
	"Assignment-1/requests"
	"Assignment-1/structs"
	"encoding/json"
	"log"
	"net/http"
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

	//resp, err := requests.Request(constants.UNI_URL+"search?name=middle&country=Turkey", http.MethodGet)

	resp, err := requests.Request(constants.UNI_URL+"search?name=technology&country=norway", http.MethodGet)

	if err != nil {
		http.Error(w, "Error in response.", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(resp.Body)
	var uni []structs.University
	if err := decoder.Decode(&uni); err != nil {
		log.Println(err)
	}

	w.Header().Add("content-type", "application/json")

	err = json.NewEncoder(w).Encode(uni)

	if err != nil {
		http.Error(w, "Error during encoding "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusNoContent)
}
