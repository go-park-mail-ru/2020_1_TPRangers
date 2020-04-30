package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"io/ioutil"
	"main/internal/chats"
	"main/internal/tools/errors"
	"main/models"
	"net/http"
)

type ChatsDelivery struct {
	chatsLogic chats.ChatUseCase
	logger     *zap.SugaredLogger
}

func NewChatsDelivery(log *zap.SugaredLogger, chatsRealisation chats.ChatUseCase) ChatsDelivery {
	return ChatsDelivery{logger: log, chatsLogic: chatsRealisation}
}

func (CD ChatsDelivery) CreateChat(rwContext echo.Context) error {
	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	newChat := new(models.NewChatUsers)

	b, err := ioutil.ReadAll(rwContext.Request().Body)
	defer rwContext.Request().Body.Close()

	if err != nil {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	err = newChat.UnmarshalJSON(b)

	if err != nil {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	err = CD.chatsLogic.CreateChat(*newChat, userId)

	if err != nil {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	CD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)
}

func (CD ChatsDelivery) ExitChat(rwContext echo.Context) error {
	uId := rwContext.Get("REQUEST_ID").(string)
	chatId := rwContext.Param("id")

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err := CD.chatsLogic.ExitChat(chatId, userId)

	if err != nil {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	CD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)
}

func (CD ChatsDelivery) GetChatMessages(rwContext echo.Context) error {
	uId := rwContext.Get("REQUEST_ID").(string)
	chatId := rwContext.Param("id")

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	chatInfo, err := CD.chatsLogic.GetChatMessages(chatId, userId)

	if err != nil {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	CD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, chatInfo)

}

func (CD ChatsDelivery) GetAllChats(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	allChats, err := CD.chatsLogic.GetAllChats(userId)

	if err != nil {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	CD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, allChats)
}

func (CD ChatsDelivery) InitHandlers(server *echo.Echo) {

	server.POST("/api/v1/chats", CD.CreateChat)
	server.DELETE("/api/v1/chats/:id", CD.ExitChat)

	server.GET("/api/v1/chats/:id", CD.GetChatMessages)
	server.GET("/api/v1/chats", CD.GetAllChats)
}
