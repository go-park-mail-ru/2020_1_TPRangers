package register

import (
	"../../errors"
	"../../models"
	"database/sql"
	_ "github.com/lib/pq"
)

type RegisterRepositoryRealisation struct {
	regDB *sql.DB
}

func NewRegisterRepositoryRealisation(db *sql.DB) RegisterRepositoryRealisation {
	return RegisterRepositoryRealisation{regDB: db}
}

func (Data RegisterRepositoryRealisation) GetDefaultProfilePhotoId() (int , error) {
	row := Data.regDB.QueryRow("SELECT photo_id FROM photos WHERE url=$1", "defaults/profile/avatar")

	var photo_id int
	errScan := row.Scan(&photo_id)

	return photo_id , errScan
}

func (Data RegisterRepositoryRealisation) AddNewUser(userData models.User) error {
	//result
	_, err := Data.regDB.Exec("insert into Users (phone, mail, name, surname, password, birthdate, login, photo_id) values ($1, $2, $3, $4, $5, $6, $7, $8)", userData.Telephone, userData.Email, userData.Name, userData.Surname, userData.Password, userData.Date, userData.Login , userData.Photo)
	if err != nil {
		return errors.FailSendToDB
	}
	//fmt.Println(result)
	return nil
}

func (Data RegisterRepositoryRealisation) IsUserExist(email string) (bool, error) {
	row := Data.regDB.QueryRow("SELECT u_id FROM users WHERE mail=$1", email)
	var u_id int
	errScan := row.Scan(&u_id)
	if errScan != nil {
		return false, nil
	}
	return true, nil
}

func (Data RegisterRepositoryRealisation) GetIdByEmail(email string) (int, error) {
	row := Data.regDB.QueryRow("SELECT u_id FROM users WHERE mail=$1", email)
	var u_id int
	errScan := row.Scan(&u_id)
	if errScan != nil {
		return 0, errors.NotExist
	}
	return u_id, nil
}
