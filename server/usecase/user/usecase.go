package user

import (
	"../../errors"
	"../../models"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type UserUseCaseRealisation struct {
	userDB REPOSITORYCASETYPE
	feedDB REPOSITORYCASETYPE
}

func (userR UserUseCaseRealisation) GetUser(rwContext echo.Context) (error, models.JsonStruct) {

	login := rwContext.Param("id")

	userData, existError := userR.userDB.GetUserDataByLogin(login)

	if existError != nil {
		return errors.NotExist, models.JsonStruct{Err: errors.NotExist.Error()}
	}

	sendData := make(map[string]interface{})

	sendData["feed"] = userR.feedDB.GetUserFeed(login, 30)
	sendData["user"] = userData

	return nil, models.JsonStruct{Body: sendData}

}

func (userR UserUseCaseRealisation) Profile(rwContext echo.Context) (error, models.JsonStruct) {
	cookie, err := rwContext.Cookie("session_id")

	if err == http.ErrNoCookie {
		return errors.CookieExpired, models.JsonStruct{Err: errors.CookieExpired.Error()}
	}

	login, err := userR.userDB.GetUserByCookie(cookie.Value)

	if err != nil {

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		return errors.InvalidCookie, models.JsonStruct{Err: errors.CookieExpired.Error()}
	}

	sendData := make(map[string]interface{})
	sendData["user"], _ = userR.userDB.GetUserDataByLogin(login)

	return nil, models.JsonStruct{Body: sendData}
}
