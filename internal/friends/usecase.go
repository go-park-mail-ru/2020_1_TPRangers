package friends

type FriendUseCase interface {
	AddFriend(int, string) error
	DeleteFriend(int, string) error
	GetAllFriends(string) (map[string]interface{}, error)
	GetUserLoginById(int) (string, error)
}
