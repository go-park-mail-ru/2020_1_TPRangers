package models

type MainUserProfileData struct {
	Feed    []Post              `json:"feed,omitempty"`
	User    Settings            `json:"user,omitempty"`
	Friends []FriendLandingInfo `json:"friends,omitempty"`
}
