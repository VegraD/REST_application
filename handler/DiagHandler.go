package handler

import (
	"Assignment-1/constants"
	"Assignment-1/requests"
	"Assignment-1/structs"
	"Assignment-1/uptime"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
A handler for the /diag path of the application. When accessed, provides information regarding the uptime of the server.
Parameters:
	w: ResponseWriter (user of application) to write error message to.
	r: A request pointer given by the user
*/

func DiagHandler(w http.ResponseWriter, r *http.Request) {

	//Use switch case to more easily extend application in the future
	switch r.Method {
	case http.MethodGet:
		handleGetRequestDiag(w, r)
	default:
		http.Error(w, "REST method '"+r.Method+"' not currently supported. At this moment "+
			"only '"+http.MethodGet+"' are supported.", http.StatusNotImplemented)
		return
	}

}

// A function for handling get requests on the /diag path.
func handleGetRequestDiag(w http.ResponseWriter, r *http.Request) {

	diag := getDiag(w)

	w.Header().Add("content-type", "application/json")

	err := json.NewEncoder(w).Encode(diag)

	// Returns if encoding is faulty
	if err != nil {
		http.Error(w, "Error during encoding of diagnostics "+err.Error(), http.StatusInternalServerError)
		return
	}

	// No content if no action is taken above this point.
	http.Error(w, "", http.StatusNoContent)
}

/*
A function for acquiring the diagnostics struct.
Parameters:

	w: ResponseWriter (user of application) to write error message to (here: optional).

Returns a Diagnostics struct filled with upkeep information about the server.
*/
func getDiag(w http.ResponseWriter) structs.Diagnostic {
	// Requests a response from the APIs
	uReq, err := requests.Request(constants.UNI_URL+"search?name=", http.MethodGet)

	// return empty struct if request is faulty
	if err != nil {
		http.Error(w, "Error in fetching university request.", http.StatusInternalServerError)
		return structs.Diagnostic{}
	}

	cReq, err := requests.Request(constants.COUNTRIES_URL+"v3.1/all", http.MethodGet)
	if err != nil {
		http.Error(w, "Error in fetching country request.", http.StatusInternalServerError)
		return structs.Diagnostic{}
	}

	universityResp := uReq.StatusCode
	countryResp := cReq.StatusCode

	return structs.Diagnostic{
		UniAPI:     fmt.Sprintf("%d", universityResp),
		CountryAPI: fmt.Sprintf("%d", countryResp),
		Version:    "v1",
		Uptime:     uptime.GetUptime(),
	}
}
