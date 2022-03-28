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
			ServiceCodes: []string{
				"ups_ground",
				"ups_3_day_select",
			},
		},
		Shipment:    Shipment{
			ValidateAddress: "no_validation",
			ShipTo:          ShipTo{
				Name:                        "Amanda Miller",
				Phone:                       "",
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
	//fmt.Println(calculated)
	fmt.Println(calculated.RateResponse.Rates)
}
