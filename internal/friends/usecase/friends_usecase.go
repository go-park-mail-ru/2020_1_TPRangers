package usecase

import (
	Friend "main/internal/friends"
	"main/internal/tools/errors"
)

type FriendUseCaseRealisation struct {
	friendDB Friend.FriendRepository
}

func (userR FriendUseCaseRealisation) GetAllFriends(login string) (map[string]interface{}, error) {

	sendData := make(map[string]interface{})
	var err error
	sendData["friends"], err = userR.friendDB.GetAllFriendsByLogin(login)

	return sendData, err

}

func (userR FriendUseCaseRealisation) AddFriend(userId int, friendLogin string) error {

	friendId, _ := userR.friendDB.GetFriendIdByLogin(friendLogin)

	err := userR.friendDB.AddFriend(userId, friendId)

	if err != nil {
		return errors.FailAddFriend
	}

	return err
}

func (userR FriendUseCaseRealisation) DeleteFriend(userId int, friendLogin string) error {

	friendId, _ := userR.friendDB.GetFriendIdByLogin(friendLogin)

	err := userR.friendDB.DeleteFriend(userId, friendId)

	if err != nil {
		return errors.FailDeleteFriend
	}

	return nil
}

func (userR FriendUseCaseRealisation) GetUserLoginById(userId int) (string, error) {
	return userR.friendDB.GetUserLoginById(userId)
}

func NewFriendUseCaseRealisation(userDB Friend.FriendRepository) FriendUseCaseRealisation {
	return FriendUseCaseRealisation{
		friendDB: userDB,
	}
}
