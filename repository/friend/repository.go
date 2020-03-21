package friend

import (
	"../../errors"
	"database/sql"
	_ "github.com/lib/pq"
)

type FriendRepositoryRealisation struct {
	friendDB *sql.DB
}

func NewFriendRepositoryRealisation(db *sql.DB) FriendRepositoryRealisation {
	return FriendRepositoryRealisation{friendDB: db}
}

func (Data FriendRepositoryRealisation) AddFriend(firstFriend, secondFriend int) error {
	_, err := Data.friendDB.Exec("INSERT INTO Friends (u_id , f_id) VALUES ($1 , $2) , ($2 , $1)", firstFriend, secondFriend)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}

func (Data FriendRepositoryRealisation) GetFriendIdByMail(mail string) (int, error) {
	var friend_id int

	row := Data.friendDB.QueryRow("SELECT u_id FROM users WHERE mail = $1", mail)

	scanErr := row.Scan(&friend_id)

	return friend_id, scanErr
}
