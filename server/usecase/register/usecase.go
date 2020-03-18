package register

import (
	".."
	"../../errors"
	"../../models"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type RegisterUseCaseRealisation struct {
	registerDB REPOSITORYCASETYPE
	sessionDB  REPOSITORYSESSIONTYPE
}

func (regR RegisterUseCaseRealisation) Register(rwContext echo.Context) error {

	jsonData, convertionError := usecase.GetDataFromJson("reg", rwContext)

	if convertionError != nil {
		return convertionError
	}

	userData := jsonData.(*models.Register)

	login := userData.Email

	if regR.registerDB.CheckExistnig(login) == false {
		return errors.AlreadyExist
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

	regR.registerDB.AddNewUser(data)

	id := regR.registerDB.GetIdByLogin(login)
	info, _ := uuid.NewV4()
	cookieValue := info.String()
	regR.sessionDB.SetCookie(id, cookie)

	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(12 * time.Hour),
	}
	rwContext.SetCookie(&cookie)

	return nil

}
