package usecase

import (
	"fmt"
	"main/internal/chats"
	"main/internal/friends"
	"main/internal/tools/errors"
	"main/models"
	"strconv"
)

type ChatUseCaseRealisation struct {
	chatRepo chats.ChatRepository
	userRepo friends.FriendRepository
}

func NewChatUseCaseRealisation(chatR chats.ChatRepository, userR friends.FriendRepository) ChatUseCaseRealisation {
	return ChatUseCaseRealisation{
		chatRepo: chatR,
		userRepo: userR,
	}
}

func (CU ChatUseCaseRealisation) CreateChat(newChat models.NewChatUsers, userId int) error {

	usersChat := make([]int, 0)
	usersChat = append(usersChat, userId)

	for iter, _ := range newChat.UsersLogin {

		uId, err := CU.userRepo.GetIdByLogin(newChat.UsersLogin[iter])

		if err != nil {
			return err
		}

		usersChat = append(usersChat, uId)

	}

	return CU.chatRepo.CreateNewChat(newChat.ChatPhoto, newChat.ChatName, usersChat)
}

func (CU ChatUseCaseRealisation) ExitChat(chatId string, userId int) error {

	if chatId[:1] != "c" {
		return errors.FailSendToDB
	}
	chat, err := strconv.ParseInt(chatId[1:], 10, 64)

	if err != nil {
		return errors.FailSendToDB
	}
	return CU.chatRepo.ExitChat(chat, userId)
}

func (CU ChatUseCaseRealisation) GetChatMessages(chatId string, userId int) (models.ChatAndMessages, error) {

	fmt.Println(chatId, chatId[:1])
	if chatId[:1] != "c" {
		chat, err := strconv.ParseInt(chatId, 10, 64)
		chAndMsg := new(models.ChatAndMessages)

		if err != nil {
			return *chAndMsg, err
		}
		fmt.Println("private chat")
		chAndMsg.ChatInfo, chAndMsg.ChatMessages, err = CU.chatRepo.GetPrivateChatMessages(chat, userId)

		return *chAndMsg, err

	}

	chat, err := strconv.ParseInt(chatId[1:], 10, 64)
	chAndMsg := new(models.ChatAndMessages)

	if err != nil {
		return *chAndMsg, err
	}
	fmt.Println("group chat", chat)
	chAndMsg.ChatInfo, chAndMsg.ChatMessages, err = CU.chatRepo.GetGroupChatMessages(chat, userId)

	return *chAndMsg, err

}

func (CU ChatUseCaseRealisation) GetAllChats(userId int) ([]models.Chat, error) {
	return CU.chatRepo.GetAllChats(userId)
}
