package handlers

import (
	"Assignment-1/application/combine"
	"Assignment-1/application/requests"
	"Assignment-1/structs"
	"encoding/json"
	"net/http"
	"strings"
)

/*
A handlers for the university info page of the application.
Parameters:

	w: ResponseWriter (user of application) to write error message to.
	r: A request pointer given by the user
*/
func UInfoHandler(w http.ResponseWriter, r *http.Request) {

	//Use switch case to more easily extend application in the future
	switch r.Method {
	case http.MethodGet:
		handleGetRequestU(w, r)
	default:
		http.Error(w, "REST method '"+r.Method+"' not currently supported. At this moment "+
			"only '"+http.MethodGet+"' are supported.", http.StatusNotImplemented)
		return
	}

}

/*
A function for handling a get-request for the university info page.
*/
func handleGetRequestU(w http.ResponseWriter, r *http.Request) {
	// Preallocating variables
	var unis []structs.University
	var unispluscountries []structs.UniAndCountry

	//search value given by user
	sValue := getSearchValue(w, r)

	unis = requests.RequestUniversities(sValue)

	countries := requests.RequestUniCountries(unis)

	// if neither countries nor unis are valid, show error
	if countries != nil && unis != nil {
		unispluscountries = combine.CombineUniAndCountry(unis, countries)
	} else {
		http.Error(w, "No results to show! Please try another search!", http.StatusBadRequest)
	}

	w.Header().Add("content-type", "application/json")

	err := json.NewEncoder(w).Encode(unispluscountries)

	if err != nil {
		http.Error(w, "Error during encoding of unis and countries "+err.Error(), http.StatusInternalServerError)
		return
	}

	// No content if no action is taken above this point.
	http.Error(w, "", http.StatusNoContent)
}

// TODO: ABSTRAHER
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
