package usecase

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
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

type UserUseCaseRealisation struct {
	userDB    users.UserRepository
	feedDB    FeedRep.FeedRepository
	sessionDB Sess.CookieRepository
}

func (userR UserUseCaseRealisation) GetUser(userLogin string) (map[string]interface{}, error) {

	userData, err := userR.userDB.GetUserProfileSettingsByLogin(userLogin)

	if err != nil {
		return nil, errors.NotExist
	}

	sendData := make(map[string]interface{})

	sendData["feed"], _ = userR.feedDB.GetUserFeedByEmail(userLogin, 30)
	sendData["user"] = userData
	sendData["friends"] , err = userR.userDB.GetUserFriendsByLogin(userLogin,6)

	return sendData, err

}

func (userR UserUseCaseRealisation) Profile(cookie string) (map[string]interface{}, error) {

	id, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return nil, errors.InvalidCookie
	}

	sendData := make(map[string]interface{})
	sendData["user"], _ = userR.userDB.GetUserProfileSettingsById(id)
	sendData["friends"] , err = userR.userDB.GetUserFriendsById(id,6)

	return sendData, err
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

	userR.userDB.UploadSettings(id, currentUserData)

	sendData := make(map[string]interface{})

	sendData["user"], _ = userR.userDB.GetUserProfileSettingsById(id)

	return sendData , nil
}

func (userR UserUseCaseRealisation) CheckFriendship(cookie , friendLogin string , answer map[string]interface{}) (map[string]interface{} , error){

	mainUserId, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return answer , errors.FailReadFromDB
	}

	friendId , err := userR.userDB.GetFriendIdByLogin(friendLogin)

	if err != nil {
		return answer , errors.FailReadFromDB
	}


	answer["isFriends"] , err = userR.userDB.CheckFriendship(mainUserId , friendId)

	if err != nil {
		return answer , errors.FailReadFromDB
	}

	return answer , nil
}

func (userR UserUseCaseRealisation) GetAllFriends(login string) (map[string]interface{},error) {

	sendData := make(map[string]interface{})
	var err error
	sendData["friends"] , err = userR.userDB.GetAllFriendsByLogin(login)

	return sendData, err

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

	uniqueUserLogin := uuid.NewV4()

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

func NewUserUseCaseRealisation(userDB UserRep.UserRepositoryRealisation, feedDB repository.FeedRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation) UserUseCaseRealisation {
	return UserUseCaseRealisation{
		userDB:    userDB,
		feedDB:    feedDB,
		sessionDB: sesDB,
	}
}