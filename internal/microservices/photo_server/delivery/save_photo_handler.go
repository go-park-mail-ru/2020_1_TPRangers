package delivery

import (
	"github.com/labstack/echo"
	"main/internal/microservices/photo_server"
	"main/internal/models"
	"net/http"
	"os"
)

type PhotoSaveDeliveryRealisation struct {
	PhotoSaveLogic photo_server.PhotoSaveUseCase
}

func (PhotoSaveD PhotoSaveDeliveryRealisation) upload (c echo.Context) error  {
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

	path := os.Getenv("FILEPATH")
	response.Message = "Файл загружен"
	response.Filename = "/" + path + filename

	return c.JSON(http.StatusOK, response)
}

func NewSavePhotoDeliveryRealisation(logic photo_server.PhotoSaveUseCase) PhotoSaveDeliveryRealisation {
	return PhotoSaveDeliveryRealisation{PhotoSaveLogic: logic}
}

func (PhotoSaveD PhotoSaveDeliveryRealisation) InitHandler(server *echo.Echo)  {
	server.POST("/upload", PhotoSaveD.upload)
}
