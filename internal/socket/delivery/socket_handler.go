package delivery

import (
	"github.com/gobwas/ws"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/models"
	"main/internal/socket"
	"main/internal/tools/errors"
	"net/http"
)

type SocketDelivery struct {
	socketLogic socket.SocketUseCase
	logger      *zap.SugaredLogger
}

func NewSocketDelivery(logger *zap.SugaredLogger, sLogic socket.SocketUseCase) SocketDelivery {
	return SocketDelivery{
		socketLogic: sLogic,
		logger:      logger,
	}
}

func (SD SocketDelivery) UpgradeToSocket(rwContext echo.Context) error {

	userId := rwContext.Get("user_id").(int)
	uId := rwContext.Get("REQUEST_ID").(string)

	if userId == -1 {

		SD.logger.Debug(
			zap.String("ID", uId),
			zap.Int("STAT",1),
			zap.String("COOKIE", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	conn, _, _, err := ws.UpgradeHTTP(rwContext.Request(), rwContext.Response())

	if err != nil {
		SD.logger.Debug(
			zap.String("ID", uId),
			zap.Int("STAT",2),
			zap.String("ERROR" , err.Error()),
			zap.Int("ANSWER STATUS", http.StatusNotModified),
		)

		return rwContext.JSON(http.StatusNotModified, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err = SD.socketLogic.AddToConnectionPool(conn, userId)

	if err != nil {
		SD.logger.Debug(
			zap.String("ID", uId),
			zap.Int("STAT",3),
			zap.String("ERROR" , err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	return rwContext.NoContent(http.StatusSwitchingProtocols)
}

func (SD SocketDelivery) InitHandlers(server *echo.Echo) {
	server.GET("/ws", SD.UpgradeToSocket)
}
