package shipengine_api

import "time"

type Address struct {
	AddressLine1  string `json:"address_line1"`
	AddressLine2  string `json:"address_line2"`
	CityLocality  string `json:"city_locality"`
	StateProvince string `json:"state_province"`
	PostalCode    string `json:"postal_code"`
	CountryCode   string `json:"country_code"`
}

type ValidateAddressResponse struct {
	Status          string `json:"status"`
	OriginalAddress struct {
		Name                        interface{} `json:"name"`
		Phone                       interface{} `json:"phone"`
		CompanyName                 interface{} `json:"company_name"`
		AddressLine1                string      `json:"address_line1"`
		AddressLine2                interface{} `json:"address_line2"`
		AddressLine3                interface{} `json:"address_line3"`
		CityLocality                string      `json:"city_locality"`
		StateProvince               string      `json:"state_province"`
		PostalCode                  string      `json:"postal_code"`
		CountryCode                 string      `json:"country_code"`
		AddressResidentialIndicator string      `json:"address_residential_indicator"`
	} `json:"original_address"`
	MatchedAddress struct {
		Name                        interface{} `json:"name"`
		Phone                       interface{} `json:"phone"`
		CompanyName                 interface{} `json:"company_name"`
		AddressLine1                string      `json:"address_line1"`
		AddressLine2                string      `json:"address_line2"`
		AddressLine3                interface{} `json:"address_line3"`
		CityLocality                string      `json:"city_locality"`
		StateProvince               string      `json:"state_province"`
		PostalCode                  string      `json:"postal_code"`
		CountryCode                 string      `json:"country_code"`
		AddressResidentialIndicator string      `json:"address_residential_indicator"`
	} `json:"matched_address"`
	Messages []interface{} `json:"messages"`
}

type TrackingResponse struct {
	TrackingNumber           string      `json:"tracking_number"`
	TrackingUrl              string      `json:"tracking_url"`
	StatusCode               string      `json:"status_code"`
	CarrierCode              string      `json:"carrier_code"`
	CarrierId                int         `json:"carrier_id"`
	CarrierDetailCode        interface{} `json:"carrier_detail_code"`
	StatusDescription        string      `json:"status_description"`
	CarrierStatusCode        string      `json:"carrier_status_code"`
	CarrierStatusDescription string      `json:"carrier_status_description"`
	ShipDate                 interface{} `json:"ship_date"`
	EstimatedDeliveryDate    interface{} `json:"estimated_delivery_date"`
	ActualDeliveryDate       interface{} `json:"actual_delivery_date"`
	ExceptionDescription     interface{} `json:"exception_description"`
	Events                   []struct {
		OccurredAt               time.Time   `json:"occurred_at"`
		CarrierOccurredAt        string      `json:"carrier_occurred_at"`
		Description              string      `json:"description"`
		CityLocality             string      `json:"city_locality"`
		StateProvince            string      `json:"state_province"`
		PostalCode               string      `json:"postal_code"`
		CountryCode              string      `json:"country_code"`
		CompanyName              string      `json:"company_name"`
		Signer                   string      `json:"signer"`
		EventCode                string      `json:"event_code"`
		CarrierDetailCode        interface{} `json:"carrier_detail_code"`
		StatusCode               interface{} `json:"status_code"`
		StatusDescription        interface{} `json:"status_description"`
		CarrierStatusCode        string      `json:"carrier_status_code"`
		CarrierStatusDescription string      `json:"carrier_status_description"`
		Latitude                 *float64    `json:"latitude"`
		Longitude                *float64    `json:"longitude"`
	} `json:"events"`
}

type CalculateShipping struct {
	RateOptions RateOptions `json:"rate_options"`
	Shipment Shipment `json:"shipment"`
}

type RateOptions struct {
	CarrierIds []CarrierID `json:"carrier_ids"`
}

type Shipment struct {
	ValidateAddress string `json:"validate_address"`
	ShipTo 		 ShipTo `json:"ship_to"`
	ShipFrom ShipFrom `json:"ship_from"`
	Packages []Packages `json:"packages"`
}

type ShipTo struct {
	Name                        string `json:"name"`
	Phone                       string `json:"phone"`
	AddressLine1                string `json:"address_line1"`
	CityLocality                string `json:"city_locality"`
	StateProvince               string `json:"state_province"`
	PostalCode                  string `json:"postal_code"`
	CountryCode                 string `json:"country_code"`
	AddressResidentialIndicator string `json:"address_residential_indicator"`
}

