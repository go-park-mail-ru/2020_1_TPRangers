package settings

import (
	"../../errors"
	"../../models"
	"../../repository"
	SessRep "../../repository/cookie"
	UserRep "../../repository/user"
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

	sendData["user"], err = stngR.settingsDB.GetUserProfileSettingsById(id)

	if err != nil {

		stngR.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", cookie.Value),
			zap.String("ID", "db error"),
		)

		return errors.FailReadFromDB, models.JsonStruct{Err: errors.FailReadFromDB.Error()}
	}

	stngR.logger.Debug(
		zap.String("ID", uId),
		sendData["user"],
	)

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

	currentUserData, _ := stngR.settingsDB.GetUserDataById(id)

	jsonData := new(models.Settings)

	convertionError :=  rwContext.Bind(jsonData)

	//когда нам будут высылать закэшированные настройки
	//if userData.Password == "" {
	//	userData.Password = currentUserData.Password
	//}
	//
	//currentUserData = userData

	if convertionError != nil {
		return convertionError, models.JsonStruct{}
	}

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
		photoId, _ := stngR.settingsDB.UploadPhoto(jsonData.Photo)

		currentUserData.Photo = photoId
	}

	if jsonData.Telephone != "" {
		currentUserData.Telephone = jsonData.Telephone
	}

	if jsonData.Email != "" {
		currentUserData.Email = jsonData.Email
	}

	stngR.settingsDB.UploadSettings(id, currentUserData)

	sendData := make(map[string]interface{})

	sendData["user"], _ = stngR.settingsDB.GetUserProfileSettingsById(id)

	return nil, models.JsonStruct{Body: sendData}
}

func NewSetUseCaseRealisation(userDB UserRep.UserRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation, log *zap.SugaredLogger) SettingsUseCaseRealisation {
	return SettingsUseCaseRealisation{
		settingsDB: userDB,
		sessionDB:  sesDB,
		logger:     log,
	}
}
