package models

type Feed struct {
	Posts []Post `json:"posts,omitempty"`
}