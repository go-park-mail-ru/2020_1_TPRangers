package delivery

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/internal/models"
	"main/mocks"
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
	chats := []string{chatData, chatData, wrongChat, "", chatData}
	mockErr := []error{nil , nil , nil , nil , errors.New("oh shit here we go again")}
	expectedBehaviour := []int{http.StatusOK, http.StatusUnauthorized, http.StatusConflict, http.StatusConflict, http.StatusConflict}

	for iter, _ := range users {
		fmt.Println(users[iter], iter, expectedBehaviour[iter])
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
