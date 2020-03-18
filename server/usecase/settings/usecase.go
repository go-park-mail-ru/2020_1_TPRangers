package settings

import (
	"../../errors"
	"../../models"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type SettingsUseCaseRealisation struct {
	settingsDB REPOSITORYCASETYPE
	sessionDB  REPOSITORYSESSIONTYPE
}

func (stngR SettingsUseCaseRealisation) GetSettings(rwContext echo.Context) (error, models.JsonStruct) {

	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {
		return errors.CookieExpired, models.JsonStruct{Err: errors.CookieExpired.Error()}
	}

	login, err := stngR.sessionDB.GetUserByCookie(cookie.Value)

	if err != nil {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		return errors.InvalidCookie, models.JsonStruct{Err: errors.InvalidCookie.Error()}
	}

	sendData := make(map[string]interface{})

	sendData["user"], _ = stngR.settingsDB.GetUserDataByLogin(login)

	return nil, models.JsonStruct{Body: sendData}

}

func (stngR SettingsUseCaseRealisation) UploadSettings(rwContext echo.Context) (error, models.JsonStruct) {

	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {
		return errors.CookieExpired, models.JsonStruct{Err: errors.CookieExpired.Error()}
	}

	login, err := stngR.sessionDB.GetUserByCookie(cookie.Value)

	if err != nil {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		return errors.InvalidCookie, models.JsonStruct{Err: errors.InvalidCookie.Error()}
	}

	uploadDataFlags := []string{"uploadedFile", "email", "login", "password", "name", "surname", "phone", "date"}

	currentUserData, _ := stngR.settingsDB.GetUserDataByLogin(login)

	for _, dataFlag := range uploadDataFlags {
		switch dataFlag {
		case "uploadedFile":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Photo = data
			}
		case "email":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Email = data
			}
		case "password":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Password = data
			}
		case "name":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Username = data
			}
		case "phone":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Telephone = data
			}
		case "date":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Date = data
			}
		case "login":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Login = data
			}
		case "surname":
			if data := rwContext.FormValue(dataFlag); data != "" {
				currentUserData.Surname = data
			}

		}
	}

	stngR.settingsDB.UploadSettings(login, currentUserData)

	sendData := make(map[string]interface{})

	sendData["user"], _ = stngR.settingsDB.GetUserDataByLogin(login)

	return nil, models.JsonStruct{Body: sendData}
}
