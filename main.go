package main

import (
	"./delivery"
	deliveryAuth "./delivery/auth"
	deliveryFeed "./delivery/feed"
	deliveryRegister "./delivery/register"
	deliverySettings "./delivery/settings"
	deliveryUser "./delivery/user"
	deliveryFriend "./delivery/friend"
	"database/sql"

	usecaseAuth "./usecase/auth"
	usecaseFeed "./usecase/feed"
	usecaseRegister "./usecase/register"
	usecaseSettings "./usecase/settings"
	usecaseUser "./usecase/user"
	usecaseFriend "./usecase/friend"

	repositoryAuth "./repository/auth"
	repositoryCookie "./repository/cookie"
	repositoryFeed "./repository/feed"
	repositoryRegister "./repository/register"
	repositoryUser "./repository/user"
	repositoryFriend "./repository/friend"

	"./middleware"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	usernameDB = "postgres"
	passwordDB = "071299"
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
	friendHandler  delivery.FriendDelivery
}

func NewRequestHandler(db *sql.DB) *RequestHandlers {

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	sessionDB := repositoryCookie.NewCookieRepositoryRealisation(redisPort, redisPas)
	authDB := repositoryAuth.NewAuthRepositoryRealisation(db)
	registerDB := repositoryRegister.NewRegisterRepositoryRealisation(db)
	feedDB := repositoryFeed.NewFeedRepositoryRealisation(db)
	userDB := repositoryUser.NewUserRepositoryRealisation(db)
	friendDB := repositoryFriend.NewFriendRepositoryRealisation(db)


	authUseCase := usecaseAuth.NewAuthUseCaseRealisation(authDB, sessionDB, logger)
	registerUseCase := usecaseRegister.NewRegisterUseCaseRealisation(registerDB, sessionDB, logger)
	feedUseCase := usecaseFeed.NewFeedUseCaseRealisation(feedDB, sessionDB, logger)
	settingsUseCase := usecaseSettings.NewSetUseCaseRealisation(userDB, sessionDB, logger)
	userUseCase := usecaseUser.NewUserUseCaseRealisation(userDB, feedDB, sessionDB, logger)
	friendUseCase := usecaseFriend.NewFriendUseCaseRealisation(userDB,friendDB,sessionDB,logger)

	authH := deliveryAuth.NewLoginDelivery(logger, authUseCase)
	regH := deliveryRegister.NewRegisterDelivery(logger, registerUseCase)
	userH := deliveryUser.NewUserDelivery(logger, userUseCase)
	setH := deliverySettings.NewSettingsDelivery(logger, settingsUseCase)
	feedH := deliveryFeed.NewFeedDelivery(logger, feedUseCase)
	friendH := deliveryFriend.NewRegisterDelivery(logger,friendUseCase)

	api := &(RequestHandlers{
		regHandler:     regH,
		authHandler:    authH,
		settingHandler: setH,
		userHandler:    userH,
		feedHandler:    feedH,
		friendHandler:friendH,
	})

	return api
}

func main() {

	server := echo.New()

	server.Use(middleware.PanicMiddleWare)
	server.Use(middleware.SetCorsMiddleware)

	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	defer db.Close()
	if err != nil {
		server.Logger.Fatal("NO CONNECTION TO BD" , err.Error())
	}

	api := NewRequestHandler(db)

	server.POST("/api/v1/login", api.authHandler.Login)           // //
	server.POST("/api/v1/registration", api.regHandler.Register) // //

	server.PUT("/api/v1/settings", api.settingHandler.UploadSettings) // //
	server.PUT("/api/v1/user/:id", api.friendHandler.AddFriend)        // //

	server.GET("/api/v1/news", api.feedHandler.Feed)               // //
	server.GET("/api/v1/profile", api.userHandler.Profile)         //
	server.GET("/api/v1/settings", api.settingHandler.GetSettings) // //
	server.GET("/api/v1/user/:id", api.userHandler.GetUser)        // //

	server.DELETE("/api/v1/auth", api.authHandler.Logout) // //

	server.Logger.Fatal(server.Start(":3001"))
}
