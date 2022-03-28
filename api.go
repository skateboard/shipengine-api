package shipengine_api

import (
	"net/http"
	"time"
)

type ShipEngine struct {
	header *http.Header
	client *http.Client
}

func New(apiKey string) ShipEngine {
	header := http.Header{
		"Content-Type": []string{"application/json"},
		"API-Key": 	 []string{apiKey},
		"Host": 	 []string{"api.shipengine.com"},
	}

	return ShipEngine{
		header: &header,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}