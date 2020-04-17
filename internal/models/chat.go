package models

type Chat struct {
	ChatName     string `json:"chatName , omitempty"`
	ChatId    string `json:"chatId , omitempty"`
	ChatPhoto    string `json:"chatName , omitempty"`
	ChatCounter  int    `json:"chatCounter , omitempty"`
	OnlineStatus bool   `json:"onlineStatus , omitempty"`
}
