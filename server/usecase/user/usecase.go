package user

import (
	"../../errors"
	"../../models"
	"github.com/labstack/echo"
)

type UserRealisation struct {
	userDB REPOSITORYCASETYPE
	feedDB REPOSITORYCASETYPE
}

func (ur UserRealisation) GetUser(rwContext echo.Context) (error, models.JsonStruct) {

	login := rwContext.Param("id")

	userData, existError := ur.userDB.GetUserDataByLogin(login)

	if existError != nil {
		return errors.NotExist, models.JsonStruct{}
	}

	sendData := make(map[string]interface{})

	sendData["feed"] = ur.feedDB.GetUserFeed(login, 30)
	sendData["user"] = userData

	return nil, models.JsonStruct{Body: sendData}

}
