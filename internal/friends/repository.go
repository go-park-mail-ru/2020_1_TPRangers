package friends

import "main/internal/models"

type FriendRepository interface {
	GetIdByLogin(string) (int, error)
	AddFriend(int, int) error
	DeleteFriend(int, int) error
	GetAllFriendsByLogin(string) ([]models.FriendLandingInfo, error)
	GetFriendIdByLogin(string) (int, error)
	GetUserFriendsById(int, int) ([]models.FriendLandingInfo, error)
	GetUserFriendsByLogin(string, int) ([]models.FriendLandingInfo, error)
	GetUserLoginById(int) (string, error)
	CheckFriendship(int, int) (bool, error)
	SearchFriends(int, string) ([]models.Person, error)
}
