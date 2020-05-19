package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	deliveryPhotoSave "main/internal/microservices/photo_save/delivery"
	usecasePhotoSave "main/internal/microservices/photo_save/usecase"
	"os"
)

type RequestHandlers struct {
	photoSaveHandler deliveryPhotoSave.PhotoSaveDeliveryRealisation
}

func NewRequestHandlers() *RequestHandlers {
	photoSaveUseCase := usecasePhotoSave.NewUserUseCaseRealisation()

	photoSaveDelivery := deliveryPhotoSave.NewSavePhotoDeliveryRealisation(photoSaveUseCase)

	api := &(RequestHandlers{photoSaveDelivery})
	return api
}

func main() {
	err := godotenv.Load("photo_save_micro.env")
	if err != nil {
		return
	}

	server := echo.New()

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://social-hub.ru"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	api := NewRequestHandlers()
	api.photoSaveHandler.InitHandler(server)

	port := os.Getenv("PORT")
	server.Logger.Fatal(server.Start(port))
}
