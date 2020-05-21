package models

type Sticker struct {
	Name   *string `json:"name,omitempty"`
	Phrase *string `json:"phrase,omitempty"`
	Link   *string `json:"link,omitempty"`
}
