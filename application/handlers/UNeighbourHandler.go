package handlers

import (
	"Assignment-1/application/combine"
	"Assignment-1/application/requests"
	"Assignment-1/constants"
	"Assignment-1/structs"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

/*
A handlers for the university neighbour searcher part of the application.
Parameters:

	w: ResponseWriter (user of application) to write error message to.
	r: A request pointer given by the user
*/
func UNeighbourHandler(w http.ResponseWriter, r *http.Request) {

	//Use switch case to more easily extend application in the future
	//TODO: IMPLEMENTER LIMIT
	switch r.Method {
	case http.MethodGet:
		handleGetRequestN(w, r)
	default:
		http.Error(w, "REST method '"+r.Method+"' not currently supported. At this moment"+
			"only '"+http.MethodGet+"' are supported.", http.StatusNotImplemented)
		return
	}

}

/*
A function for handling a get-request for the university neighbour page.
*/
func handleGetRequestN(w http.ResponseWriter, r *http.Request) {
	var unispluscountries []structs.UniAndCountry
	var country []structs.Country

	countryValue := getCountryName(w, r)
	uniValue := getUniName(w, r)
	limit := getLimit(w, r)

	resp, err := requests.Request(constants.COUNTRIES_URL+"v3.1/name/"+countryValue, http.MethodGet)

	if err != nil {
		http.Error(w, "Faulty request", http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(resp.Body)
	errC := decoder.Decode(&country)

	if errC != nil {
		http.Error(w, "Error while decoding country values", http.StatusInternalServerError)
	}

	unis := requests.RequestUniversities(uniValue)

	if country != nil {
		country = requests.RequestCountriesByCCA(country[0].Borders)
	}

	if country != nil && unis != nil {
		unispluscountries = combine.CombineUniAndCountry(unis, country, limit)
	} else {
		http.Error(w, "No results to show! Please try another search!", http.StatusBadRequest)
	}

	w.Header().Add("content-type", "application/json")

	err = json.NewEncoder(w).Encode(unispluscountries)

	if err != nil {
		http.Error(w, "Error during encoding of unis and countries "+err.Error(), http.StatusInternalServerError)
		return
	}

	// No content if no action is taken above this point.
	http.Error(w, "", http.StatusNoContent)
}

/*
A function for finding the search word (country) of the users request (r).

Returns the user request in string format.
*/
func getCountryName(w http.ResponseWriter, r *http.Request) string {
	parts := strings.Split(r.URL.String(), "/")

	if len(parts) > 7 {
		http.Error(w, "Too many arguments, please enter input as such: "+
			"'neighbourunis/{country_name}/{partial_or_complete_university_name}'", http.StatusBadRequest)
		return ""
		//ERROR, NEED LESS ARGUMENTS
	}
	value := parts[4]
	if len(value) == 0 {
		http.Error(w, "Kindly provide a valid country name!", http.StatusBadRequest)
		return ""
		//ERROR, VALUE IS EMPTY AND NOT A VALID COUNTRY
	}

	return value
}

/*
A function for finding the search word (uni) of the users request (r).

Returns the user request in string format.
*/
func getUniName(w http.ResponseWriter, r *http.Request) string {
	var value string
	parts := strings.Split(r.URL.Path, "/")
	name := strings.Split(parts[5], "?")

	if len(parts) > 7 {
		http.Error(w, "Too many arguments, please enter input as such: "+
			"'neighbourunis/{country_name}/{partial_or_complete_university_name}{?limit={:number}}'", http.StatusBadRequest)
		return ""
	}
	if name != nil {
		value = name[0]
	} else {
		value = parts[5]
	}

	if len(value) == 0 {
		http.Error(w, "Kindly provide a valid university name!", http.StatusBadRequest)
		return ""
	}

	return value
}

/*
A function for finding the limit of the users request (r).

Returns the limit as an integer.
*/
func getLimit(w http.ResponseWriter, r *http.Request) int {
	var number []string
	var parts []string

	//If user has specified limit
	if strings.Contains(r.URL.RequestURI(), "?") {
		parts = strings.Split(r.URL.RequestURI(), "?")
	} else {
		return 0
	}

	//If parts is valid, and no / is in the string
	if parts != nil && !strings.Contains(parts[1], "/") {
		number = strings.Split(parts[1], "=")
	} else {
		http.Error(w, "Please enter value without a '/' in the end!", http.StatusBadRequest)
		return -1
	}

	if len(parts) > 2 {
		http.Error(w, "Invalid search, please enter input as such: "+
			"'neighbourunis/{country_name}/{partial_or_complete_university_name}{?limit={:number}}'", http.StatusBadRequest)
		return -1
	}

	//convert string to integer
	limit, err := strconv.Atoi(number[1])
	if limit == 0 || err != nil {
		http.Error(w, "Kindly provide a valid limit!", http.StatusBadRequest)
		return -1
	}

	return limit
}
