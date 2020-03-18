package user

import (
	"../../usecase"
	"../../usecase/user"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type UserDeliveryRealisation struct {
	userLogic usecase.UserUseCase
}

func (userD UserDeliveryRealisation) GetUser(rwContext echo.Context) error {
	uniqueID, _ := uuid.NewV4()
	//start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	err, answerJson := userD.userLogic.GetUser(rwContext)

	if err != nil {
		return rwContext.JSON(http.StatusNotFound, answerJson)
	}

	return rwContext.JSON(http.StatusOK, answerJson)
}

func (userD UserDeliveryRealisation) Profile(rwContext echo.Context) error {
	uniqueID, _ := uuid.NewV4()
	//start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	err, answerJson := userD.userLogic.Profile(rwContext)

	if err != nil {
		return rwContext.JSON(http.StatusUnauthorized, answerJson)
	}

	return rwContext.JSON(http.StatusOK, answerJson)
}

func NewUserDelivery() UserDeliveryRealisation{
	userHandler := user.UserUseCaseRealisation{}
	return UserDeliveryRealisation{userLogic:userHandler}
}
