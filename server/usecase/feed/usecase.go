package feed

import (
	"../../errors"
	"../../models"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type FeedUseCaseRealisation struct {
	feedDB    REPOSITORYCASETYPE
	sessionDB REPOSITORYSESSIONTYPE
}

func (feedR FeedUseCaseRealisation) Feed(rwContext echo.Context) (error, models.JsonStruct) {

	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {
		return errors.CookieExpired, models.JsonStruct{Err:errors.CookieExpired.Error()}
	}

	login, err := feedR.sessionDB.GetUserByCookie(cookie.Value)

	if err != nil {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)
		return errors.InvalidCookie, models.JsonStruct{Err : errors.InvalidCookie.Error()}
	}

	sendData := make(map[string]interface{})

	sendData["feed"] = feedR.feedDB.GetUserFeed(login, 30)

	return nil, models.JsonStruct{Body: sendData}
}
