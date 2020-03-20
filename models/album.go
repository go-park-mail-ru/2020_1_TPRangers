package models

type Album struct {
	Photos []Photo `json:"photos,omitempty"`
}
