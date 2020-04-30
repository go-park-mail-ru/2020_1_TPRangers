package usecase

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"main/internal/models"
	"os"
	"path/filepath"
	"time"
)

type SavePhotoUseCaseRealisation struct {
}

func (SavePhotoUseCaseRealisation) PhotoSave (info models.PhotoInfo) (string, error) {
	hash := md5.Sum([]byte(time.Now().String() + info.File.Filename))
	filename := hex.EncodeToString(hash[:])+ filepath.Ext(info.File.Filename)
	path := os.Getenv("FILEPATH")
	dst, err := os.Create(path + filename)
	if err != nil {
		return  "", err
	}
	defer dst.Close()


	if _, err = io.Copy(dst, info.Src); err != nil {
		return "", err
	}
	return filename, nil
}

func NewUserUseCaseRealisation() SavePhotoUseCaseRealisation {
	return SavePhotoUseCaseRealisation{}
}