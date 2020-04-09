package friends

type FriendUseCase interface {
	AddFriend(string, string) error
	DeleteFriend(string, string) error
	GetAllFriends(string) (map[string]interface{}, error)
	GetUserLoginByCookie(string) (string, error)
}
