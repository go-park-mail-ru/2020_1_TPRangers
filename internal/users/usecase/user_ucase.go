package usecase

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha1"
	"golang.org/x/crypto/pbkdf2"
	FeedRep "main/internal/feeds"
	FriendRep "main/internal/friends"
	sessions "main/internal/microservices/authorization/delivery"
	"main/internal/tools/errors"
	"main/internal/users"
	"main/models"
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

func (userR UserUseCaseRealisation) SearchUsers(userID int, searchOfValue string) ([]models.Person, error) {
	sendData, err := userR.userDB.SearchUsers(userID, searchOfValue)

	return sendData, err
}

type UserUseCaseRealisation struct {
	userDB   users.UserRepository
	friendDB FriendRep.FriendRepository
	feedDB   FeedRep.FeedRepository
	sess     sessions.SessionCheckerClient
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

		photoId, _ := userR.userDB.UploadProfilePhoto(jsonData.Photo)

		currentUserData.Photo = photoId
	}

	if jsonData.Telephone != "" {
		currentUserData.Telephone = jsonData.Telephone
	}

	if jsonData.Email != "" {
		currentUserData.Email = jsonData.Email
	}

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

func (userR UserUseCaseRealisation) Login(userData models.Auth) (string, error) {

	cookie, err := userR.sess.LoginUser(context.Background(), &sessions.Auth{
		Login:    userData.Login,
		Password: userData.Password,
	})

	if cookie == nil {
		return "", err
	}

	return cookie.Cookies, nil

}

func (userR UserUseCaseRealisation) Register(userData models.Register) (string, error) {

	cookie, err := userR.sess.CreateNewUser(context.Background(), &sessions.Register{
		Email:    userData.Email,
		Password: userData.Password,
		Name:     userData.Name,
		Surname:  userData.Surname,
		Phone:    userData.Phone,
		Date:     userData.Date,
	})

	if cookie == nil {
		return "", err
	}

	return cookie.Cookies, nil
}

func (userR UserUseCaseRealisation) Logout(cookie string) error {

	_, err := userR.sess.DeleteSession(context.Background(), &sessions.SessionData{
		Cookies: cookie,
	})

	return err
}

func (userR UserUseCaseRealisation) GetUserLoginByCookie(userId int) (string, error) {
	return userR.userDB.GetUserLoginById(userId)
}

func NewUserUseCaseRealisation(userDB users.UserRepository, friendDb FriendRep.FriendRepository, feedDB FeedRep.FeedRepository, sessChecker sessions.SessionCheckerClient) UserUseCaseRealisation {
	return UserUseCaseRealisation{
		userDB:   userDB,
		feedDB:   feedDB,
		sess:     sessChecker,
		friendDB: friendDb,
	}
}
