package models


type SavePhotoResponse struct {
	Message string `json:"message"`
	Filename string `json:"filename, omitempty"`
}


