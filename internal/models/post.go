package models

type Post struct {
	Text        string `json:"text,omitempty"`
	Photo       Photo  `json:"photo,omitempty"`
	Attachments string `json:"attachments,omitempty"`
	Likes       int    `json:"likes,omitempty"`
	WasLike     bool   `json:"wasLike,omitempty"`
}