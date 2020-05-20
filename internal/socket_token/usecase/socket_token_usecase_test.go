package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"main/mocks"
	"testing"
)

func TestTokenUseCaseRealisation_CreateNewToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	stockenMock := mock.NewMockTokenRepository(ctrl)
	tokeUTest := NewTokenUseCaseRealisation(stockenMock)

	customErr := errors.New("123")
	userId := 1
	stockenMock.EXPECT().AddNewToken(gomock.Any(),userId).Return(customErr)

	if _ , err := tokeUTest.CreateNewToken(userId); err != customErr {
		t.Error("unexpected")
	}

}
