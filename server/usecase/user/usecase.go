package user

import ("github.com/labstack/echo"
		"../../models"
"../../errors"
)

type UserRealisation struct {
	userDB    REPOSITORYCASETYPE
	feedDB    REPOSITORYCASETYPE
}




func (ur UserRealisation) GetUser(rwContext echo.Context) (error , models.Feed , models.Settings){

	login := rwContext.Param("id")

	userData , existError := ur.userDB.GetUserDataByLogin(login)

	if existError != nil {
		return errors.NotExist , models.Feed{} , models.Settings{}
	}



}
