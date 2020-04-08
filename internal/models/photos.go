package models


type Photos struct {
	AlbumName 	string `json:"album_name,omitempty"`
	Urls     	[]string `json:"url,omitempty"`
}
