package usecase

import (
	errors2 "errors"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	mock "main/internal/friends/usecase/mock"
	"main/internal/models"
	"main/internal/tools/errors"
	"math/rand"
	"testing"
)

func TestFriendUseCaseRealisation_GetAllFriends(t *testing.T) {

	uVal := uuid.NewV4()

	ctrl := gomock.NewController(t)
	fRepoMock := mock.NewMockFriendRepository(ctrl)

	fTest := NewFriendUseCaseRealisation(fRepoMock)

	errs := []error{nil, errors.FailReadFromDB}

	retVal := []models.FriendLandingInfo{models.FriendLandingInfo{
		Name:    "[e[",
		Surname: "]e]",
		Photo:   "]e]",
		Login:   "[e[e",
	}, models.FriendLandingInfo{}}

	for iter, _ := range errs {

		randLogin := uVal.String()

		fRepoMock.EXPECT().GetAllFriendsByLogin(randLogin).Return(retVal, errs[iter])

		val, err := fTest.GetAllFriends(randLogin)

		if val[0] != retVal[0] && err != errs[iter] {
			t.Error("got : ", err, val, " expected :", retVal[iter], errs[iter])
		}
	}
}

func TestFriendUseCaseRealisation_AddFriend(t *testing.T) {

	uVal := uuid.NewV4()

	ctrl := gomock.NewController(t)
	fRepoMock := mock.NewMockFriendRepository(ctrl)

	fTest := NewFriendUseCaseRealisation(fRepoMock)

	friendErr := []error{nil, nil, errors.InvalidCookie}
	addErr := []error{nil, errors.FailAddFriend, nil}
	expectErr := []error{nil, errors.FailAddFriend, errors.FailAddFriend}

	for iter, _ := range expectErr {

		friendLogin := uVal.String()
		uId := rand.Int()
		fId := rand.Int()

		if friendErr[iter] == nil {
			fRepoMock.EXPECT().AddFriend(uId, fId).Return(addErr[iter])

		}
		fRepoMock.EXPECT().GetFriendIdByLogin(friendLogin).Return(fId, friendErr[iter])

		if fTest.AddFriend(uId, friendLogin) != expectErr[iter] {
			t.Error("expected :", expectErr[iter], "  iter :", iter)
		}
	}
}

func TestFriendUseCaseRealisation_DeleteFriend(t *testing.T) {

	uVal := uuid.NewV4()

	ctrl := gomock.NewController(t)
	fRepoMock := mock.NewMockFriendRepository(ctrl)

	fTest := NewFriendUseCaseRealisation(fRepoMock)

	friendErr := []error{nil, nil, errors.InvalidCookie}
	deleteErr := []error{nil, errors.FailDeleteFriend, nil}
	expectErr := []error{nil, errors.FailDeleteFriend, errors.FailDeleteFriend}

	for iter, _ := range expectErr {

		friendLogin := uVal.String()
		uId := rand.Int()
		fId := rand.Int()

		if friendErr[iter] == nil {
			fRepoMock.EXPECT().DeleteFriend(uId, fId).Return(deleteErr[iter])
		}
		fRepoMock.EXPECT().GetFriendIdByLogin(friendLogin).Return(fId, friendErr[iter])

		if err := fTest.DeleteFriend(uId, friendLogin); err != expectErr[iter] {
			t.Error("expected :", expectErr[iter], "  iter :", iter, err)
		}
	}
}

func TestFriendUseCaseRealisation_GetUserLoginByCookie(t *testing.T) {

	uVal := uuid.NewV4()

	ctrl := gomock.NewController(t)
	fRepoMock := mock.NewMockFriendRepository(ctrl)

	fTest := NewFriendUseCaseRealisation(fRepoMock)

	customErr := errors2.New("new error")
	returnErr := []error{nil, errors.FailDeleteFriend, customErr}
	expectErr := []error{nil, errors.FailDeleteFriend, customErr}

	for iter, _ := range expectErr {
		friendLogin := uVal.String()
		uId := rand.Int()

		fRepoMock.EXPECT().GetUserLoginById(uId).Return(friendLogin, returnErr[iter])

		if gotV, gotE := fTest.GetUserLoginById(uId); gotV != friendLogin || gotE != expectErr[iter] {
			t.Error("expected :", expectErr[iter], "  iter :", iter)
		}
	}
}
