package repository

type FriendRepository interface {
	AddFriend(int, int) error
	GetFriendIdByMail(string) (int, error)
}
