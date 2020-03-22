package user

import (
	"../../errors"
	"../../models"
	"../../repository"
	SessRep "../../repository/cookie"
	FeedRep "../../repository/feed"
	UserRep "../../repository/user"
	uuid "github.com/satori/go.uuid"
	"time"
)

type UserUseCaseRealisation struct {
	userDB    repository.UserRepository
	feedDB    repository.FeedRepository
	sessionDB repository.CookieRepository
}

func (userR UserUseCaseRealisation) GetUser(userLogin string) (map[string]interface{}, error) {

	userData, existError := userR.userDB.GetUserProfileSettingsByLogin(userLogin)

	if existError != nil {
		return nil, errors.NotExist
	}

	sendData := make(map[string]interface{})

	sendData["feed"], _ = userR.feedDB.GetUserFeedByEmail(userLogin, 30)
	sendData["user"] = userData

	return sendData, nil

}

func (userR UserUseCaseRealisation) Profile(cookie string) (map[string]interface{}, error) {

	id, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return nil, errors.InvalidCookie
	}

	sendData := make(map[string]interface{})
	sendData["user"], _ = userR.userDB.GetUserProfileSettingsById(id)

	return sendData, nil
}

func (userR UserUseCaseRealisation) GetSettings(cookie string) (map[string]interface{}, error) {


	id, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return nil , errors.InvalidCookie
	}

	sendData := make(map[string]interface{})

	sendData["user"], err = userR.userDB.GetUserProfileSettingsById(id)

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	return sendData, nil

}

func (userR UserUseCaseRealisation) UploadSettings(cookie string , newUserSettings models.Settings) (map[string]interface{} , error) {

	id, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return  nil ,errors.InvalidCookie
	}

	currentUserData, _ := userR.userDB.GetUserDataById(id)

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
		currentUserData.Password = jsonData.Password
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
		photoId, _ := userR.userDB.UploadPhoto(jsonData.Photo)

		currentUserData.Photo = photoId
	}

	if jsonData.Telephone != "" {
		currentUserData.Telephone = jsonData.Telephone
	}

	if jsonData.Email != "" {
		currentUserData.Email = jsonData.Email
	}

	userR.userDB.UploadSettings(id, currentUserData)

	sendData := make(map[string]interface{})

	sendData["user"], _ = userR.userDB.GetUserProfileSettingsById(id)

	return sendData , nil
}

func (userR UserUseCaseRealisation) Login(userData models.Auth , cookieValue string,exprTime time.Duration) error {


	login := userData.Login
	password := userData.Password
	dbPassword, existErr := userR.userDB.GetPassword(login)

	if existErr != nil {
		return errors.WrongLogin
	}

	if password != dbPassword {
		return errors.WrongPassword
	}

	id, existErr := userR.userDB.GetIdByEmail(login)

	if existErr != nil {
		return errors.WrongLogin
	}

	userR.sessionDB.AddCookie(id, cookieValue ,exprTime)

	return nil

}

func (userR UserUseCaseRealisation) Register(userData models.Register, cookieValue string , exprTime time.Duration) error {


	email := userData.Email


	if flag, _ := userR.userDB.IsUserExist(email); flag == true {
		return errors.AlreadyExist
	}

	uniqueUserLogin , _ := uuid.NewV4()

	defaultPhotoId , _:= userR.userDB.GetDefaultProfilePhotoId()

	data := models.User{
		Login:     uniqueUserLogin.String(),
		Telephone: userData.Phone,
		Email:     email,
		Name:      userData.Name,
		Password:  userData.Password,
		Surname:   userData.Surname,
		Date:      userData.Date,
		Photo:     defaultPhotoId,
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

func (userR UserUseCaseRealisation) AddFriend(cookie , friendLogin string) error {


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

func NewUserUseCaseRealisation(userDB UserRep.UserRepositoryRealisation, feedDB FeedRep.FeedRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation) UserUseCaseRealisation {
	return UserUseCaseRealisation{
		userDB:    userDB,
		feedDB:    feedDB,
		sessionDB: sesDB,
	}
}
