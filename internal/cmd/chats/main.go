package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	metrics "main/internal/metrics/delivery"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"

	repositoryMessage "main/internal/message/repository"
	deliveryMessage "main/internal/socket/delivery"
	usecaseMessage "main/internal/socket/usecase"

	deliveryToken "main/internal/socket_token/delivery"
	repositoryToken "main/internal/socket_token/repository"
	usecaseToken "main/internal/socket_token/usecase"

	deliveryChat "main/internal/chats/delivery"
	repositoryChat "main/internal/chats/repository"
	usecaseChat "main/internal/chats/usecase"

	"main/internal/middleware"
	"os"

	repositoryFriends "main/internal/friends/repository"
	authorMicro "main/internal/microservices/authorization/delivery"
)

type ChatsHandler struct {
	messageHandler     deliveryMessage.SocketDelivery
	chatHandler        deliveryChat.ChatsDelivery
	socketTokenHandler deliveryToken.TokenDelivery
}

func InitializeDataBases(server *echo.Echo) (*sql.DB, repositoryMessage.MessageRepositoryRealisation,
	repositoryToken.TokenRepositoryRealisation) {
	err := godotenv.Load("chat_micro.env")
	if err != nil {
		server.Logger.Fatal("can't load .env file :", err.Error())
	}
	usernameDB := os.Getenv("CHAT_POSTGRES_USERNAME")
	passwordDB := os.Getenv("CHAT_POSTGRES_PASSWORD")
	nameDB := os.Getenv("CHAT_POSTGRES_NAME")

	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		server.Logger.Fatal("NO CONNECTION TO BD", err.Error())
	}

	redisPas := os.Getenv("REDIS_CHAT_PASSWORD")
	redisPort := os.Getenv("REDIS_CHAT_PORT")

	redisChatPort := os.Getenv("REDIS_CHAT_PORT")

	tokenDB := repositoryToken.NewTokenRepositoryRealisation(redisPort, redisPas)
	chatDB := repositoryMessage.NewMessageRepositoryRealisation(redisChatPort, "", db)

	return db, chatDB, tokenDB
}

func NewRequestHandler(db *sql.DB, messageDB repositoryMessage.MessageRepositoryRealisation,
	tokenDB repositoryToken.TokenRepositoryRealisation, logger *zap.SugaredLogger) *ChatsHandler {

	chatDB := repositoryChat.NewChatRepositoryRealisation(db)
	friendsDB := repositoryFriends.NewFriendRepositoryRealisation(db)

	messageUseCase := usecaseMessage.NewSocketUseCaseRealisation(messageDB, tokenDB)
	chatUse := usecaseChat.NewChatUseCaseRealisation(chatDB, friendsDB)
	socketTokenUse := usecaseToken.NewTokenUseCaseRealisation(tokenDB)

	messageH := deliveryMessage.NewSocketDelivery(logger, messageUseCase)
	chatH := deliveryChat.NewChatsDelivery(logger, chatUse)
	socketTokenH := deliveryToken.NewTokenDelivery(logger, socketTokenUse)

	api := &(ChatsHandler{
		messageHandler:     messageH,
		chatHandler:        chatH,
		socketTokenHandler: socketTokenH,
	})

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

	server := echo.New()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	db, messages, tokens := InitializeDataBases(server)
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	origin := os.Getenv("ORIGIN_POLICY")

	auth, authConn := LoadMicroservices(server)

	defer func() {
		if authConn != nil {
			authConn.Close()
		}
	}()

	tracker := metrics.RegisterMetrics(server)


	midHandler := middleware.NewMiddlewareHandler(logger, auth, tracker,origin)
	midHandler.SetMiddleware(server)

	api := NewRequestHandler(db, messages, tokens, logger)

	api.messageHandler.InitHandlers(server)
	api.chatHandler.InitHandlers(server)
	api.socketTokenHandler.InitHandlers(server)

	port := os.Getenv("CHAT_PORT")

	server.Logger.Fatal(server.Start(port))
}
