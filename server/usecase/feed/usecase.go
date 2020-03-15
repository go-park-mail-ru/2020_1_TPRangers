package feed

import (
	"../../models"
	"../../errors"
	"github.com/labstack/echo"
	"net/http"
)

type FeedRealisation struct {
	feedDB    REPOSITORYCASETYPE
	sessionDB REPOSITORYSESSIONTYPE
}

func (feed FeedRealisation) Feed(rwContext echo.Context) (error, models.Feed) {

	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie{
		return errors.CookieExpired , models.Feed{}
	}

	login , err := feed.sessionDB.GetUserByCookie(cookie.Value)

	if err != nil {
		return errors.InvalidCookie , models.Feed{}
	}



}
