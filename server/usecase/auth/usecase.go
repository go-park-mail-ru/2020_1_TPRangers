package auth

import (
	".."
	"../../errors"
	"../../models"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"time"
)

type AuthRealisation struct {
	loginDB   REPOSITORYCASETYPE
	sessionDB REPOSITORYSESSIONTYPE
}

func (lg AuthRealisation) Login(rwContext echo.Context) (error, string) {

	jsonData, convertionError := usecase.GetDataFromJson("log", rwContext)

	if convertionError != nil {
		return convertionError, ""
	}

	userData := jsonData.(*models.Auth)
	login := userData.Login
	password := userData.Password
	dbPassword := lg.loginDB.GetPassword(login)

	if dbPassword == "" {
		return errors.WrongLogin, ""
	}

	if password != dbPassword {
		return errors.WrongPassword, ""
	}

	id := lg.loginDB.GetIdByLogin(login)
	info, _ := uuid.NewV4()
	cookie := info.String()
	lg.sessionDB.SetCookie(id, cookie)

	return nil , cookie

}

func (lg AuthRealisation) Logout(rwContext echo.Context) error {

	cookie, err := rwContext.Cookie("session_id")

	if err == nil {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)
	}

	return nil
}
