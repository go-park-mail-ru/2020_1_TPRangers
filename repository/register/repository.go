package register

import (
	"../../errors"
	"../../models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type RegisterRepositoryRealisation struct {
	regDB *sql.DB
}

func NewRegisterRepositoryRealisation(db *sql.DB) RegisterRepositoryRealisation {
	return RegisterRepositoryRealisation{regDB: db}

}

func (Data RegisterRepositoryRealisation) AddNewUser(userData models.User) error {
	//result
	_, err := Data.regDB.Exec("insert into Users (phone, mail, name, surname, password, birthdate, login) values ($1, $2, $3, $4, $5, $6, $7)", userData.Telephone, userData.Email, userData.Name, userData.Surname, userData.Password, userData.Date, userData.Login)
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
	fmt.Println("ID", u_id, "\n err", errScan)
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
