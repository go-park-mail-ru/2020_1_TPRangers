package chats

import "main/internal/models"

type ChatUseCase interface {
	CreateChat(models.NewChatUsers , int) error
	ExitChat(int64 , int) error
	GetChatMessages(int64 , int) (models.Chat , []models.Message , error)
	GetAllChats(int) ([]models.Chat , error)
}
