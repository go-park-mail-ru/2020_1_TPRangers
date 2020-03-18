package main

import (
	"./delivery"
	"./delivery/auth"
	"./delivery/feed"
	"./delivery/register"
	"./delivery/settings"
	"./delivery/user"
	"./middleware"
	"github.com/labstack/echo"
)

type RequestHandlers struct {
	regHandler     delivery.RegisterDelivery
	loginHandler   delivery.AuthDelivery
	settingHandler delivery.SettingsDelivery
	userHandler    delivery.UserDelivery
	feedHandler    delivery.FeedDelivery
}

func NewRequestHandler() *RequestHandlers {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	//
	logH := auth.NewLoginDelivery()
	regH := register.NewRegisterDelivery()
	userH := user.NewUserDelivery()
	setH := settings.NewSettingsDelivery()
	feedH := feed.NewFeedDelivery()

	api := &(RequestHandlers{
		regHandler:     regH,
		loginHandler:   logH,
		settingHandler: setH,
		userHandler:    userH,
		feedHandler:    feedH,
	})

	return api
}

func main() {

	server := echo.New()

	server.Use(middleware.PanicMiddleWare)
	server.Use(middleware.SetCorsMiddleware)

	api := NewRequestHandler()

	server.POST("/api/v1/auth", api.loginHandler.Login)          // //
	server.POST("/api/v1/registration", api.regHandler.Register) // //

	server.PUT("/api/v1/settings", api.settingHandler.UploadSettings) // //

	server.GET("/api/v1/news", api.feedHandler.Feed)               // //
	server.GET("/api/v1/profile", api.userHandler.Profile)         //
	server.GET("/api/v1/settings", api.settingHandler.GetSettings) // //
	server.GET("/api/v1/user/:id", api.userHandler.GetUser)        // //

	server.DELETE("/api/v1/auth", api.loginHandler.Logout) // //

	server.Logger.Fatal(server.Start(":3001"))
}
