package models

type Chat struct {
	IsGroupChat              bool   `json:"isGroupChat"`
	ChatName                 string `json:"chatName, omitempty"`
	ChatId                   string `json:"chatId, omitempty"`
	ChatPhoto                string `json:"chatPhoto, omitempty"`
	ChatCounter              int    `json:"chatCounter, omitempty"`
	OnlineStatus             bool   `json:"onlineStatus, omitempty"`
	LastMessageAuthorName    string `json:"lastMessageAuthorName, omitempty"`
	LastMessageAuthorSurname string `json:"lastMessageAuthorSurname, omitempty"`
	LastMessageTxt           string `json:"lastMessageTxt, omitempty"`
}
