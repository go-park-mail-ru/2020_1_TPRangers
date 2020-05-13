package models

type Group struct {
	Name        string  `json:"name, omitempty"`
	About      *string  `json:"about, omitempty"`
	PhotoUrl   *string  `json:"photoUrl, omitempty"`
}

