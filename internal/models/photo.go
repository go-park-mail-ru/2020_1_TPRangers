package models

type Photo struct {
	Url     string `json:"url,omitempty"`
	Likes   int    `json:"likes"`
	WasLike bool   `json:"wasLike"`
}
