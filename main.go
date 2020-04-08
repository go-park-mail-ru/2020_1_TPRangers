package main

import (
	"database/sql"
	"github.com/labstack/echo"
	middleware2 "github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	repositoryCookie "main/internal/cookies/repository"
	deliveryFeed "main/internal/feeds/delivery"
	repositoryFeed "main/internal/feeds/repository"
	usecaseFeed "main/internal/feeds/usecase"
	"main/internal/middleware"
	deliveryUser "main/internal/users/delivery"
	repositoryUser "main/internal/users/repository"
	usecaseUser "main/internal/users/usecase"

	repositoryLikes "main/internal/like/repository"
	usecaseLikes "main/internal/like/usecase"
	deliveryLikes "main/internal/like/delivery"
)

const (
	usernameDB = "postgres"
	passwordDB = "postgres"
	nameDB     = "vk"
	redisPas   = ""
	redisPort  = "127.0.0.1:6379"
)

type RequestHandlers struct {
	userHandler deliveryUser.UserDeliveryRealisation
	feedHandler deliveryFeed.FeedDeliveryRealisation
	likeHandler deliveryLikes.LikeDelivery
}

func NewRequestHandler(db *sql.DB, logger *zap.SugaredLogger) *RequestHandlers {



	sessionDB := repositoryCookie.NewCookieRepositoryRealisation(redisPort, redisPas)
	feedDB := repositoryFeed.NewFeedRepositoryRealisation(db)
	userDB := repositoryUser.NewUserRepositoryRealisation(db)

	likesDB := repositoryLikes.NewLikeRepositoryRealisation(db)
	likesUse := usecaseLikes.NewLikeUseRealisation(likesDB,sessionDB)
	likeH := deliveryLikes.NewLikeDelivery(logger , likesUse)

	feedUseCase := usecaseFeed.NewFeedUseCaseRealisation(feedDB, sessionDB)
	userUseCase := usecaseUser.NewUserUseCaseRealisation(userDB, feedDB, sessionDB)

	userH := deliveryUser.NewUserDelivery(logger, userUseCase)
	feedH := deliveryFeed.NewFeedDelivery(logger, feedUseCase)

	api := &(RequestHandlers{
		userHandler: userH,
		feedHandler: feedH,
		likeHandler: likeH,
	})

	return api
}

func main() {

	server := echo.New()
	server.Use(middleware2.CSRFWithConfig(middleware2.CSRFConfig{
		TokenLength:    32,
		TokenLookup:    "header" + echo.HeaderXCSRFToken,
		ContextKey:     "csrf",
	}))

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
	api.likeHandler.InitHandlers(server)

	server.Logger.Fatal(server.Start(":3001"))
	//server.Logger.Fatal(server.StartTLS(":3001","./internal/tools/ssl/bundle.pem","./internal/tools/ssl/private.key"))
}
