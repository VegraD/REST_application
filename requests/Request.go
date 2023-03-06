package requests

import (
	"Assignment-1/constants"
	"Assignment-1/structs"
	"encoding/json"
	"net/http"
	"strings"
)

func Request(url string, method string) (*http.Response, error) {

	url = strings.ReplaceAll(url, " ", "%20")

	req, err := http.NewRequest(method, url, nil)
	
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")

	client := &http.Client{}

	resp, errC := client.Do(req) //Do in case of extending application at a later point in time with more methods.

	if errC != nil {
		return nil, errC
	}

	return resp, nil
}

func RequestUniversities(value string) []structs.University {
	var unis []structs.University

	if value == "" {
		return nil
	}

	resp, err := Request(constants.UNI_URL+
		"search?name="+value, http.MethodGet)
	if err != nil {
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&unis); err != nil {
		return nil
	}
	return unis
}

func RequestUniCountries(unis []structs.University) []structs.Country {
	var countries []structs.Country
	var countryCodes string

	for _, i := range unis {
		countryCodes += i.AlphaTwoCode + ","
	}
	resp, err := Request(constants.COUNTRIES_URL+"v3.1/alpha?codes="+countryCodes, http.MethodGet)
	if err != nil {
		return nil
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&countries)

	if err != nil {
		return nil
	}

	return countries
}

// KAN VURDERE Å SAMMENSVEISE DENNE MED DEN OVER!!!
func RequestCountriesByCCA(CCAS []string) []structs.Country {
	var countries []structs.Country
	var countryCodes string

	for _, i := range CCAS {
		countryCodes += i + ","
	}
	resp, err := Request(constants.COUNTRIES_URL+"v3.1/alpha?codes="+countryCodes, http.MethodGet)
	if err != nil {
		return nil
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&countries)

	if err != nil {
		return nil
	}

	return countries
}
