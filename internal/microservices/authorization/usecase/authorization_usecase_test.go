package usecase

//import (
//	"bytes"
//	"context"
//	cr "crypto/rand"
//	"crypto/sha1"
//	"errors"
//	"github.com/golang/mock/gomock"
//	"golang.org/x/crypto/pbkdf2"
//	"main/internal/microservices/authorization/delivery"
//	"main/mocks"
//	"testing"
//)
//
//func TestCryptPasswordAndCheckPassword(t *testing.T) {
//
//	passes := []string{"love", "h8", "sex", "god"}
//
//	for iter, _ := range passes {
//
//		salt := make([]byte, 8)
//		cr.Read(salt)
//		cryptedPass := pbkdf2.Key([]byte(passes[iter]), salt, 4096, 32, sha1.New)
//		cryptedPass = append(salt, cryptedPass...)
//
//		currentCryptedPass := CryptPassword(passes[iter], salt)
//
//		if !bytes.Equal(currentCryptedPass, cryptedPass) {
//			t.Error("PASS CRYPT IS WRONG")
//			return
//		}
//
//		if !CheckPassword(passes[iter], currentCryptedPass) {
//			t.Error("PASS ENCRYPT IS WRONG")
//			return
//		}
//
//	}
//}
//
//func TestAuthorizationUseCaseRealisation_CheckSession(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	userDBMock := mock.NewMockUserRepository(ctrl)
//	sesDBMock := mock.NewMockCookieRepository(ctrl)
//
//	authUse := NewAuthorizationUseCaseRealisation(userDBMock, sesDBMock)
//	auth := &session.SessionData{
//		Cookies: "123",
//		Csrf:    "",
//	}
//	sesDBMock.EXPECT().GetUserIdByCookie("123").Return(1, nil)
//
//	ses, _ := authUse.CheckSession(context.Background(), auth)
//
//	if ses.UserId != int32(1) {
//		t.Error("error")
//	}
//
//	sesDBMock.EXPECT().GetUserIdByCookie("123").Return(0, nil)
//
//	ses, _ = authUse.CheckSession(context.Background(), auth)
//
//	if ses.UserId != int32(-1) {
//		t.Error("error")
//	}
//
//}
//
//func TestAuthorizationUseCaseRealisation_DeleteSession(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	userDBMock := mock.NewMockUserRepository(ctrl)
//	sesDBMock := mock.NewMockCookieRepository(ctrl)
//	customErr := errors.New("123")
//
//	authUse := NewAuthorizationUseCaseRealisation(userDBMock, sesDBMock)
//	auth := &session.SessionData{
//		Cookies: "123",
//		Csrf:    "",
//	}
//	sesDBMock.EXPECT().DeleteCookie("123").Return(customErr)
//
//	ses, err := authUse.DeleteSession(context.Background(), auth)
//
//	if ses.UserId != int32(-1) || err != customErr {
//		t.Error("error")
//	}
//
//}
