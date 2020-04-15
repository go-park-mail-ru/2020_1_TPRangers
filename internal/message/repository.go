package message

import (
	"main/internal/models"
)

type MessageRepository interface {
	// author , message instance
	AddNewMessage(int, models.Message) error
	// get new messages , return array of new messages
	ReceiveNewMessages(int) ([]models.Message, error)
	// main user , chat with user , count of needed messages , from counter of message , return array of messages
	GetMessagesFromChat(int, int, int, int) ([]models.Message, error)
}
