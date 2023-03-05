package handler

import (
	"Assignment-1/convert"
	"Assignment-1/requests"
	"Assignment-1/structs"
	"encoding/json"
	"net/http"
	"strings"
)

func UInfoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetRequestU(w, r)
	default:
		http.Error(w, "REST method '"+r.Method+"' not currently supported. At this moment "+
			"only '"+http.MethodGet+"' are supported.", http.StatusNotImplemented)
		return
	}

}

func handleGetRequestU(w http.ResponseWriter, r *http.Request) {
	var unis []structs.University
	var unispluscountries []structs.UniAndCountry

	sValue := getSearchValue(w, r)

	unis = requests.RequestUniversities(sValue)

	countries := requests.RequestUniCountries(unis)

	if countries != nil && unis != nil {
		unispluscountries = CombineUniAndCountry(unis, countries)
	} else {
		http.Error(w, "No results to show! Please try another search!", http.StatusBadRequest)
	}

	w.Header().Add("content-type", "application/json")

	err := json.NewEncoder(w).Encode(unispluscountries)

	if err != nil {
		http.Error(w, "Error during encoding of unis and countries "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusNoContent)
}

func getSearchValue(w http.ResponseWriter, r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 5 {
		http.Error(w, "Too many arguments, please enter input as such: "+
			"'uniinfo/{partial_or_complete_university_name}'", http.StatusBadRequest)
		return ""
		//ERROR, NEED ONE MORE ARGUMENT
	}
	value := parts[4]
	if len(value) == 0 {
		http.Error(w, "Kindly provide a valid university name!", http.StatusBadRequest)
		return ""
		//ERROR, VALUE IS NULL AND NOT A VALID COUNTRY
	}

	return value
}

// TODO: ADD IN CONVERT PACKAGE OR SOMETHING SIMILAR
func CombineUniAndCountry(unis []structs.University, countries []structs.Country) []structs.UniAndCountry {
	var outputs []structs.UniAndCountry

	for _, i := range unis {
		for _, j := range countries {
			if i.Country == j.Name["common"] {
				outputs = append(outputs, convert.ToUniAndCountry(i, j))
			}
		}

	}
	return outputs
}
