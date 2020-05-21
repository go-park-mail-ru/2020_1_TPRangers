package models

type JSONEvent struct {
	Event   string  `json:"event,omitempty"`
	Message Message `json:"message,omitempty"`
}
