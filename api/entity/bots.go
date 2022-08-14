package entity

type Bot struct {
	Id       string `json:"id"`
	Status   string `json:"status"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
	ZoneId string `json:"zone_id"`
}

type PostBot struct {
	Status   string `json:"status"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
	ZoneId string `json:"zone_id"`
}
