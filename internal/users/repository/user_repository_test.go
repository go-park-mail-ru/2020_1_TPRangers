package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

const (
	usernameDB = "postgres"
	passwordDB = "071299"
	nameDB     = "vk"
)
//UserFriendsTest
func TestUserRepositoryRealisation_GetUserFriendsById(t *testing.T) {

	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	defer db.Close()
	if err != nil {
		fmt.Println("NO DB")
	}

	testHandler := NewUserRepositoryRealisation(db)


	testValue , err :=testHandler.GetUserFriendsById(2,6)

	fmt.Println(testValue , err)

}

func TestUserRepositoryRealisation_GetUserFriendsByLogin(t *testing.T) {

	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	defer db.Close()
	if err != nil {
		fmt.Println("NO DB")
	}

	testHandler := NewUserRepositoryRealisation(db)


	testValue , err :=testHandler.GetUserFriendsByLogin("TEST@yandex.ru",6)

	fmt.Println(testValue , err)

}

func TestUserRepositoryRealisation_GetAllUserFriendsByLogin(t *testing.T) {

	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	defer db.Close()
	if err != nil {
		fmt.Println("NO DB")
	}

	testHandler := NewUserRepositoryRealisation(db)


	testValue , err :=testHandler.GetAllFriendsByLogin("123124TEST@yandex.ru")

	fmt.Println(testValue , err)

}


func TestUserRepositoryRealisation_CheckFriendship(t *testing.T) {
	connectString := "user=" + usernameDB + " password=" + passwordDB + " dbname=" + nameDB + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)
	defer db.Close()
	if err != nil {
		fmt.Println("NO DB")
	}

	testHandler := NewUserRepositoryRealisation(db)


	testValue , err :=testHandler.CheckFriendship(1,2)
	fmt.Println(testValue , err)

	testValue , err =testHandler.CheckFriendship(1,10)

	fmt.Println(testValue , err)
}


