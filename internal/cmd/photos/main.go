package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	photos "main/internal/microservices/photos/delivery"
	phs "main/internal/microservices/photos/usecase"
	repositoryPhotos "main/internal/photos/repository"
	"net"
	"os"
)

func InitializeDataBases() repositoryPhotos.PhotoRepositoryRealisation {
	err := godotenv.Load("photo_micro.env")
	if err != nil {
		log.Fatal("ERROR AT LOADING .ENV", err.Error())
	}
	usernameDB := os.Getenv("PHOTO_POSTGRES_USERNAME")
	passwordDB := os.Getenv("PHOTO_POSTGRES_PASSWORD")
	nameDB := os.Getenv("PHOTO_POSTGRES_NAME")

	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal("ERROR AT LOADING CONNECTING TO DB : ", err.Error())
	}

	photoDB := repositoryPhotos.NewPhotoRepositoryRealisation(db)

	return photoDB
}

func main() {

	photoDB := InitializeDataBases()
	port := os.Getenv("PHOTO_PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("CANNOT LISTEN PORT : ", port, err.Error())
	}

	server := grpc.NewServer()

	photos.RegisterPhotoCheckerServer(server, phs.NewPhotoUseCaseChecker(photoDB))

	fmt.Println("starting server at " + port)
	server.Serve(lis)
}
