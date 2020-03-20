package settings

import (
	"../../errors"
	"../../models"
	"../../repository"
	"../../usecase"
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

	sendData["user"], err = stngR.settingsDB.GetUserDataById(id)

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

	jsonData, convertionError := usecase.GetDataFromJson("data", rwContext)

	userData := jsonData.(models.User)

	//когда нам будут высылать закэшированные настройки
	//if userData.Password == "" {
	//	userData.Password = currentUserData.Password
	//}
	//
	//currentUserData = userData

	if convertionError != nil {
		return convertionError , models.JsonStruct{}
	}

	if userData.Login != "" {
		currentUserData.Login = userData.Login
	}

	if userData.Password != "" {
		currentUserData.Password = userData.Password
	}

	if userData.Date != "" {
		currentUserData.Date = userData.Date
	}

	if userData.Surname != ""{
		currentUserData.Surname = userData.Surname
	}

	if userData.Name != ""{
		currentUserData.Name = userData.Name
	}

	if userData.Photo != ""{
		currentUserData.Photo = userData.Photo
	}

	if userData.Telephone != ""{
		currentUserData.Telephone = userData.Telephone
	}

	if userData.Email != "" {
		currentUserData.Email = userData.Email
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
