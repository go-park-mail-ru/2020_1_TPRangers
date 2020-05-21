package models

type Group struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	About      *string  `json:"about,omitempty"`
	PhotoUrl   *string  `json:"photoUrl,omitempty"`
}

