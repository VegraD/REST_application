package requests

import "net/http"

func Request(url string, method string) (*http.Response, error) {

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	defer client.CloseIdleConnections()

	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req) //Do in case of extending application at a later point in time with more methods.

	if err != nil {
		return nil, err
	}

	return resp, nil
}
