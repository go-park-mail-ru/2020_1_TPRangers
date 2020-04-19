package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	repositoryMessage "main/internal/message/repository"
	usecaseMessage "main/internal/socket/usecase"
	deliveryMessage "main/internal/socket/delivery"

	repositoryChat 	"main/internal/chats/repository"
	deliveryChat 	"main/internal/chats/delivery"
	usecaseChat 	"main/internal/chats/usecase"


	deliveryAlbum "main/internal/albums/delivery"
	repositoryAlbum "main/internal/albums/repository"
	repositoryCookie "main/internal/cookies/repository"
	deliveryFeed "main/internal/feeds/delivery"
	repositoryFeed "main/internal/feeds/repository"
	usecaseFeed "main/internal/feeds/usecase"
	"main/internal/middleware"
	deliveryPhoto "main/internal/photos/delivery"
	repositoryPhoto "main/internal/photos/repository"
	deliveryUser "main/internal/users/delivery"
	repositoryUser "main/internal/users/repository"
	usecaseUser "main/internal/users/usecase"

	"os"

	usecaseAlbum "main/internal/albums/usecase"
	deliveryFriends "main/internal/friends/delivery"
	repositoryFriends "main/internal/friends/repository"
	usecaseFriends "main/internal/friends/usecase"
	deliveryLikes "main/internal/like/delivery"
	repositoryLikes "main/internal/like/repository"
	usecaseLikes "main/internal/like/usecase"
	usecasePhoto "main/internal/photos/usecase"

)

type RequestHandlers struct {
	userHandler   deliveryUser.UserDeliveryRealisation
	feedHandler   deliveryFeed.FeedDeliveryRealisation
	likeHandler   deliveryLikes.LikeDelivery
	photoHandler  deliveryPhoto.PhotoDeliveryRealisation
	albumHandler  deliveryAlbum.AlbumDeliveryRealisation
	friendHandler deliveryFriends.FriendDeliveryRealisation
	messageHandler deliveryMessage.SocketDelivery
	chatHandler 	deliveryChat.ChatsDelivery

}

func InitializeDataBases(server *echo.Echo) (*sql.DB, repositoryCookie.CookieRepositoryRealisation , repositoryMessage.MessageRepositoryRealisation) {
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

	redisChatPort := os.Getenv("REDIS_CHAT_PORT")

	sessionDB := repositoryCookie.NewCookieRepositoryRealisation(redisPort, redisPas)
	chatDB := repositoryMessage.NewMessageRepositoryRealisation(redisChatPort , "" , db)

	return db, sessionDB , chatDB
}

func NewRequestHandler(db *sql.DB, sessionDB repositoryCookie.CookieRepositoryRealisation, messageDB repositoryMessage.MessageRepositoryRealisation,logger *zap.SugaredLogger) *RequestHandlers {

	feedDB := repositoryFeed.NewFeedRepositoryRealisation(db)
	userDB := repositoryUser.NewUserRepositoryRealisation(db)
	chatDB := repositoryChat.NewChatRepositoryRealisation(db)
	photoDB := repositoryPhoto.NewPhotoRepositoryRealisation(db)
	likesDB := repositoryLikes.NewLikeRepositoryRealisation(db)
	albumDB := repositoryAlbum.NewAlbumRepositoryRealisation(db)

	friendsDB := repositoryFriends.NewFriendRepositoryRealisation(db)

	photoUseCase := usecasePhoto.NewPhotoUseCaseRealisation(photoDB)
	albumUseCase := usecaseAlbum.NewAlbumUseCaseRealisation(albumDB)
	messageUseCase := usecaseMessage.NewSocketUseCaseRealisation(messageDB)
	feedUseCase := usecaseFeed.NewFeedUseCaseRealisation(feedDB)
	userUseCase := usecaseUser.NewUserUseCaseRealisation(userDB, friendsDB, feedDB, sessionDB)
	likesUse := usecaseLikes.NewLikeUseRealisation(likesDB)
	friendsUse := usecaseFriends.NewFriendUseCaseRealisation(friendsDB)
	chatUse := usecaseChat.NewChatUseCaseRealisation(chatDB, friendsDB)

	likeH := deliveryLikes.NewLikeDelivery(logger, likesUse)

	userH := deliveryUser.NewUserDelivery(logger, userUseCase)
	feedH := deliveryFeed.NewFeedDelivery(logger, feedUseCase)
	messageH := deliveryMessage.NewSocketDelivery(logger , messageUseCase)
	photoH := deliveryPhoto.NewPhotoDelivery(logger, photoUseCase)
	albumH := deliveryAlbum.NewAlbumDelivery(logger, albumUseCase)
	chatH := deliveryChat.NewChatsDelivery(logger , chatUse)

	friendH := deliveryFriends.NewFriendDelivery(logger, friendsUse)

	api := &(RequestHandlers{

		photoHandler:  photoH,
		albumHandler:  albumH,
		userHandler:   userH,
		feedHandler:   feedH,
		likeHandler:   likeH,
		friendHandler: friendH,
		messageHandler: messageH,
		chatHandler: chatH,
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

	db, sessions , messages := InitializeDataBases(server)
	defer db.Close()

	midHandler := middleware.NewMiddlewareHandler(logger, sessions)
	midHandler.SetMiddleware(server)

	api := NewRequestHandler(db, sessions, messages ,logger)

	api.userHandler.InitHandlers(server)
	api.feedHandler.InitHandlers(server)
	api.likeHandler.InitHandlers(server)
	api.photoHandler.InitHandlers(server)
	api.albumHandler.InitHandlers(server)
	api.friendHandler.InitHandlers(server)
	api.messageHandler.InitHandlers(server)
	api.chatHandler.InitHandlers(server)

	port := os.Getenv("PORT")

	server.Logger.Fatal(server.Start(port))
}
