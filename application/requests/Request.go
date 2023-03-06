package requests

import (
	"Assignment-1/constants"
	"Assignment-1/structs"
	"encoding/json"
	"net/http"
	"strings"
)

/*
A function for requesting a response from the APIs.
Parameters:

	url: The url of the request
	method: The method of which to use when requesting

Returns the response given from the API and an appropriate error code.
*/
func Request(url string, method string) (*http.Response, error) {

	//Replaces all spaces with %20 (space in searchbar), enables searches to be made with spaces.
	url = strings.ReplaceAll(url, " ", "%20")

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")

	client := &http.Client{}

	// Use Do instead if Get() in case of extending application at a later point in time with more methods.
	resp, errC := client.Do(req)

	if errC != nil {
		return nil, errC
	}

	return resp, nil
}

/*
A function for requesting university from a given keyword,
Parameters:

	value: The value of which to conduct the search with

Returns all universities in a slice which are partially equal or equal to the keyword.
*/
func RequestUniversities(value string) []structs.University {
	var unis []structs.University

	//return nil if value is empty
	if value == "" {
		return nil
	}

	//Request response from University API
	resp, err := Request(constants.UNI_URL+
		"search?name="+value, http.MethodGet)
	if err != nil {
	}

	//decode, and return
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&unis); err != nil {
		return nil
	}
	return unis
}

/*
A function for requesting countries for a slice of universities.
Parameters:

	unis: a slice of unis to find countries for.

Returns a slice of countries, where each country is one which a university is located in.
*/
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

/*
A function for requesting countries by CCAs.
Parameters:

	CCAS: a slice of CCAS to search by.

Returns a slice of countries which match the CCAS given as parameters.
*/
func RequestCountriesByCCA(CCAS []string) []structs.Country {
	var countries []structs.Country
	var countryCodes string

	// Combine into a search string
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
