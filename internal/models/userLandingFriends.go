package models

type FriendLandingInfo struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Photo   string `json:"avatar,omitempty"`
	Login   string `json:"url,omitempty"`
}
