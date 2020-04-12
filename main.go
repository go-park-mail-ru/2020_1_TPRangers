package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	repositoryCookie "main/internal/cookies/repository"
	repositoryPhoto "main/internal/photos/repository"
	repositoryAlbum "main/internal/albums/repository"
	deliveryFeed "main/internal/feeds/delivery"
	deliveryAlbum "main/internal/albums/delivery"
	repositoryFeed "main/internal/feeds/repository"
	usecaseFeed "main/internal/feeds/usecase"
	"main/internal/middleware"
	deliveryUser "main/internal/users/delivery"
	repositoryUser "main/internal/users/repository"
	usecaseUser "main/internal/users/usecase"
	deliveryPhoto "main/internal/photos/delivery"

	"os"


	deliveryLikes "main/internal/like/delivery"
	repositoryLikes "main/internal/like/repository"
	usecaseLikes "main/internal/like/usecase"
	usecasePhoto "main/internal/photos/usecase"
	usecaseAlbum "main/internal/albums/usecase"
	deliveryFriends "main/internal/friends/delivery"
	repositoryFriends "main/internal/friends/repository"
	usecaseFriends "main/internal/friends/usecase"
)

type RequestHandlers struct {

	userHandler deliveryUser.UserDeliveryRealisation
	feedHandler deliveryFeed.FeedDeliveryRealisation
	likeHandler deliveryLikes.LikeDelivery
	photoHandler deliveryPhoto.PhotoDeliveryRealisation
	albumHandler deliveryAlbum.AlbumDeliveryRealisation
	friendHandler deliveryFriends.FriendDeliveryRealisation
}

func InitializeDataBases(server *echo.Echo) (*sql.DB, repositoryCookie.CookieRepositoryRealisation) {
	err := godotenv.Load("project.env")
	if err != nil {
		server.Logger.Fatal("can't load .env file :", err.Error())
	}
	usernameDB := os.Getenv("POSTGRES_USERNAME")
	passwordDB := os.Getenv("POSTGRES_PASSWORD")
	nameDB := os.Getenv("POSTGRES_NAME")

	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		server.Logger.Fatal("NO CONNECTION TO BD", err.Error())
	}

	redisPas := os.Getenv("REDIS_PASSWORD")
	redisPort := os.Getenv("REDIS_PORT")

	sessionDB := repositoryCookie.NewCookieRepositoryRealisation(redisPort, redisPas)

	return db, sessionDB
}

func NewRequestHandler(db *sql.DB, sessionDB repositoryCookie.CookieRepositoryRealisation, logger *zap.SugaredLogger) *RequestHandlers {

	feedDB := repositoryFeed.NewFeedRepositoryRealisation(db)
	userDB := repositoryUser.NewUserRepositoryRealisation(db)

	photoDB := repositoryPhoto.NewPhotoRepositoryRealisation(db)
	likesDB := repositoryLikes.NewLikeRepositoryRealisation(db)
	albumDB := repositoryAlbum.NewAlbumRepositoryRealisation(db)




	friendsDB := repositoryFriends.NewFriendRepositoryRealisation(db)




	photoUseCase := usecasePhoto.NewPhotoUseCaseRealisation(photoDB, sessionDB)
	albumUseCase := usecaseAlbum.NewAlbumUseCaseRealisation(albumDB, sessionDB)


	feedUseCase := usecaseFeed.NewFeedUseCaseRealisation(feedDB)
	userUseCase := usecaseUser.NewUserUseCaseRealisation(userDB, friendsDB, feedDB, sessionDB)
	likesUse := usecaseLikes.NewLikeUseRealisation(likesDB)
	friendsUse := usecaseFriends.NewFriendUseCaseRealisation(friendsDB)

	likeH := deliveryLikes.NewLikeDelivery(logger, likesUse)

	userH := deliveryUser.NewUserDelivery(logger, userUseCase)
	feedH := deliveryFeed.NewFeedDelivery(logger, feedUseCase)

	photoH := deliveryPhoto.NewPhotoDelivery(logger, photoUseCase)
	albumH := deliveryAlbum.NewAlbumDelivery(logger, albumUseCase)

	friendH := deliveryFriends.NewUserDelivery(logger, friendsUse)


	api := &(RequestHandlers{


		photoHandler: photoH,
		albumHandler: albumH,
		userHandler:   userH,
		feedHandler:   feedH,
		likeHandler:   likeH,
		friendHandler: friendH,

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

	db, sessions := InitializeDataBases(server)
	defer db.Close()

	midHandler := middleware.NewMiddlewareHandler(logger, sessions)
	midHandler.SetMiddleware(server)

	api := NewRequestHandler(db, sessions, logger)

	api.userHandler.InitHandlers(server)
	api.feedHandler.InitHandlers(server)
	api.likeHandler.InitHandlers(server)
	api.photoHandler.InitHandlers(server)
	api.albumHandler.InitHandlers(server)
	api.friendHandler.InitHandlers(server)

	port := os.Getenv("PORT")

	server.Logger.Fatal(server.Start(port))
}
