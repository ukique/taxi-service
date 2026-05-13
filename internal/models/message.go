package models

type IncomingMessage struct {
	Type string `json:"type"`
	Page int    `json:"page"`
}
	
type OutgoingMessage[T any] struct {
	Type string `json:"type"`
	Data T      `json:"data"`
}
