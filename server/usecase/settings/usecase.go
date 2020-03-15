package settings

import (
	"../../errors"
	"../../models"
	"github.com/labstack/echo"
	"net/http"
)

type SettingsRealisation struct {
	settingsDB REPOSITORYCASETYPE
	sessionDB  REPOSITORYSESSIONTYPE
}

func (stng SettingsRealisation) GetSettings(rwContext echo.Context) (error, models.JsonStruct) {

	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {
		return errors.CookieExpired, models.JsonStruct{}
	}

	login, err := stng.sessionDB.GetUserByCookie(cookie.Value)

	if err != nil {
		return errors.InvalidCookie, models.JsonStruct{}
	}

	sendData := make(map[string]interface{})

	sendData["user"], _ = stng.settingsDB.GetUserDataByLogin(login)

	return nil, models.JsonStruct{Body: sendData}

}

func (stng SettingsRealisation) UploadSettings(rwContext echo.Context) (error, models.JsonStruct) {

	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {
		return errors.CookieExpired, models.JsonStruct{}
	}

	login, err := stng.sessionDB.GetUserByCookie(cookie.Value)

	if err != nil {
		return errors.InvalidCookie, models.JsonStruct{}
	}

	uploadDataFlags := []string{"uploadedFile", "email", "login", "password", "name", "surname", "phone", "date"}

	currentUserData, _ := stng.settingsDB.GetUserDataByLogin(login)

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

	stng.settingsDB.UploadSettings(login, currentUserData)

	sendData := make(map[string]interface{})

	sendData["user"], _ = stng.settingsDB.GetUserDataByLogin(login)

	return nil, models.JsonStruct{Body: sendData}
}
