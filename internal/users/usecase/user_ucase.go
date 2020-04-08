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
	FeedRep "main/internal/feeds"
	"main/internal/feeds/repository"
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

func CheckPassword(plainPass string , hashPass []byte) bool {

	fmt.Println(plainPass , hashPass)
	salt := hashPass[0:8]
	checkPass := CryptPassword(plainPass, salt)
	return bytes.Equal(hashPass,checkPass)
}

type UserUseCaseRealisation struct {
	userDB    users.UserRepository
	feedDB    FeedRep.FeedRepository
	sessionDB Sess.CookieRepository
}

func (userR UserUseCaseRealisation) GetAlbums(cookie string) ([]models.Album, error) {
	id, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return nil, errors.InvalidCookie
	}

	albums, err := userR.userDB.GetAlbums(id)

	fmt.Println(albums)

	if len(albums) == 0 {
		albums, err = userR.userDB.GetAlbums(0) // FIXME user id 0 have default album - maybe it's not cool
	}
	return albums, nil

}

func (userR UserUseCaseRealisation) GetPhotosFromAlbum(cookie string, albumID int) ([]models.Photos, error) {
	_, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return nil, errors.InvalidCookie
	}

	photos, err := userR.userDB.GetPhotosFromAlbum(albumID)

	fmt.Println(photos)

	return photos, nil
}

func (userR UserUseCaseRealisation) CreateAlbum(cookie string, albumData models.AlbumReq) error {
	uID, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return errors.InvalidCookie
	}

	err = userR.userDB.CreateAlbum(uID, albumData)

	if err != nil {
		return errors.FailReadFromDB
	}

	return nil
}

func (userR UserUseCaseRealisation) UploadPhotoToAlbum(cookie string, photoData models.PhotoInAlbum) error {
	_, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return errors.InvalidCookie
	}

	err = userR.userDB.UploadPhotoToAlbum(photoData)

	if err != nil {
		return errors.FailReadFromDB
	}

	return nil
}

func (userR UserUseCaseRealisation) GetUser(userLogin string) (map[string]interface{}, error) {

	userData, err := userR.userDB.GetUserProfileSettingsByLogin(userLogin)

	if err != nil {
		return nil, errors.NotExist
	}

	sendData := make(map[string]interface{})

	sendData["feed"], _ = userR.feedDB.GetUserPostsByLogin(userLogin)
	sendData["user"] = userData
	sendData["friends"], err = userR.userDB.GetUserFriendsByLogin(userLogin, 6)

	return sendData, err

}

func (userR UserUseCaseRealisation) Profile(cookie string) (map[string]interface{}, error) {

	id, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return nil, errors.InvalidCookie
	}

	sendData := make(map[string]interface{})
	sendData["user"], _ = userR.userDB.GetUserProfileSettingsById(id)
	sendData["friends"], err = userR.userDB.GetUserFriendsById(id, 6)
	sendData["feed"], err = userR.feedDB.GetUserPostsById(id)

	return sendData, err
}

func (userR UserUseCaseRealisation) GetSettings(cookie string) (map[string]interface{}, error) {

	id, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return nil, errors.InvalidCookie
	}

	sendData := make(map[string]interface{})

	sendData["user"], err = userR.userDB.GetUserProfileSettingsById(id)

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	return sendData, nil

}

func (userR UserUseCaseRealisation) UploadSettings(cookie string, newUserSettings models.Settings) (map[string]interface{}, error) {

	id, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return nil, errors.InvalidCookie
	}

	currentUserData, _ := userR.userDB.GetUserDataById(id)

	fmt.Println(currentUserData)

	jsonData := newUserSettings

	//когда нам будут высылать закэшированные настройки
	//if userData.Password == "" {
	//	userData.Password = currentUserData.Password
	//}
	//
	//currentUserData = userData

	if jsonData.Login != "" {
		currentUserData.Login = jsonData.Login
	}

	if jsonData.Password != "" {

		salt := make([]byte,8)
		rand.Read(salt)
		currentUserData.CryptedPassword =  CryptPassword(jsonData.Password,salt)
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

	err = userR.userDB.UploadSettings(id, currentUserData)

	if err != nil {
		return nil , err
	}

	sendData := make(map[string]interface{})

	sendData["user"], err = userR.userDB.GetUserProfileSettingsById(id)

	return sendData, err
}

func (userR UserUseCaseRealisation) CheckFriendship(cookie, friendLogin string, answer map[string]interface{}) (map[string]interface{}, error) {

	mainUserId, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return answer, errors.FailReadFromDB
	}

	friendId, err := userR.userDB.GetFriendIdByLogin(friendLogin)

	if err != nil {
		return answer, errors.FailReadFromDB
	}

	answer["isFriends"], err = userR.userDB.CheckFriendship(mainUserId, friendId)

	if err != nil {
		return answer, errors.FailReadFromDB
	}

	return answer, nil
}

func (userR UserUseCaseRealisation) GetAllFriends(login string) (map[string]interface{}, error) {

	sendData := make(map[string]interface{})
	var err error
	sendData["friends"], err = userR.userDB.GetAllFriendsByLogin(login)

	return sendData, err

}

func (userR UserUseCaseRealisation) Login(userData models.Auth, cookieValue string, exprTime time.Duration) error {

	login := userData.Login
	password := userData.Password
	dbPassword, existErr := userR.userDB.GetPassword(login)

	if existErr != nil {
		return errors.WrongLogin
	}

	if ! CheckPassword(password, dbPassword) {
		return errors.WrongPassword
	}

	id, existErr := userR.userDB.GetIdByEmail(login)

	if existErr != nil {
		return errors.WrongLogin
	}

	userR.sessionDB.AddCookie(id, cookieValue, exprTime)

	return nil

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

	userR.sessionDB.AddCookie(id, cookieValue, exprTime)

	return nil

}

func (userR UserUseCaseRealisation) Logout(cookie string) error {

	err := userR.sessionDB.DeleteCookie(cookie)

	return err
}

func (userR UserUseCaseRealisation) AddFriend(cookie, friendLogin string) error {

	id, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return errors.InvalidCookie
	}

	friendId, _ := userR.userDB.GetFriendIdByLogin(friendLogin)

	err = userR.userDB.AddFriend(id, friendId)

	if err != nil {
		return errors.FailAddFriend
	}

	return err
}

func (userR UserUseCaseRealisation) GetUserLoginByCookie(cookieValue string) (string, error) {
	id, _ := userR.sessionDB.GetUserIdByCookie(cookieValue)

	return userR.userDB.GetUserLoginById(id)
}

func NewUserUseCaseRealisation(userDB UserRep.UserRepositoryRealisation, feedDB repository.FeedRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation) UserUseCaseRealisation {
	return UserUseCaseRealisation{
		userDB:    userDB,
		feedDB:    feedDB,
		sessionDB: sesDB,
	}
}
