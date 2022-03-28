package shipengine_api

import (
	"fmt"
	"os"
	"testing"
)

func TestLabel(t *testing.T) {
	shipEngine := New(os.Getenv("API_KEY"))

	label := Label{
		Shipment: LabelShipment{
			ServiceCode: "ups_ground",
			ShipTo:      ShipTo{
				Name:                        "Amanda Miller",
				Phone:                       "555-555-5555",
				AddressLine1:                "525 S Winchester Blvd",
				CityLocality:                "San Jose",
				StateProvince:               "CA",
				PostalCode:                  "95128",
				CountryCode:                 "US",
				AddressResidentialIndicator: "yes",
			},
			ShipFrom:    ShipFrom{
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
			Packages:  []LabelPackages{
				{
					Weight: Weight{
						Value: 1.0,
						Unit:  "ounce",
					},
					Dimensions: Dimensions{
						Height: 6,
						Width:  12,
						Length: 24,
						Unit:   "inch",
					},
				},
			},
		},
	}
	labelResponse, err := shipEngine.CreateLabel(label)
	if err != nil {
		t.Errorf("Error creating label: %s", err.Error())
	}
	fmt.Println(labelResponse.Status)
}