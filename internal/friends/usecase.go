package friends

import "main/internal/models"

type FriendUseCase interface {
	AddFriend(int, string) error
	DeleteFriend(int, string) error
	GetAllFriends(string) ([]models.FriendLandingInfo, error)
	GetUserLoginById(int) (string, error)
	SearchFriends(int, string) ([]models.Person, error)
}
