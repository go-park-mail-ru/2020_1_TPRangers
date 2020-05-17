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

func (groupD GroupDeliveryRealisation) LeaveTheGroup(rwContext echo.Context) error {
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
	err = groupD.groupLogic.LeaveTheGroup(userId, groupID)

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

func (groupD GroupDeliveryRealisation) CreatePostInGroup(rwContext echo.Context) error {
	uId := rwContext.Get("REQUEST_ID").(string)
	userId := rwContext.Get("user_id").(int)

	if userId == -1 {

		groupD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}


	newPost := new(models.Post)

	err := rwContext.Bind(newPost)

	if err != nil {
		groupD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	groupID, err := strconv.Atoi(rwContext.Param("id"))
	err = groupD.groupLogic.CreatePostInGroup(userId, groupID, *newPost)

	var responseStatus int
	switch err {
	case errors.DontHavePermission:
		responseStatus = http.StatusForbidden
	case errors.FailSendToDB:
		responseStatus = http.StatusConflict
	}

	if err != nil {
		groupD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", responseStatus),
		)

		return rwContext.JSON(responseStatus, models.JsonStruct{Err: err.Error()})
	}

	groupD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)
}
func (groupD GroupDeliveryRealisation) GetGroupProfile(rwContext echo.Context) error {
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

	groupID, err := strconv.Atoi(rwContext.Param("id"))
	groupData, err := groupD.groupLogic.GetGroupProfile(userId, groupID)

	if err != nil {
		groupD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusNotFound),
		)

		return rwContext.JSON(http.StatusNotFound, models.JsonStruct{Err: err.Error()})
	}

	groupD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, groupData)
}

func (groupD GroupDeliveryRealisation) GetGroupFeeds(rwContext echo.Context) error {
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

	groupID, err := strconv.Atoi(rwContext.Param("id"))
	groupFeed, err := groupD.groupLogic.GetGroupFeeds(userId, groupID)

	if err != nil {
		groupD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusNotFound),
		)

		return rwContext.JSON(http.StatusNotFound, models.JsonStruct{Err: err.Error()})
	}

	groupD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, groupFeed)
}

func (groupD GroupDeliveryRealisation) GetUserGroupsList(rwContext echo.Context) error {
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

	groupsList, err := groupD.groupLogic.GetUserGroupsList(userId)

	if err != nil {
		groupD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusNotFound),
		)

		return rwContext.JSON(http.StatusNotFound, models.JsonStruct{Err: err.Error()})
	}

	groupD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, groupsList)
}

func (groupD GroupDeliveryRealisation) SearchAllGroups(rwContext echo.Context) error {
	rId := rwContext.Get("REQUEST_ID").(string)
	userId := rwContext.Get("user_id").(int)
	valueOfSearch := rwContext.Param("value")

	if userId == -1 {
		groupD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	jsonAnswer, err := groupD.groupLogic.SearchAllGroups(userId, valueOfSearch)

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
	return rwContext.JSON(http.StatusOK, jsonAnswer)
}



func NewGroupDelivery(log *zap.SugaredLogger, groupRealisation groups.GroupUseCase) GroupDeliveryRealisation {
	return GroupDeliveryRealisation{groupLogic: groupRealisation, logger: log}
}

func (groupD GroupDeliveryRealisation) InitHandlers(server *echo.Echo) {
	server.POST("/api/v1/group/:id/join", groupD.JoinTheGroup)
	server.DELETE("/api/v1/group/:id/join", groupD.LeaveTheGroup)
	server.POST("/api/v1/group/create", groupD.CreateGroup)
	server.POST("/api/v1/group/:id/post/create", groupD.CreatePostInGroup)
	server.GET("/api/v1/group/:id/profile", groupD.GetGroupProfile)
	server.GET("/api/v1/group/:id/feed", groupD.GetGroupFeeds)
	server.GET("/api/v1/group/list", groupD.GetUserGroupsList)
	server.GET("/api/v1/group/search/:value", groupD.SearchAllGroups)

}

