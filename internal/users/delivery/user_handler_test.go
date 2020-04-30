package delivery

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/internal/csrf"
	"main/internal/models"
	mock_users "main/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
	expectedBehaviour := []int{http.StatusOK, http.StatusOK, http.StatusOK}

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

func TestFriendDeliveryRealisation_Register(t *testing.T) {
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
	expectedBehaviour := []int{http.StatusInternalServerError, http.StatusInternalServerError, http.StatusInternalServerError}

	for iter, _ := range usersId {
		regData := models.Register{}
		if expectedBehaviour[iter] != http.StatusUnauthorized {
			answer := ""
			lUseCase.EXPECT().Register(regData).Return(answer, likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/registration")
		c.Set("REQUEST_ID", "123")

		if assert.NoError(t, userD.Register(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_UploadSettings(t *testing.T) {
	ctrl := gomock.NewController(t)
	lUseCase := mock_users.NewMockUserUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	friendD := NewUserDelivery(logger, lUseCase)

	usersId := []int{-1, 1, 2}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusConflict, http.StatusConflict}

	for iter, _ := range usersId {
		settings := models.Settings{}

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().UploadSettings(usersId[iter], settings).Return(settings, likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.PUT, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/settings")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, friendD.UploadSettings(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_Profile(t *testing.T) {
	ctrl := gomock.NewController(t)
	lUseCase := mock_users.NewMockUserUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	friendD := NewUserDelivery(logger, lUseCase)

	usersId := []int{-1, 1}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusUnauthorized}

	for iter, _ := range usersId {
		profile := models.MainUserProfileData{}

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().GetMainUserProfile(usersId[iter]).Return(profile, likeBehaviour[iter])
		}
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/profile")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, friendD.Profile(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_GetSettings(t *testing.T) {
	ctrl := gomock.NewController(t)
	aUseCase := mock_users.NewMockUserUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	feedD := NewUserDelivery(logger, aUseCase)
	usersId := []int{-1, 1, 2}
	userBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusInternalServerError}

	for iter, _ := range usersId {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			settings := models.Settings{}
			aUseCase.EXPECT().GetSettings(usersId[iter]).Return(settings, userBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("id", usersId[iter])
		c.SetPath("/api/v1/settings")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])
		if assert.NoError(t, feedD.GetSettings(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_SearchUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	aUseCase := mock_users.NewMockUserUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	valueOfSearch := "al"
	feedD := NewUserDelivery(logger, aUseCase)
	usersId := []int{-1, 1, 2}
	userBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusConflict}

	for iter, _ := range usersId {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			persons := make([]models.Person,2, 2)
			aUseCase.EXPECT().SearchUsers(usersId[iter], valueOfSearch).Return(persons, userBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/users/search/:value")
		c.SetParamNames("value")
		c.SetParamValues(valueOfSearch)
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])
		if assert.NoError(t, feedD.SearchUsers(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_GetCsrf(t *testing.T) {
	ctrl := gomock.NewController(t)
	aUseCase := mock_users.NewMockUserUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	SessionID := "123"
	feedD := NewUserDelivery(logger, aUseCase)
	usersId := []int{-1, 1, 2}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusUnauthorized, http.StatusUnauthorized}

	for iter, _ := range usersId {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			csrf.Tokens.Create(SessionID, 900+time.Now().Unix())
		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/csrf")
		c.Set("session_id", SessionID)
		if assert.NoError(t, feedD.GetCsrf(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	lUseCase := mock_users.NewMockUserUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	userD := NewUserDelivery(logger, lUseCase)
	cookeVal := "12345"
	usersId := []int{-1, 1, 2}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusConflict, http.StatusConflict, http.StatusConflict}

	for iter, _ := range usersId {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			user := models.Auth{}
			lUseCase.EXPECT().Login(user).Return(cookeVal,likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/login")
		c.Set("REQUEST_ID", "123")
		if assert.NoError(t, userD.Login(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}




