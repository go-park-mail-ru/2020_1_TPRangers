package models

type GroupProfile struct {
	GroupID    int                 `json:"id,omitempty"`
	Owner      FriendLandingInfo   `json:"owner,omitempty"`
	GroupInfo  Group               `json:"groupInfo,omitempty"`
	Members    []FriendLandingInfo `json:"members,omitempty"`
	IsJoined   bool                `json:"isJoined"`
}