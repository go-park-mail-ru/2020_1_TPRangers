package models

type Message struct {
	ChatId        string `json:"chatId, omitempty"`
	ChatPhoto     string `json:"chatPhoto"`
	ChatName      string `json:"chatName"`
	AuthorName    string `json:"authorName, omitempty"`
	AuthorSurname string `json:"authorSurname, omitempty"`
	AuthorUrl     string `json:"authorUrl, omitempty"`
	AuthorPhoto   string `json:"authorPhoto, omitempty"`
	Text          string `json:"text, omitempty"`
	Time          string `json:"time, omitempty"`
	Sticker       string `json:"sticker,omitempty"`
	IsMe          bool   `json:"isMe"`
}
