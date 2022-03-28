package shipengine_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *ShipEngine) TrackPackage(carrier Carrier, trackingNumber string) (*TrackingResponse, error) {
	request, err := http.NewRequest("GET", fmt.Sprintf(TrackPackage, carrier, trackingNumber), nil)
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

	var trackingResponse TrackingResponse
	err = json.Unmarshal(body, &trackingResponse)
	if err != nil {
		return nil, err
	}

	return &trackingResponse, nil
}