package delivery

import (
	"fmt"
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

	fmt.Println(rwContext.Request())

	uId := rwContext.Get("REQUEST_ID").(string)
	token := rwContext.Param("token")

	fmt.Println(len(token))

	userId, err := SD.socketLogic.CheckToken(token)

	fmt.Println(userId)

	if userId == -1 || err != nil {
		SD.logger.Debug(
			zap.String("ID", uId),
			zap.String("TOKEN", errors.InvalidToken.Error()),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.InvalidToken.Error()})
	}

	conn, _, _, err := ws.UpgradeHTTP(rwContext.Request(), rwContext.Response())

	if err != nil {
		SD.logger.Debug(
			zap.String("ID", uId),
			zap.Int("STAT", 2),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusNotModified),
		)

		return rwContext.JSON(http.StatusNotModified, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err = SD.socketLogic.AddToConnectionPool(conn, userId)

	if err != nil {
		SD.logger.Debug(
			zap.String("ID", uId),
			zap.Int("STAT", 3),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	return rwContext.NoContent(http.StatusSwitchingProtocols)
}

func (SD SocketDelivery) InitHandlers(server *echo.Echo) {
	server.GET("/api/v1/ws/:token", SD.UpgradeToSocket)
}
