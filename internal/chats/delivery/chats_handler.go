package delivery

import (
	"fmt"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/chats"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
	"strconv"
)

type ChatsDelivery struct {
	chatsLogic chats.ChatUseCase
	logger    *zap.SugaredLogger
}

func NewChatsDelivery(log * zap.SugaredLogger , chatsRealisation chats.ChatUseCase) ChatsDelivery {
	return ChatsDelivery{logger: log , chatsLogic: chatsRealisation}
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

	err_Bind := rwContext.Bind(newChat)
	fmt.Println(err_Bind)
	fmt.Println(*newChat)

	err := CD.chatsLogic.CreateChat(*newChat , userId)

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
	chatId, err := strconv.ParseInt(rwContext.Param("id"), 10 , 64)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err = CD.chatsLogic.ExitChat(chatId , userId)

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
	chatId, err := strconv.ParseInt(rwContext.Param("id"), 10 , 64)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		CD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	chatInfo , messages ,err := CD.chatsLogic.GetChatMessages(chatId , userId)

	chatAndMsgs := new(models.ChatAndMessages)

	chatAndMsgs.ChatInfo = chatInfo
	chatAndMsgs.ChatMessages = messages

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

	return rwContext.JSON(http.StatusOK, *chatAndMsgs)

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

	chats , err := CD.chatsLogic.GetAllChats(userId)


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

	return rwContext.JSON(http.StatusOK, chats)

}

func (CD ChatsDelivery) InitHandlers(server *echo.Echo) {

	server.POST("/api/v1/chats", CD.CreateChat )
	server.DELETE("/api/v1/chats/:id", CD.ExitChat)

	server.GET("/api/v1/chats/:id", CD.GetChatMessages)
	server.GET("/api/v1/chats", CD.GetAllChats)
}
