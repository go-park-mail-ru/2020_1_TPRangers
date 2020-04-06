package models

type Album struct {
	Url 		*string `json:"url, omitempty"`
	Name 		*string `json:"name, omitempty"`
	PhotoUrl 	*string `json:"photo_url, omitempty"`

}

