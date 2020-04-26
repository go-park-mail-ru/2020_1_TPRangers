package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	repositoryCookie "main/internal/cookies/repository"
	"main/internal/microservices/authorization/delivery"
	"main/internal/microservices/authorization/usecase"
	repositoryUser "main/internal/users/repository"
	"net"
	"os"
)

func InitializeDataBases() (repositoryUser.UserRepositoryRealisation, repositoryCookie.CookieRepositoryRealisation) {
	err := godotenv.Load("author_micro.env")
	if err != nil {
		log.Fatal("ERROR AT LOADING .ENV", err.Error())
	}
	usernameDB := os.Getenv("AUTHORIZ_POSTGRES_USERNAME")
	passwordDB := os.Getenv("AUTHORIZ_POSTGRES_PASSWORD")
	nameDB := os.Getenv("AUTHORIZ_POSTGRES_NAME")

	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal("ERROR AT LOADING CONNECTING TO DB : ", err.Error())
	}

	redisPas := os.Getenv("AUTHORIZ_REDIS_PASSWORD")
	redisPort := os.Getenv("AUTHORIZ_REDIS_PORT")

	sessionDB := repositoryCookie.NewCookieRepositoryRealisation(redisPort, redisPas)
	userDB := repositoryUser.NewUserRepositoryRealisation(db)

	return userDB, sessionDB
}

func main() {

	users, sessions := InitializeDataBases()
	port := os.Getenv("AUTHORIZ_PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("CANNOT LISTEN PORT : ", port, err.Error())
	}

	server := grpc.NewServer()

	session.RegisterSessionCheckerServer(server, usecase.NewAuthorizationUseCaseRealisation(users, sessions))

	fmt.Println("starting server at :3080")
	server.Serve(lis)
}
