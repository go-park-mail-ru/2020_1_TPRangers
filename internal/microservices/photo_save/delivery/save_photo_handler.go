package delivery

import (
	"github.com/labstack/echo"
	"main/internal/microservices/photo_save"
	"main/models"
	"net/http"
	"os"
)

type PhotoSaveDeliveryRealisation struct {
	PhotoSaveLogic photo_save.PhotoSaveUseCase
}

func (PhotoSaveD PhotoSaveDeliveryRealisation) Upload (c echo.Context) error  {
	file, err := c.FormFile("fileData")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	data := models.PhotoInfo{File: file, Src: src}

	filename, err := PhotoSaveD.PhotoSaveLogic.PhotoSave(data)
	if err != nil {
		return err
	}

	response := new(models.SavePhotoResponse)

	path := os.Getenv("UPLOAD_PATH")
	response.Message = "Файл загружен"
	response.Filename = path + filename

	return c.JSON(http.StatusOK, response)
}

func NewSavePhotoDeliveryRealisation(logic photo_save.PhotoSaveUseCase) PhotoSaveDeliveryRealisation {
	return PhotoSaveDeliveryRealisation{PhotoSaveLogic: logic}
}

func (PhotoSaveD PhotoSaveDeliveryRealisation) InitHandler(server *echo.Echo)  {
	server.POST("/upload", PhotoSaveD.Upload)
}
