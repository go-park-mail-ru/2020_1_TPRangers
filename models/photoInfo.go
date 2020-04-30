package models

import "mime/multipart"

type PhotoInfo struct {
	File *multipart.FileHeader
	Src  multipart.File
}
