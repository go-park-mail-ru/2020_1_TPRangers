package models

type Comment struct {
	CommentID     string `json:"comment_id,omitempty"`
	PostID        string `json:"post_id,omitempty"`
	Text          string `json:"text,omitempty"`
	Photo         Photo  `json:"photo"`
	Attachments   string `json:"attachments,omitempty"`
	Likes         int    `json:"likes, omitempty"`
	WasLike       bool   `json:"wasLike, omitempty"`
	Creation      string `json:"date,omitempty"`
	AuthorName    string `json:"authorName,omitempty"`
	AuthorSurname string `json:"authorSurname,omitempty"`
	AuthorUrl     string `json:"authorUrl,omitempty"`
	AuthorPhoto   string `json:"authorPhoto,omitempty"`
}
