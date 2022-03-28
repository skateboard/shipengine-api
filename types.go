package shipengine_api

var (
	ValidateAddress = "https://api.shipengine.com/v1/addresses/validate"
	TrackPackage = "https://api.shipengine.com/v1/tracking?carrier_code=%v&tracking_number=%v"
	Rates = "https://api.shipengine.com/v1/rates"
	Labels = "https://api.shipengine.com/v1/labels"
)

type Carrier string
const (
	USPS Carrier = "usps"
	STAMPS Carrier = "stamps_com"
	FEDEX Carrier = "fedex"
	UPS Carrier = "ups"
	DHL_EXPRESS Carrier = "dhl_express"
	DHL_ECOMMERCE Carrier = "dhl_global_mail"
	CANADA_POST Carrier = "canada_post"
	AUSTRALIA_POST Carrier = "australia_post"
	FIRSTMILE Carrier = "firstmile"
	ASENDIA Carrier = "asendia"
	ONTRAC Carrier = "ontrac"
	APC Carrier = "apc"
	NEWGISTICS Carrier = "newgistics"
	GLOBEGISTICS Carrier = "globegistics"
	RR_DONNELLEY Carrier = "rr_donnelley"
	IMEX Carrier = "imex"
	ACCESS_WORLDWIDE Carrier = "access_worldwide"
	PUROLATOR_CA Carrier = "purolator_ca"
	SENDLE Carrier = "sendle"
)

type CarrierID string
const (
	USPS_SANDBOX CarrierID = "se-2055161"
	UPS_SANDBOX CarrierID = "se-2055162"
	FEDEX_SANDBOX CarrierID = "se-2055163"
)