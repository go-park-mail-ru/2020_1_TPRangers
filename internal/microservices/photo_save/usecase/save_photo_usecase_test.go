package usecase

import (
	"main/internal/models"
	"main/internal/tools/errors"
	"mime/multipart"
	"testing"
)

func TestSavePhotoUseCaseRealisation_PhotoSave(t *testing.T) {

	saverTest := SavePhotoUseCaseRealisation{}

	filenames := []string{"haha.html", "nonfilehaha" , "iamfrontend.js" }

	for _ , filename := range filenames {

		if _ , err := saverTest.PhotoSave(models.PhotoInfo{
			File: &multipart.FileHeader{
				Filename: filename,
			},
		}) ; err != errors.InvalidPhotoFormat {
			t.Error("invalid behaviour : " , err )
		}

	}

}


