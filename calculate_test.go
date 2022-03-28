package shipengine_api

import (
	"fmt"
	"os"
	"testing"
)

func TestCalculate(t *testing.T) {
	shipEngine := New(os.Getenv("API_KEY"))

	calculation := CalculateShipping{
		RateOptions: RateOptions{
			CarrierIds: []CarrierID{UPS_SANDBOX},
		},
		Shipment:    Shipment{
			ValidateAddress: "no_validation",
			ShipTo:          ShipTo{
				Name:                        "Amanda Miller",
				Phone:                       "555-555-5555",
				AddressLine1:                "525 S Winchester Blvd",
				CityLocality:                "San Jose",
				StateProvince:               "CA",
				PostalCode:                  "95128",
				CountryCode:                 "US",
				AddressResidentialIndicator: "yes",
			},
			ShipFrom:        ShipFrom{
				CompanyName:                 "Example Corp.",
				Name:                        "John Doe",
				Phone:                       "111-111-1111",
				AddressLine1:                "4009 Marathon Blvd",
				AddressLine2:                "Suite 300",
				CityLocality:                "Austin",
				StateProvince:               "TX",
				PostalCode:                  "78756",
				CountryCode:                 "US",
				AddressResidentialIndicator: "no",
			},
			Packages:        []Packages{
				{
					Weight: Weight{
						Value: 1.0,
						Unit:  "ounce",
					},
				},
			},
		},
	}
	calculated, err := shipEngine.CalculateShipping(calculation)
	if err != nil {
		t.Errorf("Error calculating shipping: %s", err.Error())
	}
	fmt.Println(calculated.RateResponse.Status)
}
