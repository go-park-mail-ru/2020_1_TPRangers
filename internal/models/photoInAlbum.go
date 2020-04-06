package models

type PhotoInAlbum struct {
	Url 		string `json:"url, omitempty"`
	AlbumID 	string `json:"id, omitempty"`
}