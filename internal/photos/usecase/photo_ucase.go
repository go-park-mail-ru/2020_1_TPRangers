package usecase

import (
	Sess "main/internal/cookies"
	SessRep "main/internal/cookies/repository"
	"main/internal/models"
	photoRep "main/internal/photos/repository"
	"main/internal/tools/errors"
	"main/internal/photos"
)

type PhotoUseCaseRealisation struct {
	photoDB    photos.PhotoRepository
	sessionDB Sess.CookieRepository
}

func (photoR PhotoUseCaseRealisation) GetPhotosFromAlbum(cookie string, albumID int) (models.Photos, error) {
	_, err := photoR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return  models.Photos{} ,errors.InvalidCookie
	}

	photos, err := photoR.photoDB.GetPhotosFromAlbum(albumID)

	return photos, nil
}

func (photoR PhotoUseCaseRealisation) UploadPhotoToAlbum(cookie string, photoData models.PhotoInAlbum) error {
	_, err := photoR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return errors.InvalidCookie
	}

	err = photoR.photoDB.UploadPhotoToAlbum(photoData)

	if err != nil {
		return errors.FailReadFromDB
	}

	return nil
}


func NewPhotoUseCaseRealisation(photoDB photoRep.PhotoRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation) PhotoUseCaseRealisation {
	return PhotoUseCaseRealisation{
		photoDB:    photoDB,
		sessionDB: sesDB,
	}
}