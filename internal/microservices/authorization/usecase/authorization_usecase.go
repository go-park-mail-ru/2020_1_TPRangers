package usecase

import (
	"bytes"
	context "context"
	"crypto/rand"
	"crypto/sha1"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/pbkdf2"
	"main/internal/cookies"
	"main/internal/microservices/authorization/delivery"
	"main/internal/models"
	"main/internal/tools/errors"
	"main/internal/users"
	"time"
)

type AuthorizationUseCaseRealisation struct {
	userDB    users.UserRepository
	sessionDB cookies.CookieRepository
}

func NewAuthorizationUseCaseRealisation(userDB users.UserRepository, sessionDB cookies.CookieRepository) AuthorizationUseCaseRealisation {
	return AuthorizationUseCaseRealisation{
		userDB:    userDB,
		sessionDB: sessionDB,
	}
}

func CryptPassword(pass string, salt []byte) []byte {
	cryptedPass := pbkdf2.Key([]byte(pass), salt, 4096, 32, sha1.New)
	return append(salt, cryptedPass...)
}

func CheckPassword(plainPass string, hashPass []byte) bool {
	salt := hashPass[0:8]
	checkPass := CryptPassword(plainPass, salt)
	return bytes.Equal(hashPass, checkPass)
}

func (AU AuthorizationUseCaseRealisation) LoginUser(ctx context.Context, auth *session.Auth) (*session.SessionData, error) {
	login := auth.Login

	password := auth.Password
	dbPassword, existErr := AU.userDB.GetPassword(login)

	cookieValue := uuid.NewV4()
	session := &session.SessionData{
		Cookies: cookieValue.String(),
	}

	if existErr != nil {
		return nil, errors.WrongLogin
	}

	if !CheckPassword(password, dbPassword) {
		return nil, errors.WrongPassword
	}

	id, existErr := AU.userDB.GetIdByEmail(login)

	if existErr != nil {
		return nil, errors.WrongLogin
	}

	err := AU.sessionDB.AddCookie(id, session.Cookies, 15*time.Hour)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (AU AuthorizationUseCaseRealisation) CreateNewUser(ctx context.Context, userData *session.Register) (*session.SessionData, error) {

	email := userData.Email

	if flag, _ := AU.userDB.IsUserExist(email); flag {
		return nil, errors.AlreadyExist
	}

	uniqueUserLogin := uuid.NewV4()

	defaultPhotoId, _ := AU.userDB.GetDefaultProfilePhotoId()

	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
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

	err = AU.userDB.AddNewUser(data)

	if err != nil {
		return nil, err
	}

	id, err := AU.userDB.GetIdByEmail(email)

	if err != nil {
		return nil, err
	}

	cookieValue := uuid.NewV4()
	session := &session.SessionData{
		Cookies: cookieValue.String(),
	}

	err = AU.sessionDB.AddCookie(id, session.Cookies, 15*time.Hour)

	return session, err

}

func (AU AuthorizationUseCaseRealisation) CheckSession(ctx context.Context, auth *session.SessionData) (*session.UserId, error) {

	userId, err := AU.sessionDB.GetUserIdByCookie(auth.Cookies)

	if userId == 0 || err != nil {
		userId = -1
	}

	userSession := &session.UserId{
		UserId: int32(userId),
	}

	return userSession, err
}

func (AU AuthorizationUseCaseRealisation) DeleteSession(ctx context.Context, auth *session.SessionData) (*session.UserId, error) {

	err := AU.sessionDB.DeleteCookie(auth.Cookies)

	return &session.UserId{
		UserId: -1,
	}, err

}
