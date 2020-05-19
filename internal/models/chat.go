package models

type Chat struct {
	IsGroupChat            bool   `json:"isGroupChat"`
	ChatId                 string `json:"chatId,omitempty"`
	ChatPhoto              string `json:"chatPhoto,omitempty"`
	ChatName               string `json:"chatName,omitempty"`
	PrivateName            string `json:"privateName,omitempty"`
	PrivateSurname         string `json:"privateSurname,omitempty"`
	PrivateUrl             string `json:"privateUrl,omitempty"`
	StatusOnline           bool   `json:"statusOnline"`
	LastMessageAuthorPhoto string `json:"lastMessageAuthorPhoto,omitempty"`
	LastMessageTime        string `json:"lastMessageTime,omitempty"`
	LastMessageTxt         string `json:"lastMessageTxt,omitempty"`
}
