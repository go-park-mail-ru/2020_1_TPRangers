package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	stoken "main/internal/socket_token"
	"main/internal/tools/errors"
	"main/models"
	"net/http"
)

type TokenDelivery struct {
	tokenUseCase stoken.TokeUseCase
	logger       *zap.SugaredLogger
}

func NewTokenDelivery(logger *zap.SugaredLogger, tokenUseCase stoken.TokeUseCase) TokenDelivery {
	return TokenDelivery{
		logger:       logger,
		tokenUseCase: tokenUseCase,
	}
}

func (TD TokenDelivery) TokenSetup(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		TD.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", errors.CookieExpired.Error()),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	token, err := TD.tokenUseCase.CreateNewToken(userId)

	if err != nil {
		TD.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", err.Error()),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	TD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, models.Token{Token: token})
}

func (TD TokenDelivery) InitHandlers(server *echo.Echo) {
	server.GET("/api/v1/ws", TD.TokenSetup)
}
