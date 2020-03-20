package models

type Photo struct {
	Url     string `json:"url,omitempty"`
	Likes   int    `json:"likes,omitempty"`
	WasLike bool   `json:"wasLike,omitempty"`
}


