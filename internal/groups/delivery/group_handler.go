package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/groups"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
	"strconv"
)
type GroupDeliveryRealisation struct {
	groupLogic groups.GroupUseCase
	logger    *zap.SugaredLogger
}

func (groupD GroupDeliveryRealisation) JoinTheGroup(rwContext echo.Context) error {
	uId := rwContext.Get("REQUEST_ID").(string)
	userId := rwContext.Get("user_id").(int)

	if userId == -1 {

		groupD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	groupID, err := strconv.Atoi(rwContext.Param("id"))
	err = groupD.groupLogic.JoinTheGroup(userId, groupID)

	if err != nil {
		groupD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	groupD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)
}

func (groupD GroupDeliveryRealisation) CreateGroup(rwContext echo.Context) error {
	rId := rwContext.Get("REQUEST_ID").(string)
	userId := rwContext.Get("user_id").(int)
	if userId == -1 {
		groupD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}
	groupData := new(models.Group)

	err := rwContext.Bind(groupData)
	if err != nil {
		groupD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	err = groupD.groupLogic.CreateGroup(userId, *groupData)

	if err != nil {
		groupD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	groupD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)
}





func NewGroupDelivery(log *zap.SugaredLogger, groupRealisation groups.GroupUseCase) GroupDeliveryRealisation {
	return GroupDeliveryRealisation{groupLogic: groupRealisation, logger: log}
}

func (groupD GroupDeliveryRealisation) InitHandlers(server *echo.Echo) {
	server.POST("/api/v1/group/:id/join", groupD.JoinTheGroup)
	server.POST("/api/v1/group/create", groupD.CreateGroup)

}

