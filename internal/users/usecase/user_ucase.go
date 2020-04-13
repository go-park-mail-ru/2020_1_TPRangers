package usecase

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/pbkdf2"
	Sess "main/internal/cookies"
	SessRep "main/internal/cookies/repository"
	"main/internal/csrf"
	FeedRep "main/internal/feeds"
	"main/internal/feeds/repository"
	FriendRep "main/internal/friends"
	"main/internal/models"
	"main/internal/tools/errors"
	"main/internal/users"
	UserRep "main/internal/users/repository"
	"time"
)

func CryptPassword(pass string, salt []byte) []byte {
	cryptedPass := pbkdf2.Key([]byte(pass), salt, 4096, 32, sha1.New)
	return append(salt, cryptedPass...)
}

func CheckPassword(plainPass string, hashPass []byte) bool {
	salt := hashPass[0:8]
	checkPass := CryptPassword(plainPass, salt)
	return bytes.Equal(hashPass, checkPass)
}

type UserUseCaseRealisation struct {
	userDB    users.UserRepository
	friendDB  FriendRep.FriendRepository
	feedDB    FeedRep.FeedRepository
	sessionDB Sess.CookieRepository
}

func (userR UserUseCaseRealisation) GetOtherUserProfileNotLogged(userLogin string) (models.OtherUserProfileData, error) {

	sendData := new(models.OtherUserProfileData)
	var err error

	sendData.User, err = userR.userDB.GetUserProfileSettingsByLogin(userLogin)

	if err != nil {
		return *sendData, errors.NotExist
	}

	sendData.Feed, _ = userR.feedDB.GetUserPostsByLogin(userLogin)
	sendData.Friends, err = userR.friendDB.GetUserFriendsByLogin(userLogin, 6)

	return *sendData, err

}

func (userR UserUseCaseRealisation) GetUserProfileWhileLogged(otherUserLogin string, currentUserId int) (models.OtherUserProfileData, error) {

	sendData := new(models.OtherUserProfileData)
	var err error

	sendData.User, err = userR.userDB.GetUserProfileSettingsByLogin(otherUserLogin)

	if err != nil {
		return *sendData, errors.NotExist
	}

	sendData.Feed, _ = userR.feedDB.GetPostsOfOtherUserWhileLogged(otherUserLogin, currentUserId)
	sendData.Friends, err = userR.friendDB.GetUserFriendsByLogin(otherUserLogin, 6)

	return *sendData, err

}

func (userR UserUseCaseRealisation) GetMainUserProfile(userId int) (models.MainUserProfileData, error) {

	sendData := new(models.MainUserProfileData)
	var err error

	sendData.User, _ = userR.userDB.GetUserProfileSettingsById(userId)
	sendData.Friends, err = userR.friendDB.GetUserFriendsById(userId, 6)
	sendData.Feed, err = userR.feedDB.GetUserPostsById(userId)

	return *sendData, err
}

func (userR UserUseCaseRealisation) GetSettings(userId int) (models.Settings, error) {

	sendData, err := userR.userDB.GetUserProfileSettingsById(userId)

	if err != nil {
		return sendData, errors.FailReadFromDB
	}

	return sendData, nil

}

func (userR UserUseCaseRealisation) UploadSettings(userId int, newUserSettings models.Settings) (models.Settings, error) {

	currentUserData, _ := userR.userDB.GetUserDataById(userId)

	fmt.Println(currentUserData)

	jsonData := newUserSettings

	if jsonData.Login != "" {
		currentUserData.Login = jsonData.Login
	}

	if jsonData.Password != "" {

		salt := make([]byte, 8)
		rand.Read(salt)
		currentUserData.CryptedPassword = CryptPassword(jsonData.Password, salt)
	}

	if jsonData.Date != "" {
		currentUserData.Date = jsonData.Date
	}

	if jsonData.Surname != "" {
		currentUserData.Surname = jsonData.Surname
	}

	if jsonData.Name != "" {
		currentUserData.Name = jsonData.Name
	}

	if jsonData.Photo != "" {

		fmt.Println(jsonData.Photo)

		photoId, _ := userR.userDB.UploadPhoto(jsonData.Photo)

		currentUserData.Photo = photoId
	}

	if jsonData.Telephone != "" {
		currentUserData.Telephone = jsonData.Telephone
	}

	if jsonData.Email != "" {
		currentUserData.Email = jsonData.Email
	}

	fmt.Println(currentUserData)

	err := userR.userDB.UploadSettings(userId, currentUserData)

	sendData := models.Settings{}

	if err != nil {
		return sendData, err
	}

	sendData, err = userR.userDB.GetUserProfileSettingsById(userId)

	return sendData, err
}

func (userR UserUseCaseRealisation) CheckFriendship(mainUserId int, friendLogin string) (bool, error) {

	friendId, err := userR.friendDB.GetFriendIdByLogin(friendLogin)

	if err != nil {
		return false, errors.FailReadFromDB
	}

	friendShipStatus, err := userR.friendDB.CheckFriendship(mainUserId, friendId)

	if err != nil {
		return false, errors.FailReadFromDB
	}

	return friendShipStatus, nil
}

func (userR UserUseCaseRealisation) Login(userData models.Auth, cookieValue string, exprTime time.Duration) (string, error) {
	login := userData.Login

	token, _ := csrf.Tokens.Create(login, cookieValue, 999999)

	password := userData.Password
	dbPassword, existErr := userR.userDB.GetPassword(login)

	if existErr != nil {
		return "", errors.WrongLogin
	}

	if !CheckPassword(password, dbPassword) {
		return "", errors.WrongPassword
	}

	id, existErr := userR.userDB.GetIdByEmail(login)

	if existErr != nil {
		return "", errors.WrongLogin
	}

	err := userR.sessionDB.AddCookie(id, cookieValue, exprTime)

	return token, err

}

func (userR UserUseCaseRealisation) Register(userData models.Register, cookieValue string, exprTime time.Duration) error {

	email := userData.Email

	if flag, _ := userR.userDB.IsUserExist(email); flag == true {
		return errors.AlreadyExist
	}

	uniqueUserLogin := uuid.NewV4()

	defaultPhotoId, _ := userR.userDB.GetDefaultProfilePhotoId()

	salt := make([]byte, 8)
	rand.Read(salt)
	crypPass := CryptPassword(userData.Password, salt)

	data := models.User{
		Login:           uniqueUserLogin.String(),
		Telephone:       userData.Phone,
		Email:           email,
		Name:            userData.Name,
		CryptedPassword: crypPass,
		Surname:         userData.Surname,
		Date:            userData.Date,
		Photo:           defaultPhotoId,
	}

	userR.userDB.AddNewUser(data)

	id, err := userR.userDB.GetIdByEmail(email)

	if err != nil {
		return errors.FailReadFromDB
	}

	err = userR.sessionDB.AddCookie(id, cookieValue, exprTime)

	return err

}

func (userR UserUseCaseRealisation) Logout(cookie string) error {

	err := userR.sessionDB.DeleteCookie(cookie)

	return err
}

func (userR UserUseCaseRealisation) GetUserLoginByCookie(userId int) (string, error) {

	return userR.userDB.GetUserLoginById(userId)
}

func NewUserUseCaseRealisation(userDB UserRep.UserRepositoryRealisation, friendDb FriendRep.FriendRepository, feedDB repository.FeedRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation) UserUseCaseRealisation {
	return UserUseCaseRealisation{
		userDB:    userDB,
		feedDB:    feedDB,
		sessionDB: sesDB,
		friendDB:  friendDb,
	}
}
