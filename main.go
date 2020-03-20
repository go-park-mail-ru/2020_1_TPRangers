package main

import (
	"./delivery"
	deliveryAuth "./delivery/auth"
	deliveryFeed "./delivery/feed"
	deliveryRegister "./delivery/register"
	deliverySettings "./delivery/settings"
	deliveryUser "./delivery/user"

	usecaseAuth "./usecase/auth"
	usecaseFeed "./usecase/feed"
	usecaseRegister "./usecase/register"
	usecaseSettings "./usecase/settings"
	usecaseUser "./usecase/user"

	repositoryAuth "./repository/auth"
	repositoryCookie "./repository/cookie"
	repositoryFeed "./repository/feed"
	repositoryRegister "./repository/register"
	repositoryUser "./repository/user"

	"./middleware"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	usernameDB = "postgres"
	passwordDB = "nikita2003"
	nameDB     = "vk"
	redisPas   = ""
	redisPort  = "127.0.0.1:6379"
)

type RequestHandlers struct {
	regHandler     delivery.RegisterDelivery
	authHandler    delivery.AuthDelivery
	settingHandler delivery.SettingsDelivery
	userHandler    delivery.UserDelivery
	feedHandler    delivery.FeedDelivery
}

func NewRequestHandler() *RequestHandlers {

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	sessionDB := repositoryCookie.NewCookieRepositoryRealisation(redisPort, redisPas)
	authDB, err := repositoryAuth.NewAuthRepositoryRealisation(usernameDB, passwordDB, nameDB)

	if err != nil {
		logger.Debug(
			zap.String("AUTH DB STARTING ERROR", err.Error()),
		)
	}

	registerDB, err := repositoryRegister.NewRegisterRepositoryRealisation(usernameDB, passwordDB, nameDB)

	if err != nil {
		logger.Debug(
			zap.String("REGISTER DB STARTING ERROR", err.Error()),
		)
	}

	feedDB, err := repositoryFeed.NewFeedRepositoryRealisation(usernameDB, passwordDB, nameDB)

	if err != nil {
		logger.Debug(
			zap.String("FEED DB STARTING ERROR", err.Error()),
		)
	}

	userDB, err := repositoryUser.NewUserRepositoryRealisation(usernameDB, passwordDB, nameDB)

	if err != nil {
		logger.Debug(
			zap.String("USER DB STARTING ERROR", err.Error()),
		)
	}

	authUseCase := usecaseAuth.NewAuthUseCaseRealisation(authDB, sessionDB, logger)
	registerUseCase := usecaseRegister.NewRegisterUseCaseRealisation(registerDB, sessionDB, logger)
	feedUseCase := usecaseFeed.NewFeedUseCaseRealisation(feedDB, sessionDB, logger)
	settingsUseCase := usecaseSettings.NewSetUseCaseRealisation(userDB, sessionDB, logger)
	userUseCase := usecaseUser.NewUserUseCaseRealisation(userDB, feedDB, sessionDB, logger)

	authH := deliveryAuth.NewLoginDelivery(logger, authUseCase)
	regH := deliveryRegister.NewRegisterDelivery(logger, registerUseCase)
	userH := deliveryUser.NewUserDelivery(logger, userUseCase)
	setH := deliverySettings.NewSettingsDelivery(logger, settingsUseCase)
	feedH := deliveryFeed.NewFeedDelivery(logger, feedUseCase)

	api := &(RequestHandlers{
		regHandler:     regH,
		authHandler:    authH,
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

	server.POST("/api/v1/auth", api.authHandler.Login)           // //
	server.POST("/api/v1/registration", api.regHandler.Register) // //

	server.PUT("/api/v1/settings", api.settingHandler.UploadSettings) // //

	server.GET("/api/v1/news", api.feedHandler.Feed)               // //
	server.GET("/api/v1/profile", api.userHandler.Profile)         //
	server.GET("/api/v1/settings", api.settingHandler.GetSettings) // //
	server.GET("/api/v1/user/:id", api.userHandler.GetUser)        // //

	server.DELETE("/api/v1/auth", api.authHandler.Logout) // //

	server.Logger.Fatal(server.Start(":3001"))
}
