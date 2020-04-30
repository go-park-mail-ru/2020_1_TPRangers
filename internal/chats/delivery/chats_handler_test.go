package delivery

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/mocks"
	"main/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)
func TestChatsDelivery_CreateChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	cUseCase := mock.NewMockChatUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	chatTest := NewChatsDelivery(logger, cUseCase)

	users := []int{2, -1, 3, 4, 5}
	chatData := `{"chatName" : "nn"}`
	wrongChat := `{"haha":"lol"}`
	chats := []string{chatData, chatData, wrongChat, "asdasdasd", chatData}
	mockErr := []error{nil , nil , nil , nil , errors.New("oh shit here we go again")}
	expectedBehaviour := []int{http.StatusOK, http.StatusUnauthorized, http.StatusConflict, http.StatusConflict, http.StatusConflict}

	for iter, _ := range users {
		if expectedBehaviour[iter] == http.StatusOK || iter == len(users) -1 {
			newChat := new(models.NewChatUsers)
			if chats[iter] != wrongChat {
				newChat.ChatName = "nn"
			}
			cUseCase.EXPECT().CreateChat(*newChat,users[iter]).Return(mockErr[iter])
		}


		e := echo.New()
		var req *http.Request
		if iter == 2 {
			req = httptest.NewRequest(echo.POST, "/", nil)
		} else {
			req = httptest.NewRequest(echo.POST, "/", strings.NewReader(chats[iter]))
		}
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/chats")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", users[iter])

		if assert.NoError(t , chatTest.CreateChat(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}
	}
}

func TestChatsDelivery_ExitChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	cUseCase := mock.NewMockChatUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	chatTest := NewChatsDelivery(logger, cUseCase)


	users := []int{2, -1, 3}
	exitErr := []error{nil , nil , errors.New("smth")}
	expectedBehaviour := []int{http.StatusOK, http.StatusUnauthorized, http.StatusConflict}


	for iter , _ := range expectedBehaviour {
		chatId := uuid.NewV4()
		if expectedBehaviour[iter] != http.StatusUnauthorized {
			cUseCase.EXPECT().ExitChat(chatId.String() , users[iter]).Return(exitErr[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/chats/:id")
		c.SetParamNames("id")
		c.SetParamValues(chatId.String())
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", users[iter])

		if assert.NoError(t , chatTest.ExitChat(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}
	}
}

func TestChatsDelivery_GetChatMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	cUseCase := mock.NewMockChatUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	chatTest := NewChatsDelivery(logger, cUseCase)


	users := []int{2, -1, 3}
	exitErr := []error{nil , nil , errors.New("smth")}
	expectedBehaviour := []int{http.StatusOK, http.StatusUnauthorized, http.StatusConflict}


	for iter , _ := range expectedBehaviour {
		chatId := uuid.NewV4()
		defInfo := models.ChatAndMessages{}
		if expectedBehaviour[iter] != http.StatusUnauthorized {
			cUseCase.EXPECT().GetChatMessages(chatId.String() , users[iter]).Return(defInfo,exitErr[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/chats/:id")
		c.SetParamNames("id")
		c.SetParamValues(chatId.String())
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", users[iter])

		if assert.NoError(t , chatTest.GetChatMessages(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}
	}
}

func TestChatsDelivery_GetAllChats(t *testing.T) {
	ctrl := gomock.NewController(t)
	cUseCase := mock.NewMockChatUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	chatTest := NewChatsDelivery(logger, cUseCase)


	users := []int{2, -1, 3}
	exitErr := []error{nil , nil , errors.New("smth")}
	expectedBehaviour := []int{http.StatusOK, http.StatusUnauthorized, http.StatusConflict}


	for iter , _ := range expectedBehaviour {
		defInfo := make([]models.Chat,3,3)
		if expectedBehaviour[iter] != http.StatusUnauthorized {
			cUseCase.EXPECT().GetAllChats(users[iter]).Return(defInfo,exitErr[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/chats")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", users[iter])

		if assert.NoError(t , chatTest.GetAllChats(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}
	}
}