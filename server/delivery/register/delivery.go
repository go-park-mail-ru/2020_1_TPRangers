package register

import (
	"../../errors"
	"../../usecase"
	"../../usecase/register"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type RegisterDeliveryRealisation struct {
	registerLogic usecase.RegisterUseCase
}

func (regDev RegisterDeliveryRealisation) Register(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	//start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	err := regDev.registerLogic.Register(rwContext)

	switch err {
	case errors.AlreadyExist:
		return rwContext.NoContent(http.StatusConflict)
	}

	if err != nil {
		return rwContext.NoContent(http.StatusInternalServerError)
	}

	return rwContext.NoContent(http.StatusOK)
}


func NewRegisterDelivery() RegisterDeliveryRealisation {
	regHandler := register.RegisterUseCaseRealisation{}
	return RegisterDeliveryRealisation{registerLogic:regHandler}
}
