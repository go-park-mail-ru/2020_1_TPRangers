package models

type OtherUserProfileData struct {
	Feed      []Post              `json:"feed,omitempty"`
	User      Settings            `json:"user,omitempty"`
	Friends   []FriendLandingInfo `json:"friends,omitempty"`
	IsFriends bool                `json:"isFriends"`
	IsMe      bool                `json:"isMe"`
}
