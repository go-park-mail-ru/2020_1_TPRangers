package usecase

import (
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"main/internal/friends/usecase/mock"
	"main/internal/models"
	"main/internal/tools/errors"
	"math/rand"
	"testing"
)

func TestFriendUseCaseRealisation_GetAllFriends(t *testing.T) {

	uVal := uuid.NewV4()

	ctrl := gomock.NewController(t)
	fRepoMock := mock.NewMockFriendRepository(ctrl)
	cRepoMock := mock.NewMockCookieRepository(ctrl)

	fTest := NewFriendUseCaseRealisation(fRepoMock, cRepoMock)

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

		if val["friends"].([]models.FriendLandingInfo)[0] != retVal[0] && err != errs[iter] {
			t.Error("got : ", err, val, " expected :", retVal[iter], errs[iter])
		}
	}
}

func TestFriendUseCaseRealisation_AddFriend(t *testing.T) {

	uVal := uuid.NewV4()

	ctrl := gomock.NewController(t)
	fRepoMock := mock.NewMockFriendRepository(ctrl)
	cRepoMock := mock.NewMockCookieRepository(ctrl)

	fTest := NewFriendUseCaseRealisation(fRepoMock, cRepoMock)

	cookieErr := []error{nil, nil, errors.InvalidCookie}
	addErr := []error{nil, errors.FailAddFriend, nil}
	expectErr := []error{nil, errors.FailAddFriend, errors.InvalidCookie}

	for iter, _ := range expectErr {

		cookie := uVal.String()
		friendLogin := uVal.String()
		uId := rand.Int()
		fId := rand.Int()

		cRepoMock.EXPECT().GetUserIdByCookie(cookie).Return(uId, cookieErr[iter])

		if cookieErr[iter] == nil {
			fRepoMock.EXPECT().AddFriend(uId, fId).Return(addErr[iter])
			fRepoMock.EXPECT().GetFriendIdByLogin(friendLogin).Return(fId, nil)
		}

		if fTest.AddFriend(cookie, friendLogin) != expectErr[iter] {
			t.Error("expected :", expectErr[iter], "  iter :", iter)
		}
	}
}

func TestFriendUseCaseRealisation_DeleteFriend(t *testing.T) {

	uVal := uuid.NewV4()

	ctrl := gomock.NewController(t)
	fRepoMock := mock.NewMockFriendRepository(ctrl)
	cRepoMock := mock.NewMockCookieRepository(ctrl)

	fTest := NewFriendUseCaseRealisation(fRepoMock, cRepoMock)

	cookieErr := []error{nil, nil, errors.InvalidCookie}
	addErr := []error{nil, errors.FailDeleteFriend, nil}
	expectErr := []error{nil, errors.FailDeleteFriend, errors.InvalidCookie}

	for iter, _ := range expectErr {

		cookie := uVal.String()
		friendLogin := uVal.String()
		uId := rand.Int()
		fId := rand.Int()

		cRepoMock.EXPECT().GetUserIdByCookie(cookie).Return(uId, cookieErr[iter])

		if cookieErr[iter] == nil {
			fRepoMock.EXPECT().DeleteFriend(uId, fId).Return(addErr[iter])
			fRepoMock.EXPECT().GetFriendIdByLogin(friendLogin).Return(fId, nil)
		}

		if fTest.DeleteFriend(cookie, friendLogin) != expectErr[iter] {
			t.Error("expected :", expectErr[iter], "  iter :", iter)
		}
	}
}

func TestFriendUseCaseRealisation_GetUserLoginByCookie(t *testing.T) {

	uVal := uuid.NewV4()

	ctrl := gomock.NewController(t)
	fRepoMock := mock.NewMockFriendRepository(ctrl)
	cRepoMock := mock.NewMockCookieRepository(ctrl)

	fTest := NewFriendUseCaseRealisation(fRepoMock, cRepoMock)

	cookieErr := []error{nil, nil, errors.InvalidCookie}
	returnErr := []error{nil, errors.FailDeleteFriend, nil}
	expectErr := []error{nil, errors.FailDeleteFriend, nil}

	for iter, _ := range expectErr {

		cookie := uVal.String()
		friendLogin := uVal.String()
		uId := rand.Int()

		cRepoMock.EXPECT().GetUserIdByCookie(cookie).Return(uId, cookieErr[iter])

		fRepoMock.EXPECT().GetUserLoginById(uId).Return(friendLogin, returnErr[iter])

		if gotV, gotE := fTest.GetUserLoginByCookie(cookie); gotV != friendLogin || gotE != expectErr[iter] {
			t.Error("expected :", expectErr[iter], "  iter :", iter)
		}
	}
}
