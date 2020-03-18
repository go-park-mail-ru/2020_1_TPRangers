package auth

import (
	".."
	"../../errors"
	"../../models"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type AuthUseCaseRealisation struct {
	loginDB   REPOSITORYCASETYPE
	sessionDB REPOSITORYSESSIONTYPE
}

func (logR AuthUseCaseRealisation) Login(rwContext echo.Context) error {

	jsonData, convertionError := usecase.GetDataFromJson("log", rwContext)

	if convertionError != nil {
		return convertionError
	}

	userData := jsonData.(*models.Auth)
	login := userData.Login
	password := userData.Password
	dbPassword := logR.loginDB.GetPassword(login)

	if dbPassword == "" {
		return errors.WrongLogin
	}

	if password != dbPassword {
		return errors.WrongPassword
	}

	id := logR.loginDB.GetIdByLogin(login)
	info, _ := uuid.NewV4()
	cookieValue := info.String()
	logR.sessionDB.SetCookie(id, cookie)

	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(12 * time.Hour),
	}
	rwContext.SetCookie(&cookie)

	return nil

}

func (logR AuthUseCaseRealisation) Logout(rwContext echo.Context) error {

	cookie, err := rwContext.Cookie("session_id")

	if err == nil {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)
	}

	return nil
}
