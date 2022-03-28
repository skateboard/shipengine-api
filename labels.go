package shipengine_api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (s *ShipEngine) CreateLabel(label Label) (*LabelResponse, error) {
	jsonBytes, err := json.Marshal(label)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", Labels, bytes.NewBuffer(jsonBytes))
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

	var labelResponse LabelResponse
	err = json.Unmarshal(body, &labelResponse)
	if err != nil {
		return nil, err
	}

	return &labelResponse, nil
}