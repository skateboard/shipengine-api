package shipengine_api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (s *ShipEngine) CalculateShipping(calculateShipping CalculateShipping) (*CalculateShippingResponse, error) {
	jsonBytes, err := json.Marshal(calculateShipping)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", Rates, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = *s.header

	response, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := readBody(response)
	if err != nil {
		return nil, err
	}

	var calculateShippingResponse CalculateShippingResponse
	err = json.Unmarshal(body, &calculateShippingResponse)
	if err != nil {
		return nil, err
	}

	return &calculateShippingResponse, nil
}