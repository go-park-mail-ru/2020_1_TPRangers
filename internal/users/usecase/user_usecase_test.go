package usecase

import (
	"bytes"
	cr "crypto/rand"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/pbkdf2"
	"main/internal/models"
	_error "main/internal/tools/errors"
	"main/mocks"
	"math/rand"
	"testing"
)

func TestCryptPasswordAndCheckPassword(t *testing.T) {

	passes := []string{"love", "h8", "sex", "god"}

	for iter, _ := range passes {

		salt := make([]byte, 8)
		cr.Read(salt)
		cryptedPass := pbkdf2.Key([]byte(passes[iter]), salt, 4096, 32, sha1.New)
		cryptedPass = append(salt, cryptedPass...)

		currentCryptedPass := CryptPassword(passes[iter], salt)

		if !bytes.Equal(currentCryptedPass, cryptedPass) {
			fmt.Println(cryptedPass)
			fmt.Println(currentCryptedPass)
			t.Error("PASS CRYPT IS WRONG")
			return
		}

		if !CheckPassword(passes[iter], currentCryptedPass) {
			t.Error("PASS ENCRYPT IS WRONG")
			return
		}

	}
}

func TestUserUseCaseRealisation_GetSettings(t *testing.T) {
	ctrl := gomock.NewController(t)
	uUseMock := mock.NewMockUserRepository(ctrl)

	uTest := NewUserUseCaseRealisation(uUseMock, nil, nil, nil)

	settingsErr := []error{nil, errors.New("smth wrong")}
	expectBehaviour := []error{nil, _error.FailReadFromDB}

	for iter, _ := range expectBehaviour {
		uId := rand.Int()
		userModel := models.Settings{
			Login:     "123",
			Telephone: "123",
			Password:  "123",
			Email:     "123",
			Name:      "123",
			Surname:   "123",
			Date:      "123",
			Photo:     "123",
		}

		uUseMock.EXPECT().GetUserProfileSettingsById(uId).Return(userModel, settingsErr[iter])

		if val, err := uTest.GetSettings(uId); !(val == userModel && err == expectBehaviour[iter]) {
			t.Error(iter, err, expectBehaviour[iter])
		}
	}

}

func TestUserUseCaseRealisation_UploadSettings(t *testing.T) {
	ctrl := gomock.NewController(t)
	uUseMock := mock.NewMockUserRepository(ctrl)

	uTest := NewUserUseCaseRealisation(uUseMock, nil, nil, nil)

	settingsErr := []error{nil, nil, _error.FailReadFromDB}
	getErr := []error{nil, _error.FailReadFromDB, nil}
	expectBehaviour := []error{nil, _error.FailReadFromDB, _error.FailReadFromDB}

	for iter, _ := range expectBehaviour {
		uId := rand.Int()
		oldUserModel := models.User{
			Login:     "123",
			Telephone: "123",
			Email:     "123",
			Name:      "123",
			Surname:   "123",
			Date:      "123",
			Photo:     123,
		}

		newUploadModel := models.User{
			Login:     "222",
			Telephone: "123",
			Email:     "123",
			Name:      "123",
			Surname:   "123",
			Date:      "123",
			Photo:     4,
		}

		newUserModel := models.Settings{
			Login:     "222",
			Telephone: "",
			Password:  "",
			Email:     "",
			Name:      "",
			Surname:   "",
			Date:      "",
			Photo:     "3",
		}

		returnUserModel := models.Settings{
			Login:     "222",
			Telephone: "123",
			Email:     "123",
			Name:      "123",
			Surname:   "123",
			Date:      "123",
			Photo:     "3",
		}

		uUseMock.EXPECT().GetUserDataById(uId).Return(oldUserModel, nil)
		uUseMock.EXPECT().UploadSettings(uId, newUploadModel).Return(settingsErr[iter])

		if newUserModel.Photo != "" {
			uUseMock.EXPECT().UploadProfilePhoto(newUserModel.Photo).Return(4, nil)
		}

		if settingsErr[iter] == nil {
			uUseMock.EXPECT().GetUserProfileSettingsById(uId).Return(returnUserModel, getErr[iter])
		} else {
			returnUserModel = models.Settings{}
		}

		if val, err := uTest.UploadSettings(uId, newUserModel); !(val == returnUserModel && err == expectBehaviour[iter]) {
			t.Error(iter, err, expectBehaviour[iter])
		}
	}

}

func TestUserUseCaseRealisation_GetUserLoginByCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	uUseMock := mock.NewMockUserRepository(ctrl)

	uTest := NewUserUseCaseRealisation(uUseMock, nil, nil, nil)

	errs := []error{nil, _error.FailReadFromDB}
	expectBehaviour := []error{nil, _error.FailReadFromDB}

	for iter, _ := range expectBehaviour {
		uId := rand.Int()
		login := uuid.NewV4()

		uUseMock.EXPECT().GetUserLoginById(uId).Return(login.String(), errs[iter])

		if currLog, err := uTest.GetUserLoginByCookie(uId); !(currLog == login.String() && err == expectBehaviour[iter]) {
			t.Error(iter, currLog, login.String(), err, expectBehaviour[iter])
		}
	}
}
