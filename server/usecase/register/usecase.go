package register

import (
	".."
	"../../errors"
	"../../models"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

type RegisterRealisation struct {
	registerDB REPOSITORYCASETYPE
	sessionDB  REPOSITORYSESSIONTYPE
}

func (reg RegisterRealisation) Register(rwContext echo.Context) (error, string) {

	jsonData, convertionError := usecase.GetDataFromJson("reg", rwContext)

	if convertionError != nil {
		return convertionError, ""
	}

	userData := jsonData.(*models.Register)

	login := userData.Email

	if reg.registerDB.CheckExistnig(login) == false {
		return errors.AlreadyExist, ""
	}

	data := models.User{
		Login:     login,
		Telephone: userData.Phone,
		Email:     login,
		Name:      userData.Name,
		Password:  userData.Password,
		Surname:   userData.Surname,
		Date:      userData.Date,
		Photo:     "default",
	}

	reg.registerDB.AddNewUser(data)

	id := reg.registerDB.GetIdByLogin(login)
	info, _ := uuid.NewV4()
	cookie := info.String()
	reg.sessionDB.SetCookie(id, cookie)

	return nil, cookie

}
