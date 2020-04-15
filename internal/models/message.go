package models

type Message struct {
	Receiver int    `json:"receiver, omitempty"`
	Text     string `json:"text , omitempty"`
}
