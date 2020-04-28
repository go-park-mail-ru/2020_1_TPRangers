package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"

	deliveryAlbum "main/internal/albums/delivery"
	repositoryAlbum "main/internal/albums/repository"
	repositoryChat "main/internal/chats/repository"
	usecaseChat "main/internal/chats/usecase"
	deliveryFeed "main/internal/feeds/delivery"
	repositoryFeed "main/internal/feeds/repository"
	usecaseFeed "main/internal/feeds/usecase"
	"main/internal/middleware"
	deliveryPhoto "main/internal/photos/delivery"
	deliveryUser "main/internal/users/delivery"
	repositoryUser "main/internal/users/repository"
	usecaseUser "main/internal/users/usecase"

	"os"

	usecaseAlbum "main/internal/albums/usecase"
	deliveryFriends "main/internal/friends/delivery"
	repositoryFriends "main/internal/friends/repository"
	usecaseFriends "main/internal/friends/usecase"
	deliveryLikes "main/internal/like/delivery"
	usecaseLikes "main/internal/like/usecase"
	usecasePhoto "main/internal/photos/usecase"

	authorMicro "main/internal/microservices/authorization/delivery"
	likeMicro "main/internal/microservices/likes/delivery"
	photoMicro "main/internal/microservices/photos/delivery"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type RequestHandlers struct {
	userHandler   deliveryUser.UserDeliveryRealisation
	feedHandler   deliveryFeed.FeedDeliveryRealisation
	likeHandler   deliveryLikes.LikeDelivery
	photoHandler  deliveryPhoto.PhotoDeliveryRealisation
	albumHandler  deliveryAlbum.AlbumDeliveryRealisation
	friendHandler deliveryFriends.FriendDeliveryRealisation
}

func InitializeDataBases(server *echo.Echo) *sql.DB {
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

	return db
}

func NewRequestHandler(db *sql.DB, session authorMicro.SessionCheckerClient, likes likeMicro.LikeCheckerClient, photos photoMicro.PhotoCheckerClient, logger *zap.SugaredLogger) *RequestHandlers {

	feedDB := repositoryFeed.NewFeedRepositoryRealisation(db)
	userDB := repositoryUser.NewUserRepositoryRealisation(db)
	chatDB := repositoryChat.NewChatRepositoryRealisation(db)
	albumDB := repositoryAlbum.NewAlbumRepositoryRealisation(db)
	friendsDB := repositoryFriends.NewFriendRepositoryRealisation(db)

	photoUseCase := usecasePhoto.NewPhotoUseCaseRealisation(photos)
	albumUseCase := usecaseAlbum.NewAlbumUseCaseRealisation(albumDB)
	feedUseCase := usecaseFeed.NewFeedUseCaseRealisation(feedDB)
	userUseCase := usecaseUser.NewUserUseCaseRealisation(userDB, friendsDB, feedDB, session)
	likesUse := usecaseLikes.NewLikeUseRealisation(likes)
	friendsUse := usecaseFriends.NewFriendUseCaseRealisation(friendsDB)
	chatUse := usecaseChat.NewChatUseCaseRealisation(chatDB, friendsDB)

	likeH := deliveryLikes.NewLikeDelivery(logger, likesUse)
	userH := deliveryUser.NewUserDelivery(logger, userUseCase)
	feedH := deliveryFeed.NewFeedDelivery(logger, feedUseCase)
	photoH := deliveryPhoto.NewPhotoDelivery(logger, photoUseCase)
	albumH := deliveryAlbum.NewAlbumDelivery(logger, albumUseCase)
	friendH := deliveryFriends.NewFriendDelivery(logger, friendsUse, chatUse)

	api := &(RequestHandlers{

		photoHandler:  photoH,
		albumHandler:  albumH,
		userHandler:   userH,
		feedHandler:   feedH,
		likeHandler:   likeH,
		friendHandler: friendH,
	})

	return api
}

func LoadMicroservices(server *echo.Echo) (authorMicro.SessionCheckerClient, likeMicro.LikeCheckerClient, photoMicro.PhotoCheckerClient, []*grpc.ClientConn) {

	connections := make([]*grpc.ClientConn, 0)

	authPORT := os.Getenv("AUTHORIZ_PORT")

	authConn, err := grpc.Dial(
		"127.0.0.1"+authPORT,
		grpc.WithInsecure(),
	)
	connections = append(connections, authConn)

	if err != nil {
		server.Logger.Fatal("cant connect to grpc")
	}

	authManager := authorMicro.NewSessionCheckerClient(authConn)

	likePORT := os.Getenv("LIKE_PORT")
	likeConn, err := grpc.Dial(
		"127.0.0.1"+likePORT,
		grpc.WithInsecure(),
	)
	connections = append(connections, likeConn)

	if err != nil {
		server.Logger.Fatal("cant connect to grpc")
	}

	likeManager := likeMicro.NewLikeCheckerClient(likeConn)

	photoPORT := os.Getenv("PHOTO_PORT")
	photoConn, err := grpc.Dial(
		"127.0.0.1"+photoPORT,
		grpc.WithInsecure(),
	)
	connections = append(connections, photoConn)

	if err != nil {
		server.Logger.Fatal("cant connect to grpc")
	}

	photoManager := photoMicro.NewPhotoCheckerClient(photoConn)

	return authManager, likeManager, photoManager, connections

}

func RegisterMetrics(server *echo.Echo) *prometheus.CounterVec {
	tracker := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "url_tracker",
	}, []string{"status", "path", "method", "time"})

	prometheus.MustRegister(tracker)

	server.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	return tracker
}

func main() {

	server := echo.New()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	db := InitializeDataBases(server)
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	origin := os.Getenv("ORIGIN_POLICY")

	auth, like, photo, authConn := LoadMicroservices(server)
	defer func() {
		if len(authConn) > 0 {
			for i, _ := range authConn {
				authConn[i].Close()
			}
		}
	}()

	tracker := RegisterMetrics(server)

	midHandler := middleware.NewMiddlewareHandler(logger, auth, tracker, origin)
	midHandler.SetMiddleware(server)

	api := NewRequestHandler(db, auth, like, photo, logger)

	api.userHandler.InitHandlers(server)
	api.feedHandler.InitHandlers(server)
	api.likeHandler.InitHandlers(server)
	api.photoHandler.InitHandlers(server)
	api.albumHandler.InitHandlers(server)
	api.friendHandler.InitHandlers(server)

	port := os.Getenv("PORT")

	server.Logger.Fatal(server.Start(port))
}
