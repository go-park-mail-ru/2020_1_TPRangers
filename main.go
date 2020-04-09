package main

import (
	"database/sql"
	"github.com/labstack/echo"
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

	deliveryLikes "main/internal/like/delivery"
	repositoryLikes "main/internal/like/repository"
	usecaseLikes "main/internal/like/usecase"

	deliveryFriends "main/internal/friends/delivery"
	repositoryFriends "main/internal/friends/repository"
	usecaseFriends "main/internal/friends/usecase"
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
	friendHandler deliveryFriends.FriendDeliveryRealisation
}

func NewRequestHandler(db *sql.DB, logger *zap.SugaredLogger) *RequestHandlers {


	sessionDB := repositoryCookie.NewCookieRepositoryRealisation(redisPort, redisPas)
	feedDB := repositoryFeed.NewFeedRepositoryRealisation(db)
	userDB := repositoryUser.NewUserRepositoryRealisation(db)
	likesDB := repositoryLikes.NewLikeRepositoryRealisation(db)
	friendsDB := repositoryFriends.NewFriendRepositoryRealisation(db)



	feedUseCase := usecaseFeed.NewFeedUseCaseRealisation(feedDB, sessionDB)
	userUseCase := usecaseUser.NewUserUseCaseRealisation(userDB, friendsDB ,feedDB, sessionDB)
	likesUse := usecaseLikes.NewLikeUseRealisation(likesDB,sessionDB)
	friendsUse := usecaseFriends.NewFriendUseCaseRealisation(friendsDB, sessionDB)

	likeH := deliveryLikes.NewLikeDelivery(logger , likesUse)
	userH := deliveryUser.NewUserDelivery(logger, userUseCase)
	feedH := deliveryFeed.NewFeedDelivery(logger, feedUseCase)
	friendH := deliveryFriends.NewUserDelivery(logger, friendsUse)

	api := &(RequestHandlers{
		userHandler: userH,
		feedHandler: feedH,
		likeHandler: likeH,
		friendHandler: friendH,
	})

	return api
}

func main() {

	server := echo.New()
	//server.Use(middleware2.CSRF())

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	//server.Use(middleware.PanicMiddleWare)
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
	api.friendHandler.InitHandlers(server)

	server.Logger.Fatal(server.Start(":3001"))
	//server.Logger.Fatal(server.StartTLS(":3001","./internal/tools/ssl/bundle.pem","./internal/tools/ssl/private.key"))
}
