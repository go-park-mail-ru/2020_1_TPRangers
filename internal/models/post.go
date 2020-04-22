package models

type Post struct {
	Id            int    `json:"id,omitempty"`
	Text          string `json:"text,omitempty"`
	Photo         Photo  `json:"photo,omitempty"`
	Attachments   string `json:"attachments,omitempty"`
	Likes         int    `json:"likes"`
	WasLike       bool   `json:"wasLike"`
	Creation      string `json:"date,omitempty"`
	AuthorName    string `json:"authorName,omitempty"`
	AuthorSurname string `json:"authorSurname,omitempty"`
	AuthorUrl     string `json:"authorUrl,omitempty"`
	AuthorPhoto   string `json:"authorPhoto,omitempty"`
	Comments	  []Comment `json:"comments,omitempty"`
}
