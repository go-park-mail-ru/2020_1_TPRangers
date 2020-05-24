package delivery

import (
	"fmt"
	"github.com/labstack/echo"
	"main/internal/microservices/photo_save"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
	"os"
)

type PhotoSaveDeliveryRealisation struct {
	PhotoSaveLogic photo_save.PhotoSaveUseCase
}


func (PhotoSaveD PhotoSaveDeliveryRealisation) UploadPhoto(c echo.Context) error {

	userId := c.Get("user_id").(int)

	if userId == -1 {
		return c.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

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

	fmt.Println(data)

	filename, err := PhotoSaveD.PhotoSaveLogic.PhotoSave(data)

	fmt.Println(err)

	if err == errors.InvalidPhotoFormat {
		return c.NoContent(http.StatusTeapot)
	}

	if err != nil {
		return c.JSON(http.StatusConflict , models.JsonStruct{
			Err:  err.Error(),
		})
	}

	response := new(models.SavePhotoResponse)

	path := os.Getenv("UPLOAD_PATH")
	response.Message = "Файл загружен"
	response.Filename = path + filename
	c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
	return c.JSON(http.StatusOK, response)
}

func (PhotoSaveD PhotoSaveDeliveryRealisation) UploadAttachments(c echo.Context) error {

	userId := c.Get("user_id").(int)

	if userId == -1 {
		return c.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

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

	fmt.Println(data)

	filename, err := PhotoSaveD.PhotoSaveLogic.AttachSave(data)

	fmt.Println(err)

	if err != nil {
		return c.JSON(http.StatusConflict , models.JsonStruct{
			Err:  err.Error(),
		})
	}

	response := new(models.SavePhotoResponse)

	path := os.Getenv("UPLOAD_ATTACHPATH")
	response.Message = "Файл загружен"
	response.Filename = path + filename
	c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
	return c.JSON(http.StatusOK, response)
}

func NewSavePhotoDeliveryRealisation(logic photo_save.PhotoSaveUseCase) PhotoSaveDeliveryRealisation {
	return PhotoSaveDeliveryRealisation{PhotoSaveLogic: logic}
}

func (PhotoSaveD PhotoSaveDeliveryRealisation) InitHandler(server *echo.Echo) {
	server.POST("api/v1/upload", PhotoSaveD.UploadPhoto)
	server.POST("api/v1/attach", PhotoSaveD.UploadAttachments)

}