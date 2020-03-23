package main

import (
	deliveryFeed "./internal/feeds/delivery"
	deliveryUser "./internal/users/delivery"
	"./internal/middleware"
	repositoryCookie "./internal/cookies/repository"
	repositoryFeed "./internal/feeds/repository"
	repositoryUser "./internal/users/repository"
	usecaseFeed "./internal/feeds/usecase"
	usecaseUser "./internal/users/usecase"
	"database/sql"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	usernameDB = "alexandr"
	passwordDB = "nikita2003"
	nameDB     = "VK"
	redisPas   = ""
	redisPort  = "127.0.0.1:6379"
)

type RequestHandlers struct {
	userHandler deliveryUser.UserDeliveryRealisation
	feedHandler deliveryFeed.FeedDeliveryRealisation
}

func NewRequestHandler(db *sql.DB, logger *zap.SugaredLogger) *RequestHandlers {

	sessionDB := repositoryCookie.NewCookieRepositoryRealisation(redisPort, redisPas)
	feedDB := repositoryFeed.NewFeedRepositoryRealisation(db)
	userDB := repositoryUser.NewUserRepositoryRealisation(db)

	feedUseCase := usecaseFeed.NewFeedUseCaseRealisation(feedDB, sessionDB)
	userUseCase := usecaseUser.NewUserUseCaseRealisation(userDB, feedDB, sessionDB)

	userH := deliveryUser.NewUserDelivery(logger, userUseCase)
	feedH := deliveryFeed.NewFeedDelivery(logger, feedUseCase)

	api := &(RequestHandlers{
		userHandler: userH,
		feedHandler: feedH,
	})

	return api
}

func main() {

	server := echo.New()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	server.Use(middleware.PanicMiddleWare)
	server.Use(middleware.SetCorsMiddleware)

	logFunc := middleware.AccessLog(logger)

	server.Use(logFunc)

	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	defer db.Close()
	if err != nil {
		server.Logger.Fatal("NO CONNECTION TO BD", err.Error())
	}

	api := NewRequestHandler(db, logger)

	api.userHandler.InitHandlers(server)
	api.feedHandler.InitHandlers(server)

	server.Logger.Fatal(server.Start(":3001"))
}
