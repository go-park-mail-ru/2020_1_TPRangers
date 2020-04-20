package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/models"
	"main/internal/tools/errors"
	tokens "main/internal/tools/token_generator"
	"net/http"
)

type TokenDelivery struct {
	logger *zap.SugaredLogger
}

func NewTokenDelivery(logger *zap.SugaredLogger) TokenDelivery {
	return TokenDelivery{
		logger: logger,
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

	token := tokens.CryptToken(userId)

	TD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, models.Token{Token: token})
}

func (TD TokenDelivery) InitHandlers(server *echo.Echo) {
	server.GET("/api/v1/ws", TD.TokenSetup)
}
