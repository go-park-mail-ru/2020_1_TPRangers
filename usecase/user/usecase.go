package user

import (
	"../../errors"
	"../../models"
	"../../repository"
	FeedRep "../../repository/feed"
	UserRep "../../repository/user"
	SessRep "../../repository/cookie"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserUseCaseRealisation struct {
	userDB    repository.UserRepository
	feedDB    repository.FeedRepository
	sessionDB repository.CookieRepository
	logger    *zap.SugaredLogger
}

func (userR UserUseCaseRealisation) GetUser(rwContext echo.Context, uId string) (error, models.JsonStruct) {

	login := rwContext.Param("id")

	userData, existError := userR.userDB.GetUserProfileSettingsByLogin(login)

	if existError != nil {

		userR.logger.Debug(
			zap.String("ID", uId),
			zap.String("LOGIN", "user doesnt exist"),
		)

		return errors.NotExist, models.JsonStruct{Err: errors.NotExist.Error()}
	}

	userR.logger.Debug(
		zap.String("ID", uId),
		zap.String("LOGIN", login),
	)

	sendData := make(map[string]interface{})

	sendData["feed"], _ = userR.feedDB.GetUserFeedByEmail(login, 30)
	sendData["user"] = userData

	return nil, models.JsonStruct{Body: sendData}

}

func (userR UserUseCaseRealisation) Profile(rwContext echo.Context, uId string) (error, models.JsonStruct) {
	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {

		userR.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", "no-cookie"),
		)

		return errors.CookieExpired, models.JsonStruct{Err: errors.CookieExpired.Error()}
	}

	id, err := userR.sessionDB.GetUserIdByCookie(cookie.Value)

	if err != nil {

		userR.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", cookie.Value),
			zap.String("USER ID", "no such user"),
		)


		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		return errors.InvalidCookie, models.JsonStruct{Err: errors.CookieExpired.Error()}
	}

	userR.logger.Debug(
		zap.String("ID", uId),
		zap.String("COOKIE", cookie.Value),
		zap.Int("USER ID", id),
	)

	sendData := make(map[string]interface{})
	sendData["user"], _ = userR.userDB.GetUserProfileSettingsById(id)

	return nil, models.JsonStruct{Body: sendData}
}

func NewUserUseCaseRealisation(userDB UserRep.UserRepositoryRealisation, feedDB FeedRep.FeedRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation, log *zap.SugaredLogger) UserUseCaseRealisation {
	return UserUseCaseRealisation{
		userDB:    userDB,
		feedDB:    feedDB,
		sessionDB: sesDB,
		logger:    log,
	}
}
