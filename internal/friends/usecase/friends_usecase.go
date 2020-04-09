package usecase

import (
	Sess "main/internal/cookies"
	Friend "main/internal/friends"
	"main/internal/tools/errors"
)


type FriendUseCaseRealisation struct {
	friendDB  Friend.FriendRepository
	sessionDB Sess.CookieRepository
}


func (userR FriendUseCaseRealisation) GetAllFriends(login string) (map[string]interface{}, error) {

	sendData := make(map[string]interface{})
	var err error
	sendData["friends"], err = userR.friendDB.GetAllFriendsByLogin(login)

	return sendData, err

}

func (userR FriendUseCaseRealisation) AddFriend(cookie, friendLogin string) error {

	id, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return errors.InvalidCookie
	}

	friendId, _ := userR.friendDB.GetFriendIdByLogin(friendLogin)

	err = userR.friendDB.AddFriend(id, friendId)

	if err != nil {
		return errors.FailAddFriend
	}

	return err
}

func (userR FriendUseCaseRealisation) DeleteFriend(cookie, friendLogin string) error {

	id, err := userR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return errors.InvalidCookie
	}

	friendId, _ := userR.friendDB.GetFriendIdByLogin(friendLogin)

	err = userR.friendDB.DeleteFriend(id, friendId)

	if err != nil {
		return errors.FailDeleteFriend
	}

	return err
}

func (userR FriendUseCaseRealisation) GetUserLoginByCookie(cookieValue string) (string, error) {
	id, _ := userR.sessionDB.GetUserIdByCookie(cookieValue)

	return userR.friendDB.GetUserLoginById(id)
}

func NewUserUseCaseRealisation(userDB Friend.FriendRepository, sesDB Sess.CookieRepository) FriendUseCaseRealisation {
	return FriendUseCaseRealisation{
		friendDB:    userDB,
		sessionDB: sesDB,
	}
}
