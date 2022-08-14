package entity

type Delivery struct {
	Id           string `json:"id"`
	CreationDate string `json:"creation_date"`
	State        string `json:"state"`
	Pickup       struct {
		PickupLat float64 `json:"pickup_lat"`
		PickupLon float64 `json:"pickup_lon"`
	} `json:"pickup"`
	Dropoff struct {
		DropoffLat float64 `json:"dropoff_lat"`
		DropoffLon float64 `json:"dropoff_lon"`
	} `json:"dropoff"`
	ZoneId string `json:"zone_id"`
}

type PostDelivery struct {
	State  string `json:"state"`
	Pickup struct {
		PickupLat float64 `json:"pickup_lat"`
		PickupLon float64 `json:"pickup_lon"`
	} `json:"pickup"`
	Dropoff struct {
		DropoffLat float64 `json:"dropoff_lat"`
		DropoffLon float64 `json:"dropoff_lon"`
	} `json:"dropoff"`
	ZoneId string `json:"zone_id"`
}
