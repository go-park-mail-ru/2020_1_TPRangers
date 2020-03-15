package auth

import (
	"../../errors"
	"../../usecase"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type LoginDelivery struct {
	loginLogic usecase.AuthUseCase
}

func (log LoginDelivery) Login(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	//start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	err, cookieValue := log.loginLogic.Login(rwContext)

	switch err {
	case errors.WrongLogin, errors.WrongPassword:
		return rwContext.NoContent(http.StatusUnauthorized)
	}

	if err != nil {
		return rwContext.NoContent(http.StatusUnauthorized)
	}

	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(12 * time.Hour),
	}
	rwContext.SetCookie(&cookie)

	return rwContext.NoContent(http.StatusOK)
}

func (log LoginDelivery) Logout(rwContext echo.Context) error {
	log.loginLogic.Logout(rwContext)
	return rwContext.NoContent(http.StatusOK)
}
