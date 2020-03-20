package auth

import (
	"../../errors"
	"database/sql"
	_ "github.com/lib/pq"
)

type AuthRepositoryRealisation struct {
	authDB *sql.DB
}

func NewAuthRepositoryRealisation(db *sql.DB) AuthRepositoryRealisation {
	return AuthRepositoryRealisation{authDB: db}

}

func (Data AuthRepositoryRealisation) GetIdByEmail(email string) (int, error) {
	row := Data.authDB.QueryRow("SELECT u_id FROM users WHERE mail=$1", email)
	var u_id int
	errScan := row.Scan(&u_id)
	if errScan != nil {
		return 0, errors.NotExist
	}
	return u_id, nil
}

func (Data AuthRepositoryRealisation) GetPassword(email string) (string, error) {
	row := Data.authDB.QueryRow("SELECT password FROM users WHERE mail=$1", email)
	var password string
	errScan := row.Scan(&password)
	if errScan != nil {
		return "", errors.NotExist
	}
	return password, nil
}
