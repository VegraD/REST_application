package handler

import (
	"Assignment-1/constants"
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
		handleGetRequest(w, r)
	default:
		http.Error(w, "REST method '"+r.Method+"' not currently supported. At this moment "+
			"only '"+http.MethodGet+"' are supported.", http.StatusNotImplemented)
		return
	}

}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	var unis []structs.University
	var unispluscountries []structs.UniAndCountry

	sValue := getSearchValue(w, r)

	unis = searchUniversities(w, sValue)

	countries := searchCountries(w, unis)

	unispluscountries = combineUniAndCountry(w, unis, countries)

	w.Header().Add("content-type", "application/json")

	err := json.NewEncoder(w).Encode(unispluscountries)

	if err != nil {
		http.Error(w, "Error during encoding of unis and countries "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusNoContent)
}

//TODO: Method for combining when there are several universities
//TODO: Method for requesting university + country separately

func getSearchValue(w http.ResponseWriter, r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 5 {
		http.Error(w, "Too many arguments, please enter input as such: 'uniinfo/{name/partial name}'", http.StatusBadRequest)
		return ""
		//ERROR, NEED ONE MORE ARGUMENT
	}
	value := parts[4]
	if len(value) == 0 {
		http.Error(w, "Kindly provide a valid university name!", http.StatusBadRequest)
		return ""
		//ERROR, VALUE IS NULL AND NOT A VALID COUNTRY
	}

	return parts[4]
}

// KAN VURDERE Å FLYTTE TIL REQUEST
func searchUniversities(w http.ResponseWriter, value string) []structs.University {

	var unis []structs.University

	if value == "" {
		http.Error(w, "University search not possible!", http.StatusBadRequest)
		return nil
	}

	resp, err := requests.Request(constants.UNI_URL+
		"search?name="+value, http.MethodGet)
	if err != nil {
		http.Error(w, "Error in response.", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&unis); err != nil {
		http.Error(w, "Error in university decoding.", http.StatusInternalServerError)
	}
	return unis
}

// KAN VURDERE Å FLYTTE TIL REQUEST
func searchCountries(w http.ResponseWriter, unis []structs.University) []structs.Country {
	var countries []structs.Country
	var countryCodes string
	for _, i := range unis {
		countryCodes += i.AlphaTwoCode + ","
	}
	resp, err := requests.Request(constants.COUNTRIES_URL+"v3.1/alpha?codes="+countryCodes, http.MethodGet)
	if err != nil {
		http.Error(w, "Error in response.", http.StatusBadRequest)
		return nil
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&countries)
	if err != nil {
		http.Error(w, "Error in country decoding", http.StatusInternalServerError)
		return nil
	} else {

	}

	return countries
}

func combineUniAndCountry(w http.ResponseWriter, unis []structs.University, countries []structs.Country) []structs.UniAndCountry {
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
