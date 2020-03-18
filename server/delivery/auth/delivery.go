package auth

import (
	"../../errors"
	"../../usecase"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	"net/http"
	"../../usecase/auth"
)

type LoginDeliveryRealisation struct {
	loginLogic usecase.AuthUseCase
}

func (logD LoginDeliveryRealisation) Login(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	//start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	err := logD.loginLogic.Login(rwContext)

	switch err {
	case errors.WrongLogin, errors.WrongPassword:
		return rwContext.NoContent(http.StatusUnauthorized)
	}

	if err != nil {
		return rwContext.NoContent(http.StatusUnauthorized)
	}

	return rwContext.NoContent(http.StatusOK)
}

func (logD LoginDeliveryRealisation) Logout(rwContext echo.Context) error {
	logD.loginLogic.Logout(rwContext)
	return rwContext.NoContent(http.StatusOK)
}

func NewLoginDelivery() LoginDeliveryRealisation {
	authHandler := auth.AuthUseCaseRealisation{}
	return LoginDeliveryRealisation{loginLogic: authHandler}
}