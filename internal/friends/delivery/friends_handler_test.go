package delivery

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	mock_friends "main/internal/friends/delivery/mock"
	"main/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFriendDeliveryRealisation_AddFriend(t *testing.T) {
	ctrl := gomock.NewController(t)
	lUseCase := mock_friends.NewMockFriendUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	friendD := NewFriendDelivery(logger, lUseCase)

	usersId := []int{-1, 1, 2}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusConflict}

	for iter, _ := range usersId {
		friendLogin := uuid.NewV4()

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().AddFriend(usersId[iter], friendLogin.String()).Return(likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/user/:id")
		c.SetParamNames("id")
		c.SetParamValues(friendLogin.String())
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, friendD.AddFriend(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_DeleteFriend(t *testing.T) {
	ctrl := gomock.NewController(t)
	lUseCase := mock_friends.NewMockFriendUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	friendD := NewFriendDelivery(logger, lUseCase)

	usersId := []int{-1, 1, 2}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusConflict}

	for iter, _ := range usersId {
		friendLogin := uuid.NewV4()

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().DeleteFriend(usersId[iter], friendLogin.String()).Return(likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/user/:id")
		c.SetParamNames("id")
		c.SetParamValues(friendLogin.String())
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, friendD.DeleteFriend(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_GetMainUserFriends(t *testing.T) {
	ctrl := gomock.NewController(t)
	lUseCase := mock_friends.NewMockFriendUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	friendD := NewFriendDelivery(logger, lUseCase)

	answer := []models.FriendLandingInfo{
		{
			Name:    "asdasd",
			Surname: "asdasd",
			Photo:   "asdasd",
			Login:   "asdasd",
		},
		{
			Name:    "123",
			Surname: "123",
			Photo:   "123",
			Login:   "123",
		},
	}

	usersId := []int{-1, 1, 2}
	friendBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusNotFound}

	for iter, _ := range usersId {
		login := uuid.NewV4()

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().GetUserLoginById(usersId[iter]).Return(login.String(), nil)
			lUseCase.EXPECT().GetAllFriends(login.String()).Return(answer, friendBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/friends")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, friendD.GetMainUserFriends(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_FriendList(t *testing.T) {
	ctrl := gomock.NewController(t)
	lUseCase := mock_friends.NewMockFriendUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	friendD := NewFriendDelivery(logger, lUseCase)

	answer := []models.FriendLandingInfo{
		{
			Name:    "asdasd",
			Surname: "asdasd",
			Photo:   "asdasd",
			Login:   "asdasd",
		},
		{
			Name:    "123",
			Surname: "123",
			Photo:   "123",
			Login:   "123",
		},
	}

	usersId := []int{-1, 1, 2}
	friendBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusOK, http.StatusOK, http.StatusNotFound}

	for iter, _ := range usersId {
		login := uuid.NewV4()

		lUseCase.EXPECT().GetAllFriends(login.String()).Return(answer, friendBehaviour[iter])

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/friends/:id")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])
		c.SetParamNames("id")
		c.SetParamValues(login.String())

		if assert.NoError(t, friendD.FriendList(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}
