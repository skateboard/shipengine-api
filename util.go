package shipengine_api

import (
	"io/ioutil"
	"net/http"
)

func readBody(response *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}