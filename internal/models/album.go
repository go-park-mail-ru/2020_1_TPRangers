package models

type Album struct {
	ID			string `json:"id, omitempty"`
	Name 		string `json:"name, omitempty"`
	PhotoUrl 	*string `json:"photo_url, omitempty"`

}

