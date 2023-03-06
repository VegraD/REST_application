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

func DiagHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleGetRequestDiag(w, r)
	default:
		http.Error(w, "REST method '"+r.Method+"' not currently supported. At this moment "+
			"only '"+http.MethodGet+"' are supported.", http.StatusNotImplemented)
		return
	}

}

func handleGetRequestDiag(w http.ResponseWriter, r *http.Request) {

	diag := getDiag(w)

	w.Header().Add("content-type", "application/json")

	err := json.NewEncoder(w).Encode(diag)

	if err != nil {
		http.Error(w, "Error during encoding of diagnostics "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusNoContent)
	//sende generelle requests til både UNIAPI og COUNTRYAPI og finne statuskodene
	//Putte gitte statuskoder inn i structs.
	//LAGE uptime.go som sjekker hvor lenge nettside har vært oppe i
}

func getDiag(w http.ResponseWriter) structs.Diagnostic {
	uReq, err := requests.Request(constants.UNI_URL+"search?name=", http.MethodGet)
	if err != nil {
		http.Error(w, "Error in fething university request.", http.StatusInternalServerError)
		return structs.Diagnostic{}
	}

	cReq, err := requests.Request(constants.COUNTRIES_URL+"v3.1/all", http.MethodGet)
	if err != nil {
		http.Error(w, "Error in fething country request.", http.StatusInternalServerError)
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
