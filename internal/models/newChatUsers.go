package models


type NewChatUsers struct {
	ChatPhoto string `json:"chatPhoto,omitempty"`
	ChatName string `json:"chatName,omitempty"`
	UsersLogin []string `json:"usersLogin,omitempty"`
}