package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"main/internal/models"
	"main/internal/tools/errors"
	"strings"
)

type UserRepositoryRealisation struct {
	userDB *sql.DB
}

func NewUserRepositoryRealisation(db *sql.DB) UserRepositoryRealisation {
	return UserRepositoryRealisation{userDB: db}

}

func (Data UserRepositoryRealisation) SearchUsers(userID int, valueOfSearch string) ([]models.Person, error) {
	persons := make([]models.Person, 0)
	if strings.Contains(valueOfSearch, " ") {
		arrayOfvalue := strings.Split(valueOfSearch, " ")
		nameOrSurname := arrayOfvalue[0]
		SurnameOrName := arrayOfvalue[1]

		rows, err := Data.userDB.Query("SELECT u.name, u.surname, u.login, ph.url FROM users AS u INNER JOIN photos AS ph ON (u.photo_id = ph.photo_id)  WHERE u.u_id != $3 AND ((lower(u.name) LIKE LOWER($1) AND lower(u.surname) LIKE LOWER($2)) OR (lower(u.name) LIKE LOWER($2) AND lower(u.surname) LIKE LOWER($1)));", nameOrSurname+"%", SurnameOrName+"%", userID)

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

		rows, err := Data.userDB.Query("SELECT u.name, u.surname, u.login, ph.url FROM users AS u INNER JOIN photos AS ph ON (u.photo_id = ph.photo_id)  WHERE u.u_id != $2 AND ((lower(u.name) LIKE LOWER($1)) OR (lower(u.surname) LIKE LOWER($1)));", valueOfSearch+"%", userID)

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

func (Data UserRepositoryRealisation) GetUserLoginById(userId int) (string, error) {
	row := Data.userDB.QueryRow("SELECT login FROM Users WHERE u_id = $1", userId)
	login := ""

	err := row.Scan(&login)

	return login, err
}

func (Data UserRepositoryRealisation) GetIdByLogin(login string) (int, error) {

	var id *int

	row := Data.userDB.QueryRow("select users.u_id from users where users.login = $1", login)

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return *id, err
}

func (Data UserRepositoryRealisation) GetUserDataById(id int) (models.User, error) {
	user := models.User{}

	row := Data.userDB.QueryRow("SELECT login, phone, mail, name, surname, birthdate , photo_id , password FROM users WHERE u_id=$1", id)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date, &user.Photo, &user.CryptedPassword)

	if errScan != nil {
		return models.User{}, errors.FailReadToVar
	}

	return user, nil
}

func (Data UserRepositoryRealisation) GetUserDataByLogin(email string) (models.User, error) {
	user := models.User{}

	row := Data.userDB.QueryRow("SELECT login, phone, mail, name, surname, birthdate , photo_id FROM users WHERE mail=$1", email)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date, &user.Photo)

	if errScan != nil {
		return models.User{}, errors.FailReadToVar
	}
	return user, nil
}

func (Data UserRepositoryRealisation) UploadSettings(id int, currentUserData models.User) error {
	_, err := Data.userDB.Exec("update users set login = $1, phone = $2, mail = $3, name = $4, surname = $5, birthdate = $6, password = $7::bytea , photo_id = $8 WHERE u_id=$9", currentUserData.Login, currentUserData.Telephone, currentUserData.Email, currentUserData.Name, currentUserData.Surname, currentUserData.Date, currentUserData.CryptedPassword, currentUserData.Photo, id)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}

func (Data UserRepositoryRealisation) UploadProfilePhoto(photoUrl string) (int, error) {

	row := Data.userDB.QueryRow("INSERT INTO photos (url, photos_likes_count) VALUES ($1 , 0) RETURNING photo_id", photoUrl)
	var photo_id int

	errScan := row.Scan(&photo_id)
	return photo_id, errScan
}

func (Data UserRepositoryRealisation) GetUserProfileSettingsByLogin(login string) (models.Settings, error) {
	user := models.Settings{}

	row := Data.userDB.QueryRow("SELECT U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url FROM users U INNER JOIN photos P USING (photo_id) WHERE U.login=$1 GROUP BY U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url", login)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date, &user.Photo)
	return user, errScan
}

func (Data UserRepositoryRealisation) GetUserProfileSettingsById(id int) (models.Settings, error) {
	user := models.Settings{}

	row := Data.userDB.QueryRow("SELECT U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url FROM users U INNER JOIN photos P USING (photo_id) WHERE U.u_id=$1 GROUP BY U.login, U.phone, U.mail, U.name, U.surname, U.birthdate , P.url", id)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date, &user.Photo)

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

func (Data UserRepositoryRealisation) GetPassword(email string) ([]byte, error) {
	row := Data.userDB.QueryRow("SELECT password FROM users WHERE mail=$1", email)
	var password []byte
	errScan := row.Scan(&password)
	if errScan != nil {
		return password, errors.NotExist
	}
	return password, nil
}

func (Data UserRepositoryRealisation) GetDefaultProfilePhotoId() (int, error) {
	row := Data.userDB.QueryRow("SELECT photo_id FROM photos WHERE url=$1", "https://social-hub.ru/uploads/img/default.png")

	var photo_id int
	errScan := row.Scan(&photo_id)

	return photo_id, errScan
}

func (Data UserRepositoryRealisation) AddNewUser(userData models.User) error {
	_, err := Data.userDB.Exec("insert into Users (phone, mail, name, surname, password, birthdate, login, photo_id) values ($1, $2, $3, $4, $5::bytea, $6, $7, $8)", userData.Telephone, userData.Email, userData.Name, userData.Surname, userData.CryptedPassword, userData.Date, userData.Login, userData.Photo)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}

func (Data UserRepositoryRealisation) IsUserExist(email string) (bool, error) {
	row := Data.userDB.QueryRow("SELECT u_id FROM users WHERE mail=$1", email)
	var u_id int
	errScan := row.Scan(&u_id)
	if errScan != nil {
		return false, errScan
	}
	return true, nil
}
