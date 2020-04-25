package models

type ChatAndMessages struct {
	ChatInfo ChatInfo `json:"chatInfo"`
	ChatMessages []Message `json:"chatMessages"`
}
