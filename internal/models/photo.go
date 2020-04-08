package models

type Photo struct {
	Id  	*int  	`json:"Id,omitempty"`
	Url     *string `json:"url,omitempty"`
	Likes   *int    `json:"likes"`
	WasLike bool   `json:"wasLike"`
}
