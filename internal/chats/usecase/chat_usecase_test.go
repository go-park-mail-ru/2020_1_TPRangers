package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"main/internal/models"
	errs "main/internal/tools/errors"
	"main/mocks"
	"math/rand"
	"strconv"
	"testing"
)

func TestChatUseCaseRealisation_CreateChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	cRepoMock := mock.NewMockChatRepository(ctrl)
	fRepoMock := mock.NewMockFriendRepository(ctrl)

	cUseCase := NewChatUseCaseRealisation(cRepoMock, fRepoMock)
	customErr := errors.New("smth happend")
	chErr := errors.New("create chat error")
	expectedBehaviour := []error{nil, customErr, chErr}

	for iter  := range expectedBehaviour {

		newChat := new(models.NewChatUsers)
		uId := rand.Int()
		if expectedBehaviour[iter] != customErr {

			newChat.UsersLogin = append(newChat.UsersLogin, "123")
			fRepoMock.EXPECT().GetIdByLogin("123").Return(2, nil)
			newChat.UsersLogin = append(newChat.UsersLogin, "234")
			fRepoMock.EXPECT().GetIdByLogin("234").Return(3, nil)

			cRepoMock.EXPECT().CreateNewChat(newChat.ChatPhoto, newChat.ChatName, []int{uId, 2, 3}).Return(expectedBehaviour[iter])

		} else {
			newChat.UsersLogin = append(newChat.UsersLogin, "123")
			fRepoMock.EXPECT().GetIdByLogin("123").Return(2, customErr)
		}

		if err := cUseCase.CreateChat(*newChat, uId); err != expectedBehaviour[iter] {
			t.Error("ITER :", iter, "expected err :", expectedBehaviour[iter], "got :", err)
		}

	}

}

func TestChatUseCaseRealisation_ExitChat(t *testing.T) {
	ctrl := gomock.NewController(t)
	cRepoMock := mock.NewMockChatRepository(ctrl)
	fRepoMock := mock.NewMockFriendRepository(ctrl)

	cUseCase := NewChatUseCaseRealisation(cRepoMock, fRepoMock)
	chErr := errors.New("create chat error")

	expectedBehaviour := []error{nil, errs.FailSendToDB, errs.FailSendToDB, chErr}

	for iter := range expectedBehaviour {
		uId := rand.Int()
		chatId := ""
		if expectedBehaviour[iter] != errs.FailSendToDB {
			chatId = "c123"
			cRepoMock.EXPECT().ExitChat(int64(123), uId).Return(expectedBehaviour[iter])
		} else {
			if iter != 2 {
				chatId = "123"
			} else {
				chatId = "cdeasd"
			}
		}

		if err := cUseCase.ExitChat(chatId, uId); err != expectedBehaviour[iter] {
			t.Error("ITER :", iter, "expected err :", expectedBehaviour[iter], "got :", err)
		}
	}

}

func TestChatUseCaseRealisation_GetChatMessages(t *testing.T) {
	ctrl := gomock.NewController(t)
	cRepoMock := mock.NewMockChatRepository(ctrl)
	fRepoMock := mock.NewMockFriendRepository(ctrl)

	cUseCase := NewChatUseCaseRealisation(cRepoMock, fRepoMock)

	uId := 1
	chatId1 := "123"
	chAndMsg := new(models.ChatAndMessages)
	cRepoMock.EXPECT().GetPrivateChatMessages(int64(123), uId).Return(chAndMsg.ChatInfo, chAndMsg.ChatMessages, nil)

	if _, err := cUseCase.GetChatMessages(chatId1, uId); err != nil {
		t.Error("expected err :", nil, "got :", err)
	}

	chatId2 := "123c"
	_, errPI := strconv.ParseInt(chatId2, 10, 64)

	if _, err := cUseCase.GetChatMessages(chatId2, uId); err.Error() != errPI.Error() {
		t.Error("expected err :", errPI, "got :", err)
	}

	chatId1 = "c123"
	cRepoMock.EXPECT().GetGroupChatMessages(int64(123), uId).Return(chAndMsg.ChatInfo, chAndMsg.ChatMessages, nil)

	if _, err := cUseCase.GetChatMessages(chatId1, uId); err != nil {
		t.Error("expected err :", nil, "got :", err)
	}

	chatId1 = "c123123cc"
	_, errPI = strconv.ParseInt(chatId1[1:], 10, 64)

	if _, err := cUseCase.GetChatMessages(chatId1, uId); err.Error() != errPI.Error() {
		t.Error("expected err :", errPI, "got :", err)
	}

}

func TestChatUseCaseRealisation_GetAllChats(t *testing.T) {
	ctrl := gomock.NewController(t)
	cRepoMock := mock.NewMockChatRepository(ctrl)
	fRepoMock := mock.NewMockFriendRepository(ctrl)

	cUseCase := NewChatUseCaseRealisation(cRepoMock, fRepoMock)
	customErr := errors.New("123")
	chs := make([]models.Chat, 2)

	cRepoMock.EXPECT().GetAllChats(1).Return(chs, customErr)

	if _, err := cUseCase.GetAllChats(1); err != customErr {
		t.Error("expected :", customErr, "got :", err)
	}

}
