package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	repositoryLikes "main/internal/like/repository"
	likes "main/internal/microservices/likes/delivery"
	lks "main/internal/microservices/likes/usecase"
	"net"
	"os"
)

func InitializeDataBases() repositoryLikes.LikeRepositoryRealisation {
	err := godotenv.Load("like_micro.env")
	if err != nil {
		log.Fatal("ERROR AT LOADING .ENV", err.Error())
	}
	usernameDB := os.Getenv("LIKE_POSTGRES_USERNAME")
	passwordDB := os.Getenv("LIKE_POSTGRES_PASSWORD")
	nameDB := os.Getenv("LIKE_POSTGRES_NAME")

	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal("ERROR AT LOADING CONNECTING TO DB : ", err.Error())
	}

	likeDB := repositoryLikes.NewLikeRepositoryRealisation(db)

	return likeDB
}

func main() {

	likeDB := InitializeDataBases()
	port := os.Getenv("LIKE_PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("CANNOT LISTEN PORT : ", port, err.Error())
	}

	server := grpc.NewServer()

	likes.RegisterLikeCheckerServer(server, lks.NewLikeUseCaseChecker(likeDB))

	fmt.Println("starting server at " + port)
	err = server.Serve(lis)
	if err != nil {
		log.Fatal("CANNOT LISTEN PORT : ", port, err.Error())
	}
}
