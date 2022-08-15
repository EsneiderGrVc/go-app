package entity

type Assignment struct {
	DeliveryId     string `json:"delivery_id"`
	BotId          string `json:"bot_id"`
	AssignmentDate string `json:"assignment_date"`
}

type RequestAssignment struct {
	DeliveryId string `json:"delivery_id"`
	BotId      string `json:"bot_id"`
}
