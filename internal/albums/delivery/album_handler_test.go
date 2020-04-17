package delivery

//import (
//	"github.com/golang/mock/gomock"
//	"github.com/labstack/echo"
//	"github.com/pkg/errors"
//	uuid "github.com/satori/go.uuid"
//	"github.com/stretchr/testify/assert"
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//	mock_albums "main/internal/albums/delivery/mock"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestFriendDeliveryRealisation_AddFriend(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	lUseCase := mock_albums.NewMockAlbumUseCase(ctrl)
//	config := zap.NewDevelopmentConfig()
//	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	prLogger, _ := config.Build()
//	logger := prLogger.Sugar()
//	defer prLogger.Sync()
//	friendD := NewAlbumDelivery(logger, lUseCase)
//
//	usersId := []int{-1, 1, 2}
//	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
//	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusConflict}
//
//	for iter, _ := range usersId {
//		friendLogin := uuid.NewV4()
//
//		if expectedBehaviour[iter] != http.StatusUnauthorized {
//			lUseCase.EXPECT().AddFriend(usersId[iter], friendLogin.String()).Return(likeBehaviour[iter])
//		}
//
//		e := echo.New()
//		req := httptest.NewRequest(echo.POST, "/", nil)
//		rec := httptest.NewRecorder()
//		c := e.NewContext(req, rec)
//		c.SetPath("/api/v1/user/:id")
//		c.SetParamNames("id")
//		c.SetParamValues(friendLogin.String())
//		c.Set("REQUEST_ID", "123")
//		c.Set("user_id", usersId[iter])
//
//		if assert.NoError(t, friendD.AddFriend(c)) {
//			assert.Equal(t, expectedBehaviour[iter], rec.Code)
//		}
//
//	}
//}