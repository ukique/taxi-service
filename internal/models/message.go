package models

type IncomingMessage struct {
	Type string `json:"type"`
	Page int    `json:"page"`
}

type OutgoingOrdersMessage struct {
	Type string  `json:"type"`
	Data []Order `json:"data"`
}

type OutgoingDriversMessage struct {
	Type string   `json:"type"`
	Data []Driver `json:"data"`
}
