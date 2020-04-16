package models

type Message struct {
	ChatId        int    `json:"chatId , omitempty"`
	ChatPhoto     string `json:"chatPhoto"`
	ChatName      string `json:"chatName"`
	AuthorName    string `json:"authorName , omitempty"`
	AuthorSurname string `json:"authorSurname , omitempty"`
	AuthorUrl     string `json:"authorUrl , omitempty"`
	AuthorPhoto   string `json:"authorPhoto , omitempty"`
	Text          string `json:"text , omitempty"`
}
