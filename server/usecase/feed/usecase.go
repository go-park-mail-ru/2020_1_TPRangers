package feed

import (
	"../../errors"
	"../../models"
	"github.com/labstack/echo"
	"net/http"
)

type FeedRealisation struct {
	feedDB    REPOSITORYCASETYPE
	sessionDB REPOSITORYSESSIONTYPE
}

func (feed FeedRealisation) Feed(rwContext echo.Context) (error, models.JsonStruct) {

	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {
		return errors.CookieExpired, models.JsonStruct{}
	}

	login, err := feed.sessionDB.GetUserByCookie(cookie.Value)

	if err != nil {
		return errors.InvalidCookie, models.JsonStruct{}
	}

	sendData := make(map[string]interface{})

	sendData["feed"] = feed.feedDB.GetUserFeed(login, 30)

	return nil, models.JsonStruct{Body: sendData}
}
