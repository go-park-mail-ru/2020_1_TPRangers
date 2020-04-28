package models

type ChatInfo struct {
	IsGroupChat    bool   `json:"isGroupChat"`
	ChatId         string `json:"chatId, omitempty"`
	ChatCounter    int    `json:"chatCounter, omitempty"`
	StatusOnline   bool   `json:"statusOnline"`
	PrivateName    string `json:"privateName, omitempty"`
	PrivateSurname string `json:"privateSurname, omitempty"`
	PrivateUrl     string `json:"privateUrl, omitempty"`
	ChatName       string `json:"chatName, omitempty"`
	ChatPhoto      string `json:"chatPhoto, omitempty"`
}
