package models

type ChatAndMessages struct {
	ChatInfo Chat `json:"chatInfo"`
	ChatMessages []Message `json:"chatMessages"`
}
