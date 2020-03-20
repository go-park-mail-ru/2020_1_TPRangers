package settings

import (
	"../../errors"
	"../../models"
	"../../repository"
	UserRep "../../repository/user"
	SessRep "../../repository/cookie"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type SettingsUseCaseRealisation struct {
	settingsDB repository.UserRepository
	sessionDB  repository.CookieRepository
	logger     *zap.SugaredLogger
}

func (stngR SettingsUseCaseRealisation) GetSettings(rwContext echo.Context, uId string) (error, models.JsonStruct) {

	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {
		stngR.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", "no-cookie"),
		)
		return errors.CookieExpired, models.JsonStruct{Err: errors.CookieExpired.Error()}
	}

	id, err := stngR.sessionDB.GetUserIdByCookie(cookie.Value)

	if err != nil {

		stngR.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", cookie.Value),
			zap.String("ID", "no such user"),
		)

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		return errors.InvalidCookie, models.JsonStruct{Err: errors.InvalidCookie.Error()}
	}

	stngR.logger.Debug(
		zap.String("ID", uId),
		zap.String("COOKIE", cookie.Value),
		zap.Int("ID", id),
	)

	sendData := make(map[string]interface{})

	sendData["user"], _ = stngR.settingsDB.GetUserDataById(id)

	return nil, models.JsonStruct{Body: sendData}

}

func (stngR SettingsUseCaseRealisation) UploadSettings(rwContext echo.Context, uId string) (error, models.JsonStruct) {

	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {

		stngR.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", "no-cookie"),
		)

		return errors.CookieExpired, models.JsonStruct{Err: errors.CookieExpired.Error()}
	}

	id, err := stngR.sessionDB.GetUserIdByCookie(cookie.Value)

	if err != nil {

		stngR.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", cookie.Value),
			zap.String("USER-ID", "no such user"),
		)

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		return errors.InvalidCookie, models.JsonStruct{Err: errors.InvalidCookie.Error()}
	}

	stngR.logger.Debug(
		zap.String("ID", uId),
		zap.String("COOKIE", cookie.Value),
		zap.Int("USER-ID", id),
	)

	uploadDataFlags := []string{"uploadedFile", "email", "login", "password", "name", "surname", "phone", "date"}

	currentUserData, _ := stngR.settingsDB.GetUserDataById(id)

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
				currentUserData.Name = data
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

	stngR.settingsDB.UploadSettings(id, currentUserData)

	sendData := make(map[string]interface{})

	sendData["user"], _ = stngR.settingsDB.GetUserDataById(id)

	return nil, models.JsonStruct{Body: sendData}
}

func NewSetUseCaseRealisation(userDB UserRep.UserRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation, log *zap.SugaredLogger) SettingsUseCaseRealisation {
	return SettingsUseCaseRealisation{
		settingsDB: userDB,
		sessionDB:  sesDB,
		logger:     log,
	}
}
