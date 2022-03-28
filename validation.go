package shipengine_api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (s *ShipEngine) ValidateAddress(address Address) (*ValidateAddressResponse, error) {
	addresses, err := s.ValidateAddresses([]Address{address})
	if err != nil {
		return nil, err
	}

	return &addresses[0], nil
}

func (s *ShipEngine) ValidateAddresses(addresses []Address) ([]ValidateAddressResponse, error) {
	jsonBytes, err := json.Marshal(addresses)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ValidateAddress, bytes.NewBuffer(jsonBytes))
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

	var validateAddressResponse []ValidateAddressResponse
	err = json.Unmarshal(body, &validateAddressResponse)
	if err != nil {
		return nil, err
	}

	return validateAddressResponse, nil
}