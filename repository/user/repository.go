package user

import (
	"../../errors"
	"../../models"
	"database/sql"
	_ "github.com/lib/pq"
)

type UserRepositoryRealisation struct {
	userDB *sql.DB
}

func NewUserRepositoryRealisation(username, password, dbName string) (UserRepositoryRealisation, error) {
	connectString := "user=" + username + " password=" + password + " dbname=" + dbName + " sslmode=disable"

	db, err := sql.Open("postgres", connectString)

	if err != nil {
		return UserRepositoryRealisation{}, errors.FailConnect
	}

	defer db.Close()

	return UserRepositoryRealisation{userDB: db}, nil

}

func (Data UserRepositoryRealisation) GetUserDataById(id int) (models.User, error) {
	user := models.User{}

	row := Data.userDB.QueryRow("SELECT login, phone, mail, name, surname, birthdate FROM users WHERE u_id=$1", id)
	errScan := row.Scan(&user.Login, &user.Telephone, &user.Email, &user.Name, &user.Surname, &user.Date)

	if errScan != nil {
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
	_, err := Data.userDB.Exec("update users set login = $1, phone = $2, mail = $3, name = $4, surname = $5, birthdate = $6, password = $7 WHERE u_id=$8", currentUserData.Login, currentUserData.Telephone, currentUserData.Email, currentUserData.Name, currentUserData.Surname, currentUserData.Date, currentUserData.Password, id)
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}
