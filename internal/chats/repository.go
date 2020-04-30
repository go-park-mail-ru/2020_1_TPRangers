package chats

import "main/models"

type ChatRepository interface {
	CreateNewChat(string, string, []int) error
	ExitChat(int64, int) error
	GetGroupChatMessages(int64, int) (models.ChatInfo, []models.Message, error)
	GetPrivateChatMessages(int64, int) (models.ChatInfo, []models.Message, error)
	GetAllChats(int) ([]models.Chat, error)
}
