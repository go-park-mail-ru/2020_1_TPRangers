package chats

import "main/internal/models"

type ChatUseCase interface {
	CreateChat(models.NewChatUsers, int) error
	ExitChat(string, int) error
	GetChatMessages(string, int) (models.ChatAndMessages, error)
	GetAllChats(int) ([]models.Chat, error)
}
