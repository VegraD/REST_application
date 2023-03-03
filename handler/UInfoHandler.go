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
	var country []structs.Country
	var unispluscountries []structs.UniAndCountry

	sValue := getSearchValue(w, r)
	//resp, err := requests.Request(constants.UNI_URL+"search?name=middle&country=Turkey", http.MethodGet)

	unis = searchUniversities(w, sValue)
	if len(unis) != 0 {
		country = searchCountry(w, unis[0])
	} else {
		country = nil
	}

	if country != nil {
		unispluscountries = combineCountries(unis, country)
	}

	w.Header().Add("content-type", "application/json")

	err := json.NewEncoder(w).Encode(unispluscountries)

	if err != nil {
		http.Error(w, "Error during encoding "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusNoContent)
}

//TODO: Method for combining when there are several universities
//TODO: Method for requesting university + country separately

func getSearchValue(w http.ResponseWriter, r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 5 {
		http.Error(w, "Too many arguments, please enter input as such: 'uniinfo/{country}'", http.StatusBadRequest)
		return ""
		//ERROR, NEED ONE MORE ARGUMENT
	}
	value := parts[4]
	if len(value) == 0 {
		http.Error(w, "Kindly provide a valid country name!", http.StatusBadRequest)
		return ""
		//ERROR, VALUE IS NULL AND NOT A VALID COUNTRY
	}

	return parts[4]
}

// KAN VURDERE Å FLYTTE TIL REQUEST
func searchUniversities(w http.ResponseWriter, value string) []structs.University {
	var unis []structs.University

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
func searchCountry(w http.ResponseWriter, uni structs.University) []structs.Country {
	var country []structs.Country

	resp, err := requests.Request(constants.COUNTRIES_URL+"v3.1/name/"+uni.Country, http.MethodGet)
	if err != nil {
		http.Error(w, "Error in response.", http.StatusBadRequest)
		return nil
	}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&country); err != nil {
		http.Error(w, "Error in country decoding", http.StatusInternalServerError)
		return nil
	}

	return country
}

func combineCountries(unis []structs.University, countries []structs.Country) []structs.UniAndCountry {
	var outputs []structs.UniAndCountry

	if countries != nil {
		country := countries[0]
		for _, i := range unis {
			outputs = append(outputs, convert.ToUniAndCountry(i, country))
		}
	}

	return outputs
}
