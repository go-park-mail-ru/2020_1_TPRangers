package usecase

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"main/internal/models"
	"main/internal/tools/errors"
	"os"
	"path/filepath"
	"time"
)

type SavePhotoUseCaseRealisation struct {
	uploader *s3manager.Uploader
}

func (SPH SavePhotoUseCaseRealisation) PhotoSave(info models.PhotoInfo) (string, error) {
	hash := md5.Sum([]byte(time.Now().String() + info.File.Filename))

	ext := filepath.Ext(info.File.Filename)

	if ext != ".png" && ext != ".jpg" && ext != ".svg" && ext != ".bmp" && ext != ".gif"{
		return "", errors.InvalidPhotoFormat
	}

	filename := hex.EncodeToString(hash[:]) + ext
	path := os.Getenv("FILEPATH")


	_, err := SPH.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("tpvk"),
		Key:    aws.String(path + filename),
		ACL:    aws.String("public-read"),
		Body:   info.Src,
		ServerSideEncryption: aws.String("AES256"),
	})

	return filename, err
}

func (SPH SavePhotoUseCaseRealisation) AttachSave(info models.PhotoInfo) (string, error) {
	hash := md5.Sum([]byte(time.Now().String() + info.File.Filename))

	ext := filepath.Ext(info.File.Filename)
	filename := hex.EncodeToString(hash[:]) + ext
	path := os.Getenv("ATTACHPATH")

	_, err := SPH.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("tpvk"),
		Key:    aws.String(path + filename),
		ACL:    aws.String("public-read"),
		Body:   info.Src,
		ServerSideEncryption: aws.String("AES256"),
	})

	return filename, err
}

func NewUserUseCaseRealisation() SavePhotoUseCaseRealisation {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://hb.bizmrg.com"),
		Region:   aws.String("ru-msk"),
	}))
	return SavePhotoUseCaseRealisation{uploader: s3manager.NewUploader(sess)}
}
