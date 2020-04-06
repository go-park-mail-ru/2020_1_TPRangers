package models

type Post struct {
	Id  		int  	`json:"Id,omitempty"`
	Text        string `json:"text,omitempty"`
	Photo       Photo  `json:"photo,omitempty"`
	Attachments string `json:"attachments,omitempty"`
	Likes       int    `json:"likes"`
	WasLike     bool   `json:"wasLike"`
	Creation	string `json:"date,omitempty"`
}