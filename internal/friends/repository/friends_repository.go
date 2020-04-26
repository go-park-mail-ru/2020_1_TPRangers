package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"main/internal/models"
	"main/internal/tools/errors"
	"strings"
)

type FriendRepositoryRealisation struct {
	friendDB *sql.DB
}

func NewFriendRepositoryRealisation(db *sql.DB) FriendRepositoryRealisation {
	return FriendRepositoryRealisation{friendDB: db}

}

func (Data FriendRepositoryRealisation) SearchFriends(userID int, valueOfSearch string) ([]models.Person, error) {
	persons := make([]models.Person, 0)
	if strings.Contains(valueOfSearch, " ") {
		arrayOfvalue := strings.Split(valueOfSearch, " ")
		nameOrSurname := arrayOfvalue[0]
		SurnameOrName := arrayOfvalue[1]
		rows, err := Data.friendDB.Query("SELECT u.name, u.surname, u.login, ph.url FROM users AS u INNER JOIN photos AS ph ON (u.photo_id = ph.photo_id) INNER JOIN friends AS f ON (f.u_id = $1 AND f.f_id = u.u_id) WHERE u.u_id != $1 AND ((u.name LIKE $2 AND u.surname LIKE $3) OR (u.name LIKE $3 AND u.surname LIKE $2));", userID,nameOrSurname + "%", SurnameOrName + "%")
		if err != nil {
			return nil, errors.FailReadFromDB
		}
		person := models.Person{}
		for rows.Next() {
			err = rows.Scan(&person.Name, &person.Surname, &person.Login, &person.PhotoUrl)
			if err != nil {
				return nil, errors.FailReadToVar
			}
			persons = append(persons, person)
		}

	} else {
		rows, err := Data.friendDB.Query("SELECT u.name, u.surname, u.login, ph.url FROM users AS u INNER JOIN photos AS ph ON (u.photo_id = ph.photo_id) INNER JOIN friends AS f ON (f.u_id = $1 AND f.f_id = u.u_id) WHERE u.u_id != $1 AND ((u.name LIKE $2) OR (u.surname LIKE $2));",userID, valueOfSearch + "%")
		if err != nil {
			return nil, errors.FailReadFromDB
		}
		person := models.Person{}
		for rows.Next() {
			err = rows.Scan(&person.Name, &person.Surname, &person.Login, &person.PhotoUrl)
			if err != nil {
				return nil, errors.FailReadToVar
			}
			persons = append(persons, person)
		}
	}



	return persons, nil
}

func (Data FriendRepositoryRealisation) GetUserLoginById(userId int) (string, error) {
	row := Data.friendDB.QueryRow("SELECT login FROM Users WHERE u_id = $1", userId)
	login := ""

	err := row.Scan(&login)

	return login, err
}

func (Data FriendRepositoryRealisation) GetUserFriendsById(id, friendsCount int) ([]models.FriendLandingInfo, error) {
	userFriends := make([]models.FriendLandingInfo, 0, 6)

	row, err := Data.friendDB.Query("select name, url , login from friends F inner join users U on F.f_id=U.u_id INNER JOIN photos P ON U.photo_id=P.photo_id "+
		"WHERE F.u_id=$1 GROUP BY F.u_id,F.f_id,U.u_id,P.photo_id LIMIT $2", id, friendsCount)
	defer func() {
		if row != nil {
			row.Close()
		}
	}()

	if err != nil {
		return nil, errors.FailReadFromDB
	}
	for row.Next() {

		var friendInfo models.FriendLandingInfo

		err = row.Scan(&friendInfo.Name, &friendInfo.Photo, &friendInfo.Login)
		if err != nil {
			return nil, errors.FailReadToVar
		}

		userFriends = append(userFriends, friendInfo)

	}

	return userFriends, nil
}

func (Data FriendRepositoryRealisation) GetAllFriendsByLogin(login string) ([]models.FriendLandingInfo, error) {
	userFriends := make([]models.FriendLandingInfo, 0, 20)

	id, err := Data.GetIdByLogin(login)

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	row, err := Data.friendDB.Query("select name, url , login , surname from friends F inner join users U on F.f_id=U.u_id INNER JOIN photos P ON U.photo_id=P.photo_id "+
		"WHERE F.u_id=$1 GROUP BY F.u_id,F.f_id,U.u_id,P.photo_id", id)
	defer func() {
		if row != nil {
			row.Close()
		}
	}()
	if err != nil {
		return nil, errors.FailReadFromDB
	}

	for row.Next() {

		var friendInfo models.FriendLandingInfo

		err = row.Scan(&friendInfo.Name, &friendInfo.Photo, &friendInfo.Login, &friendInfo.Surname)

		if err != nil {
			return nil, errors.FailReadToVar
		}

		userFriends = append(userFriends, friendInfo)

	}

	return userFriends, nil
}

func (Data FriendRepositoryRealisation) GetUserFriendsByLogin(login string, friendsCount int) ([]models.FriendLandingInfo, error) {

	id, err := Data.GetIdByLogin(login)

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	return Data.GetUserFriendsById(id, friendsCount)
}

func (Data FriendRepositoryRealisation) AddFriend(firstFriend, secondFriend int) error {
	var err error

	if firstFriend != secondFriend {
		_, err = Data.friendDB.Exec("INSERT INTO Friends (u_id , f_id) VALUES ($1 , $2) , ($2 , $1)", firstFriend, secondFriend)
	} else {
		err = errors.FailAddFriend
	}
	return err
}

func (Data FriendRepositoryRealisation) DeleteFriend(firstFriend, secondFriend int) error {
	_, err := Data.friendDB.Exec("DELETE FROM Friends WHERE ((u_id = $1 AND f_id = $2) OR (u_id = $2 AND f_id = $1))", firstFriend, secondFriend)
	return err
}

func (Data FriendRepositoryRealisation) GetFriendIdByLogin(login string) (int, error) {
	var friend_id int

	row := Data.friendDB.QueryRow("SELECT u_id FROM users WHERE login = $1", login)

	scanErr := row.Scan(&friend_id)

	return friend_id, scanErr
}

func (Data FriendRepositoryRealisation) GetIdByLogin(login string) (int, error) {

	var i *int

	row := Data.friendDB.QueryRow("SELECT users.u_id FROM users WHERE users.login = $1", login)

	err := row.Scan(&i)
	if err != nil {
		return 0, err
	}

	return *i, err
}

func (Data FriendRepositoryRealisation) CheckFriendship(id1, id2 int) (bool, error) {
	row := Data.friendDB.QueryRow("SELECT f_id FROM friends WHERE u_id=$1 AND f_id=$2", id1, id2)

	f_id := -1

	errScan := row.Scan(&f_id)

	if errScan == sql.ErrNoRows {
		return false, nil
	}

	if errScan != nil {
		return false, errScan
	}

	return true, nil

}
