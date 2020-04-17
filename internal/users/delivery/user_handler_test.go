package delivery

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	mock_users "main/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFriendDeliveryRealisation_Logout(t *testing.T) {
	ctrl := gomock.NewController(t)
	lUseCase := mock_users.NewMockUserUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	userD := NewUserDelivery(logger, lUseCase)

	usersId := []int{-1, 1, 2}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusOK, http.StatusOK,  http.StatusOK}

	for iter, _ := range usersId {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().Logout(usersId[iter]).Return(likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/login")
		c.SetParamNames("id")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, userD.Logout(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}
