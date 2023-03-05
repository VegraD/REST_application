package handler

import (
	"Assignment-1/constants"
	"Assignment-1/requests"
	"Assignment-1/structs"
	"encoding/json"
	"net/http"
	"strings"
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
	var unispluscountries []structs.UniAndCountry
	var country []structs.Country

	//TODO: CAN ABSTRACT LATER INTO COMMON FUNCTION WHICH TAKES AN "lines" parameter
	countryValue := getCountryName(w, r)

	uniValue := getUniName(w, r)

	resp, err := requests.Request(constants.COUNTRIES_URL+"v3.1/name/"+countryValue, http.MethodGet)

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&country)

	if err != nil {
		http.Error(w, "Error while decoding country values", http.StatusInternalServerError)
	}

	borderCountries := requests.RequestCountriesByCCA(country[0].Borders)

	unis := requests.RequestUniversities(uniValue)

	//Finne countryValue sine borders[]
	//requeste disse countriesene
	//Finne universities for disse landene gjennom requestcountries
	//5
	if borderCountries != nil && unis != nil {
		unispluscountries = CombineUniAndCountry(unis, borderCountries)
	} else {
		http.Error(w, "No results to show! Please try another search!", http.StatusBadRequest)
	}

	w.Header().Add("content-type", "application/json")

	err = json.NewEncoder(w).Encode(unispluscountries)

	if err != nil {
		http.Error(w, "Error during encoding of unis and countries "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusNoContent)
}

func getCountryName(w http.ResponseWriter, r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 6 {
		http.Error(w, "Too many arguments, please enter input as such: "+
			"'neighbourunis/{country_name}/{partial_or_complete_university_name}'", http.StatusBadRequest)
		return ""
		//ERROR, NEED ONE MORE ARGUMENT
	}
	value := parts[4]
	if len(value) == 0 {
		http.Error(w, "Kindly provide a valid country name!", http.StatusBadRequest)
		return ""
		//ERROR, VALUE IS NULL AND NOT A VALID COUNTRY
	}

	return value
}

func getUniName(w http.ResponseWriter, r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 6 {
		http.Error(w, "Too many arguments, please enter input as such: "+
			"'neighbourunis/{country_name}/{partial_or_complete_university_name}'", http.StatusBadRequest)
		return ""
		//ERROR, NEED ONE MORE ARGUMENT
	}
	value := parts[5]
	if len(value) == 0 {
		http.Error(w, "Kindly provide a valid university name!", http.StatusBadRequest)
		return ""
		//ERROR, VALUE IS NULL AND NOT A VALID COUNTRY
	}

	return value
}
