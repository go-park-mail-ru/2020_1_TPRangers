package main

import(
	"../models"
	"../errors"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)


type DataBaseAPI interface {
	AddNewUser(userData models.User) error
	IsUserExist(login string) bool
	GetIdByLogin(login string) int
	GetPassword(login string) string
	GetUserFeed(login string, count int) models.Feed
	GetUserDataByLogin(login string) models.User
	UploadSettings(login string, currentUserData models.User)
}

type DataBase struct{

}



//var connStr = "user=alexandr password=nikita2003 dbname=VK sslmode=disable"

func (Data DataBase) AddNewUser(userData models.User) error {

	db, err := sql.Open("postgres", "user=alexandr password=nikita2003 dbname=VK sslmode=disable")
	if err != nil {
		return errors.FailConnect
	}
	defer db.Close()
	fmt.Println("I am here")
	result, err := db.Exec("insert into Users (phone, mail, name, surname, password, birthdate) values ($1, $2, $3, $4, $5, $6)", userData.Telephone, userData.Email, userData.Name, userData.Surname, userData.Password, userData.Date)
	if err != nil{
		return errors.FailSendToDB
	}
	fmt.Println(result)
	return nil
}

func main() {
	var kek DataBase
	data := models.User{
		Telephone: "222",
		Email:     "sanya554455@gmail.com",
		Name:      "Alexandr",
		Password:  "nikita2003",
		Surname:   "Dolgavin",
		Date:      "16.05.2000",
	}

	err := kek.AddNewUser(data)

	fmt.Println(err)
}