type ShipFrom struct {
	CompanyName                 string `json:"company_name"`
	Name                        string `json:"name"`
	Phone                       string `json:"phone"`
	AddressLine1                string `json:"address_line1"`
	AddressLine2                string `json:"address_line2"`
	CityLocality                string `json:"city_locality"`
	StateProvince               string `json:"state_province"`
	PostalCode                  string `json:"postal_code"`
	CountryCode                 string `json:"country_code"`
	AddressResidentialIndicator string `json:"address_residential_indicator"`
}

type Packages struct {
	Weight Weight `json:"weight"`
}

type Weight struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type Dimensions struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Length int    `json:"length"`
	Unit   string `json:"unit"`
}

type LabelPackages struct {
	Weight Weight `json:"weight"`
	Dimensions Dimensions `json:"dimensions"`
}

type CalculateShippingResponse struct {
	RateResponse struct {
		RateRequestId string       `json:"rate_request_id"`
		ShipmentId    string    `json:"shipment_id"`
		Status        string    `json:"status"`
		CreatedAt     time.Time `json:"created_at"`
		Rates         []struct {
			RateId         string `json:"rate_id"`
			RateType       string `json:"rate_type"`
			CarrierId      string `json:"carrier_id"`
			ShippingAmount struct {
				Currency string  `json:"currency"`
				Amount   float64 `json:"amount"`
			} `json:"shipping_amount"`
			InsuranceAmount struct {
				Currency string  `json:"currency"`
				Amount   float64 `json:"amount"`
			} `json:"insurance_amount"`
			ConfirmationAmount struct {
				Currency string  `json:"currency"`
				Amount   float64 `json:"amount"`
			} `json:"confirmation_amount"`
			OtherAmount struct {
				Currency string  `json:"currency"`
				Amount   float64 `json:"amount"`
			} `json:"other_amount"`
			DeliveryDays          int           `json:"delivery_days"`
			GuaranteedService     bool          `json:"guaranteed_service"`
			EstimatedDeliveryDate time.Time     `json:"estimated_delivery_date"`
			CarrierDeliveryDays   string        `json:"carrier_delivery_days"`
			ShipDate              time.Time     `json:"ship_date"`
			NegotiatedRate        bool          `json:"negotiated_rate"`
			ServiceType           string        `json:"service_type"`
			ServiceCode           string        `json:"service_code"`
			Trackable             bool          `json:"trackable"`
			ValidationStatus      string        `json:"validation_status"`
			WarningMessages       []interface{} `json:"warning_messages"`
			ErrorMessages         []interface{} `json:"error_messages"`
			CarrierCode           string        `json:"carrier_code"`
			CarrierNickname       string        `json:"carrier_nickname"`
			CarrierFriendlyName   string        `json:"carrier_friendly_name"`
		} `json:"rates"`
		InvalidRates []interface{} `json:"invalid_rates"`
	} `json:"rate_response"`
	ShipmentId         string      `json:"shipment_id"`
	CarrierId          string      `json:"carrier_id"`
	ExternalShipmentId interface{} `json:"external_shipment_id"`
	ShipDate           time.Time   `json:"ship_date"`
	CreatedAt          time.Time   `json:"created_at"`
	ModifiedAt         time.Time   `json:"modified_at"`
	ShipmentStatus     string      `json:"shipment_status"`
	ShipTo             struct {
		Name                        string `json:"name"`
		Phone                       string `json:"phone"`
		AddressLine1                string `json:"address_line1"`
		CityLocality                string `json:"city_locality"`
		StateProvince               string `json:"state_province"`
		PostalCode                  string `json:"postal_code"`
		CountryCode                 string `json:"country_code"`
		AddressResidentialIndicator string `json:"address_residential_indicator"`
	} `json:"ship_to"`
	ShipFrom struct {
		CompanyName                 string `json:"company_name"`
		Name                        string `json:"name"`
		Phone                       string `json:"phone"`
		AddressLine1                string `json:"address_line1"`
		AddressLine2                string `json:"address_line2"`
		CityLocality                string `json:"city_locality"`
		StateProvince               string `json:"state_province"`
		PostalCode                  string `json:"postal_code"`
		CountryCode                 string `json:"country_code"`
		AddressResidentialIndicator string `json:"address_residential_indicator"`
	} `json:"ship_from"`
	ReturnTo struct {
		CompanyName                 string `json:"company_name"`
		Name                        string `json:"name"`
		Phone                       string `json:"phone"`
		AddressLine1                string `json:"address_line1"`
		AddressLine2                string `json:"address_line2"`
		CityLocality                string `json:"city_locality"`
		StateProvince               string `json:"state_province"`
		PostalCode                  string `json:"postal_code"`
		CountryCode                 string `json:"country_code"`
		AddressResidentialIndicator string `json:"address_residential_indicator"`
	} `json:"return_to"`
	Confirmation    string `json:"confirmation"`
	AdvancedOptions struct {
		BillToAccount     interface{} `json:"bill_to_account"`
		BillToCountryCode interface{} `json:"bill_to_country_code"`
		BillToParty       interface{} `json:"bill_to_party"`
		BillToPostalCode  interface{} `json:"bill_to_postal_code"`
		ContainsAlcohol   bool        `json:"contains_alcohol"`
		CustomField1      interface{} `json:"custom_field1"`
		CustomField2      interface{} `json:"custom_field2"`
		CustomField3      interface{} `json:"custom_field3"`
		NonMachinable     bool        `json:"non_machinable"`
		SaturdayDelivery  bool        `json:"saturday_delivery"`
	} `json:"advanced_options"`
	InsuranceProvider string        `json:"insurance_provider"`
	Tags              []interface{} `json:"tags"`
	TotalWeight       struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	} `json:"total_weight"`
	Packages []struct {
		Weight struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"weight"`
		Dimensions struct {
			Unit   string  `json:"unit"`
			Length float64 `json:"length"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"dimensions"`
		InsuredValue struct {
			Currency string  `json:"currency"`
			Amount   float64 `json:"amount"`
		} `json:"insured_value"`
	} `json:"packages"`
}

type Label struct {
	Shipment LabelShipment `json:"shipment"`
}

type LabelShipment struct {
	ServiceCode string `json:"service_code"`
	ShipTo ShipTo `json:"ship_to"`
	ShipFrom ShipFrom `json:"ship_from"`
	Packages []LabelPackages `json:"packages"`
}

type LabelResponse struct {
	LabelId      string    `json:"label_id"`
	Status       string    `json:"status"`
	ShipmentId   string    `json:"shipment_id"`
	ShipDate     time.Time `json:"ship_date"`
	CreatedAt    time.Time `json:"created_at"`
	ShipmentCost struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
	} `json:"shipment_cost"`
	InsuranceCost struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
	} `json:"insurance_cost"`
	ChargeEvent     string      `json:"charge_event"`
	TrackingNumber  string      `json:"tracking_number"`
	IsReturnLabel   bool        `json:"is_return_label"`
	RmaNumber       interface{} `json:"rma_number"`
	IsInternational bool        `json:"is_international"`
	BatchId         string      `json:"batch_id"`
	CarrierId       string      `json:"carrier_id"`
	ServiceCode     string      `json:"service_code"`
	PackageCode     string      `json:"package_code"`
	Voided          bool        `json:"voided"`
	VoidedAt        interface{} `json:"voided_at"`
	LabelFormat     string      `json:"label_format"`
	LabelLayout     string      `json:"label_layout"`
	Trackable       bool        `json:"trackable"`
	LabelImageId    interface{} `json:"label_image_id"`
	CarrierCode     string      `json:"carrier_code"`
	TrackingStatus  string      `json:"tracking_status"`
	LabelDownload   struct {
		Pdf  string `json:"pdf"`
		Png  string `json:"png"`
		Zpl  string `json:"zpl"`
		Href string `json:"href"`
	} `json:"label_download"`
	FormDownload   interface{} `json:"form_download"`
	InsuranceClaim interface{} `json:"insurance_claim"`
	Packages       []struct {
		PackageCode string `json:"package_code"`
		Weight      struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"weight"`
		Dimensions struct {
			Unit   string  `json:"unit"`
			Length float64 `json:"length"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"dimensions"`
		InsuredValue struct {
			Currency string  `json:"currency"`
			Amount   float64 `json:"amount"`
		} `json:"insured_value"`
		TrackingNumber string `json:"tracking_number"`
		LabelMessages  struct {
			Reference1 interface{} `json:"reference1"`
			Reference2 interface{} `json:"reference2"`
			Reference3 interface{} `json:"reference3"`
		} `json:"label_messages"`
		ExternalPackageId interface{} `json:"external_package_id"`
	} `json:"packages"`
}