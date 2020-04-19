package usecase

import (
	"main/internal/chats"
	"main/internal/models"
	"main/internal/friends"
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

func (CU ChatUseCaseRealisation) CreateChat(newChat models.NewChatUsers , userId int) error {

	usersChat := make([]int , 0)
	usersChat = append(usersChat , userId)

	for iter , _ := range newChat.UsersLogin {

		uId , err := CU.userRepo.GetIdByLogin(newChat.UsersLogin[iter])

		if err != nil {
			return err
		}

		usersChat = append(usersChat, uId)

	}


	return CU.chatRepo.CreateNewChat(newChat.ChatPhoto , newChat.ChatName , usersChat)
}

func (CU ChatUseCaseRealisation) ExitChat(chatId int64 , userId int) error {
	return CU.chatRepo.ExitChat(chatId, userId)
}

func (CU ChatUseCaseRealisation) GetChatMessages(chatId int64 , userId int) (models.Chat , []models.Message ,error) {
	return CU.chatRepo.GetChatMessages(chatId , userId)
}

func (CU ChatUseCaseRealisation) GetAllChats(userId int) ([]models.Chat , error) {
	return CU.chatRepo.GetAllChats(userId)
}

