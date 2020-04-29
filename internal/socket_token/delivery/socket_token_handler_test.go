package delivery

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTokenDelivery_TokenSetup(t *testing.T) {
	ctrl := gomock.NewController(t)
	tokenUseMock := mock.NewMockTokeUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	token := NewTokenDelivery(logger, tokenUseMock)
	usersId := []int{-1, 1, 2}
	errs := []error{nil, errors.New("123"), nil}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusConflict, http.StatusOK}

	for iter, _ := range usersId {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			tokenUseMock.EXPECT().CreateNewToken(usersId[iter]).Return("HAHA" ,errs[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/ws")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, token.TokenSetup(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}
