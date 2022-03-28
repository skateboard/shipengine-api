package shipengine_api

import (
	"fmt"
	"os"
	"testing"
)

func TestValidation(t *testing.T) {
	shipEngine := New(os.Getenv("API_KEY"))

	address := Address{
		AddressLine1:  "525 S Winchester Blvd",
		AddressLine2:  "",
		CityLocality:  "San Jose",
		StateProvince: "CA",
		PostalCode:    "95128",
		CountryCode:   "US",
	}
	validation, err := shipEngine.ValidateAddress(address)
	if err != nil {
		t.Errorf("Error validating address: %s", err.Error())
	}
	fmt.Println(validation.Status)
}
