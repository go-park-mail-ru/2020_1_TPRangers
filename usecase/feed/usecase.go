package feed

import (
	"../../errors"
	"../../models"
	"../../repository"
	FeedRep "../../repository/feed"
	SessRep "../../repository/cookie"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type FeedUseCaseRealisation struct {
	feedDB    repository.FeedRepository
	sessionDB repository.CookieRepository
	logger *zap.SugaredLogger
}

func (feedR FeedUseCaseRealisation) Feed(rwContext echo.Context ,uId string) (error, models.JsonStruct) {

	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {
		feedR.logger.Info(
			zap.String("ID" , uId) ,
			zap.String("COOKIE" , "no-cookie"),
			)
		return errors.CookieExpired, models.JsonStruct{Err: errors.CookieExpired.Error()}
	}

	id, err := feedR.sessionDB.GetUserIdByCookie(cookie.Value)

	if err != nil {

		feedR.logger.Info(
			zap.String("ID" , uId) ,
			zap.String("USER-ID" , "no such user"),
		)

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)
		return errors.InvalidCookie, models.JsonStruct{Err: errors.InvalidCookie.Error()}
	}

	feedR.logger.Debug(
		zap.String("ID" , uId) ,
		zap.String("COOKIE" , cookie.Value),
		zap.Int("USER-ID" , id),
		"feed successfullty loaded",
	)

	sendData := make(map[string]interface{})

	sendData["feed"], err = feedR.feedDB.GetUserFeedById(id, 30)

	return err, models.JsonStruct{Body: sendData}
}

func NewFeedUseCaseRealisation(feedDB FeedRep.FeedRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation, log *zap.SugaredLogger) FeedUseCaseRealisation {
	return FeedUseCaseRealisation{
		feedDB:   feedDB,
		sessionDB: sesDB,
		logger:    log,
	}
}
