package register

import (
	"../../errors"
	"../../usecase"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type RegisterDelivery struct {
	registerLogic usecase.RegisterUseCase
}

func (rDev RegisterDelivery) Register(rwContext echo.Context) error {
	uniqueID, _ := uuid.NewV4()
	//start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	err, cookieValue := rDev.registerLogic.Register(rwContext)

	switch err {
	case errors.AlreadyExist:
		return rwContext.NoContent(http.StatusConflict)
	}

	if err != nil {
		return rwContext.NoContent(http.StatusInternalServerError)
	}

	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(12 * time.Hour),
	}
	rwContext.SetCookie(&cookie)

	return rwContext.NoContent(http.StatusOK)
}
