package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	metrics "main/internal/metrics/delivery"
	authorMicro "main/internal/microservices/authorization/delivery"
	deliveryPhotoSave "main/internal/microservices/photo_save/delivery"
	usecasePhotoSave "main/internal/microservices/photo_save/usecase"
	mw2 "main/internal/middleware"
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

func LoadMicroservices(server *echo.Echo) (authorMicro.SessionCheckerClient, *grpc.ClientConn) {

	authPORT := os.Getenv("AUTHORIZ_PORT")

	grpcConn, err := grpc.Dial(
		"127.0.0.1"+authPORT,
		grpc.WithInsecure(),
	)
	if err != nil {
		server.Logger.Fatal("cant connect to grpc")
	}

	authManager := authorMicro.NewSessionCheckerClient(grpcConn)

	return authManager, grpcConn

}

func main() {
	err := godotenv.Load("photo_save_micro.env")
	if err != nil {
		return
	}

	server := echo.New()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	auth, authConn := LoadMicroservices(server)

	defer func() {
		if authConn != nil {
			authConn.Close()
		}
	}()

	tracker := metrics.RegisterMetrics(server)

	midHandler := mw2.NewMiddlewareHandler(logger, auth, tracker, "https://social-hub.ru")
	midHandler.SetMiddleware(server)

	//server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"http://localhost:3000", "https://social-hub.ru"},
	//	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	//}))
	//server.Use(middleware.Logger())
	//server.Use(middleware.Recover())

	api := NewRequestHandlers()
	api.photoSaveHandler.InitHandler(server)

	port := os.Getenv("PORT_SAVE")
	server.Logger.Fatal(server.Start(port))
}