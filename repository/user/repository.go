package user

import (
	"../../errors"
	"../../models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type UserRepositoryRealisation struct {
	userDB *sql.DB
}

func NewUserRepositoryRealisation(db *sql.DB) UserRepositoryRealisation {
	return UserRepositoryRealisation{userDB: db}

}

// фотка !!!
func (Data UserRepositoryRealisation) GetUserDataById(id int) (models.User, error) {
	user := models.User{}

	row := Data.userDB.QueryRow("SELECT login, phone, mail, name, surname, birthdate FROM users WHERE u_id=$1", id)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date)

	if errScan != nil {
		fmt.Println("ERROR", errScan.Error())
		return models.User{}, errors.FailReadToVar
	}

	return user, nil
}

func (Data UserRepositoryRealisation) GetUserDataByLogin(email string) (models.User, error) {
	user := models.User{}

	row := Data.userDB.QueryRow("SELECT login, phone, mail, name, surname, birthdate FROM users WHERE mail=$1", email)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date)

	if errScan != nil {
		return models.User{}, errors.FailReadToVar
	}
	return user, nil
}

func (Data UserRepositoryRealisation) UploadSettings(id int, currentUserData models.User) error {
	_, err := Data.userDB.Exec("update users set login = $1, phone = $2, mail = $3, name = $4, surname = $5, birthdate = $6, password = $7 , photo_id = $8 WHERE u_id=$9", currentUserData.Login, currentUserData.Telephone, currentUserData.Email, currentUserData.Name, currentUserData.Surname, currentUserData.Date, currentUserData.Password , currentUserData.Photo, id)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}

func (Data UserRepositoryRealisation) UploadPhoto(photoUrl string) (int, error) {
		row := Data.userDB.QueryRow("INSERT INTO photos (url, photos_likes_count) VALUES ($1 , 0) RETURNING photo_id", photoUrl)
	var photo_id int

	errScan := row.Scan(&photo_id)

	return photo_id, errScan
}

func (Data UserRepositoryRealisation) GetUserProfileSettingsByLogin(login string) (models.Settings, error) {
	user := models.Settings{}

	row := Data.userDB.QueryRow("SELECT U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url FROM users U INNER JOIN photos P USING (photo_id) WHERE U.login=$1 GROUP BY U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url", login)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date , &user.Photo)

	return user, errScan
}

func (Data UserRepositoryRealisation) GetUserProfileSettingsById(id int) (models.Settings, error) {
	user := models.Settings{}

	row := Data.userDB.QueryRow("SELECT U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url FROM users U INNER JOIN photos P USING (photo_id) WHERE U.u_id=$1 GROUP BY U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url", id)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date , &user.Photo)

	return user, errScan
}

func (Data UserRepositoryRealisation) GetIdByEmail(email string) (int, error) {
	row := Data.userDB.QueryRow("SELECT u_id FROM users WHERE mail=$1", email)
	var u_id int
	errScan := row.Scan(&u_id)
	if errScan != nil {
		return 0, errors.NotExist
	}
	return u_id, nil
}

func (Data UserRepositoryRealisation) GetPassword(email string) (string, error) {
	row := Data.userDB.QueryRow("SELECT password FROM users WHERE mail=$1", email)
	var password string
	errScan := row.Scan(&password)
	if errScan != nil {
		return "", errors.NotExist
	}
	return password, nil
}

func (Data UserRepositoryRealisation) GetDefaultProfilePhotoId() (int , error) {
	row := Data.userDB.QueryRow("SELECT photo_id FROM photos WHERE url=$1", "defaults/profile/avatar")

	var photo_id int
	errScan := row.Scan(&photo_id)

	return photo_id , errScan
}

func (Data UserRepositoryRealisation) AddNewUser(userData models.User) error {
	//result
	_, err := Data.userDB.Exec("insert into Users (phone, mail, name, surname, password, birthdate, login, photo_id) values ($1, $2, $3, $4, $5, $6, $7, $8)", userData.Telephone, userData.Email, userData.Name, userData.Surname, userData.Password, userData.Date, userData.Login , userData.Photo)
	if err != nil {
		return errors.FailSendToDB
	}
	//fmt.Println(result)
	return nil
}

func (Data UserRepositoryRealisation) IsUserExist(email string) (bool, error) {
	row := Data.userDB.QueryRow("SELECT u_id FROM users WHERE mail=$1", email)
	var u_id int
	errScan := row.Scan(&u_id)
	if errScan != nil {
		return false, nil
	}
	return true, nil
}

func (Data UserRepositoryRealisation) AddFriend(firstFriend, secondFriend int) error {
	_, err := Data.userDB.Exec("INSERT INTO Friends (u_id , f_id) VALUES ($1 , $2) , ($2 , $1)", firstFriend, secondFriend)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}

func (Data UserRepositoryRealisation) GetFriendIdByLogin(login string) (int, error) {
	var friend_id int

	row := Data.userDB.QueryRow("SELECT u_id FROM users WHERE login = $1", login)

	scanErr := row.Scan(&friend_id)

	return friend_id, scanErr
}


