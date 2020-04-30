package repository

import (
	"bytes"
	cr "crypto/rand"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"
	"main/internal/models"
	_errors "main/internal/tools/errors"
	"math/rand"
	"testing"
)

func TestUserRepositoryRealisation_AddNewUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{nil, errors.New("smth wrong")}
	expectBehaviour := []error{nil, _errors.FailSendToDB}
	userData := models.User{
		Login:           "123123",
		Telephone:       "123123",
		Email:           "123123",
		Name:            "123123",
		Password:        "123123",
		Surname:         "123123",
		Date:            "123123",
		Photo:           2,
		CryptedPassword: []byte("asdasd"),
	}

	for iter, _ := range expectBehaviour {
		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectExec(` insert into Users \(phone, mail, name, surname, password, birthdate, login, photo_id\) values \(\$1, \$2, \$3, \$4, \$5\:\:bytea, \$6, \$7, \$8\) `).WithArgs(userData.Telephone, userData.Email, userData.Name, userData.Surname, userData.CryptedPassword, userData.Date, userData.Login, userData.Photo).WillReturnError(errs[iter])
		} else {
			mock.ExpectExec(` insert into Users \(phone, mail, name, surname, password, birthdate, login, photo_id\) values \(\$1, \$2, \$3, \$4, \$5\:\:bytea, \$6, \$7, \$8\) `).WithArgs(userData.Telephone, userData.Email, userData.Name, userData.Surname, userData.CryptedPassword, userData.Date, userData.Login, userData.Photo).WillReturnResult(sqlmock.NewResult(1, 1))
		}
		mock.ExpectCommit()

		var err error
		tx, _ := db.Begin()

		if err = uTest.AddNewUser(userData); err != expectBehaviour[iter] {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}

		err = nil

		switch err {
		case nil:
			err = tx.Commit()
			if err != nil {
				fmt.Println(err)
			}
		default:
			tx.Rollback()
		}
	}

}

func TestUserRepositoryRealisation_IsUserExist(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{nil, _errors.FailReadFromDB}
	expectBehaviour := []error{nil, _errors.FailReadFromDB}
	existStatus := []bool{true, false}

	for iter, _ := range expectBehaviour {
		login := uuid.NewV4()
		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectQuery(` SELECT u_id FROM users WHERE mail\=\$1 `).WithArgs(login.String()).WillReturnError(errs[iter])
		} else {
			mock.ExpectQuery(` SELECT u_id FROM users WHERE mail\=\$1 `).WithArgs(login.String()).WillReturnRows(sqlmock.NewRows([]string{"u_id"}).AddRow(rand.Int()))
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if flag, err := uTest.IsUserExist(login.String()); !(err == expectBehaviour[iter] && flag == existStatus[iter]) {
			t.Error(err, expectBehaviour[iter], iter)
		}

		err = nil

		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}
}

func TestUserRepositoryRealisation_GetDefaultProfilePhotoId(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{nil, _errors.FailReadFromDB}
	expectBehaviour := []error{nil, _errors.FailReadFromDB}

	for iter, _ := range expectBehaviour {
		photoId := rand.Int()
		photoUrl := "https://social-hub.ru/uploads/img/default.png"
		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectQuery(`  SELECT photo_id FROM photos WHERE url\=\$1  `).WithArgs(photoUrl).WillReturnError(errs[iter])
			photoId = 0
		} else {
			mock.ExpectQuery(`  SELECT photo_id FROM photos WHERE url\=\$1  `).WithArgs(photoUrl).WillReturnRows(sqlmock.NewRows([]string{"photo_id"}).AddRow(photoId))
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if currentPhoto, err := uTest.GetDefaultProfilePhotoId(); !(err == expectBehaviour[iter] && currentPhoto == photoId) {
			t.Error(err, expectBehaviour[iter], iter)
		}

		err = nil

		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}
}

func TestUserRepositoryRealisation_GetPassword(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{_errors.FailReadFromDB, nil}
	expectBehaviour := []error{_errors.NotExist, nil}

	for iter, _ := range expectBehaviour {
		login := uuid.NewV4()
		password := make([]byte, 8)
		cr.Read(password)

		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectQuery(` SELECT password FROM users WHERE mail\=\$1 `).WithArgs(login.String()).WillReturnError(errs[iter])
			password = []byte{}
		} else {
			mock.ExpectQuery(` SELECT password FROM users WHERE mail\=\$1 `).WithArgs(login.String()).WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(password))
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if cPass, err := uTest.GetPassword(login.String()); !(bytes.Equal(cPass, password) && err == expectBehaviour[iter]) {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}

		err = nil

		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}

	}

}

func TestUserRepositoryRealisation_GetIdByEmail(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{_errors.FailReadFromDB, nil}
	expectBehaviour := []error{_errors.NotExist, nil}

	for iter, _ := range expectBehaviour {
		email := uuid.NewV4()
		uId := rand.Int()

		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectQuery(`  SELECT u_id FROM users WHERE mail\=\$1 `).WithArgs(email.String()).WillReturnError(errs[iter])
			uId = 0
		} else {
			mock.ExpectQuery(`  SELECT u_id FROM users WHERE mail\=\$1  `).WithArgs(email.String()).WillReturnRows(sqlmock.NewRows([]string{"u_id"}).AddRow(uId))
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if currentId, err := uTest.GetIdByEmail(email.String()); !(currentId == uId && err == expectBehaviour[iter]) {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}

		err = nil

		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}

	}

}

func TestUserRepositoryRealisation_GetUserProfileSettingsByLogin(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{_errors.NotExist, nil}
	expectBehaviour := []error{_errors.NotExist, nil}

	for iter, _ := range expectBehaviour {
		login := uuid.NewV4()
		user := models.Settings{
			Login:     "123123",
			Telephone: "123123",
			Email:     "123123",
			Name:      "123123",
			Surname:   "123123",
			Date:      "123123",
			Photo:     "123123",
		}

		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectQuery(`  SELECT U\.login, U\.phone, U\.mail, U\.name, U\.surname, U\.birthdate , P\.url FROM users U INNER JOIN photos P USING \(photo_id\) WHERE U\.login\=\$1 GROUP BY U\.login, U\.phone, U\.mail, U\.name, U\.surname, U\.birthdate , P\.url  `).WithArgs(login.String()).WillReturnError(errs[iter])
			user = models.Settings{}
		} else {
			mock.ExpectQuery(`   SELECT U\.login, U\.phone, U\.mail, U\.name, U\.surname, U\.birthdate , P\.url FROM users U INNER JOIN photos P USING \(photo_id\) WHERE U\.login\=\$1 GROUP BY U\.login, U\.phone, U\.mail, U\.name, U\.surname, U\.birthdate , P\.url `).WithArgs(login.String()).WillReturnRows(sqlmock.NewRows([]string{"U.login", "U.phone", "U.mail", "U.name", "U.surname", "U.birthdate", "P.url"}).AddRow(user.Login, user.Telephone, user.Email, user.Name, user.Surname, user.Date, user.Photo))
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if currentSettings, err := uTest.GetUserProfileSettingsByLogin(login.String()); !(currentSettings == user && err == expectBehaviour[iter]) {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}

		err = nil

		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}

	}

}

func TestUserRepositoryRealisation_GetUserProfileSettingsById(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{_errors.NotExist, nil}
	expectBehaviour := []error{_errors.NotExist, nil}

	for iter, _ := range expectBehaviour {
		uId := rand.Int()
		user := models.Settings{
			Login:     "123123",
			Telephone: "123123",
			Email:     "123123",
			Name:      "123123",
			Surname:   "123123",
			Date:      "123123",
			Photo:     "123123",
		}

		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectQuery(`  SELECT U\.login, U\.phone, U\.mail, U\.name, U\.surname, U\.birthdate , P\.url FROM users U INNER JOIN photos P USING \(photo_id\) WHERE U\.u_id\=\$1 GROUP BY U\.login, U\.phone, U\.mail, U\.name, U\.surname, U\.birthdate , P\.url  `).WithArgs(uId).WillReturnError(errs[iter])
			user = models.Settings{}
		} else {
			mock.ExpectQuery(`  SELECT U\.login, U\.phone, U\.mail, U\.name, U\.surname, U\.birthdate , P\.url FROM users U INNER JOIN photos P USING \(photo_id\) WHERE U\.u_id\=\$1 GROUP BY U\.login, U\.phone, U\.mail, U\.name, U\.surname, U\.birthdate , P\.url  `).WithArgs(uId).WillReturnRows(sqlmock.NewRows([]string{"U.login", "U.phone", "U.mail", "U.name", "U.surname", "U.birthdate", "P.url"}).AddRow(user.Login, user.Telephone, user.Email, user.Name, user.Surname, user.Date, user.Photo))
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if currentSettings, err := uTest.GetUserProfileSettingsById(uId); !(currentSettings == user && err == expectBehaviour[iter]) {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}

		err = nil

		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}

	}

}

func TestUserRepositoryRealisation_UploadProfilePhoto(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{_errors.NotExist, nil}
	expectBehaviour := []error{_errors.NotExist, nil}

	for iter, _ := range expectBehaviour {
		photoId := rand.Int()
		photoUrl := uuid.NewV4()

		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectQuery(`  INSERT INTO photos \(url, photos_likes_count\) VALUES \(\$1 , 0\) RETURNING photo_id `).WithArgs(photoUrl.String()).WillReturnError(errs[iter])
			photoId = 0
		} else {
			mock.ExpectQuery(`  INSERT INTO photos \(url, photos_likes_count\) VALUES \(\$1 , 0\) RETURNING photo_id  `).WithArgs(photoUrl.String()).WillReturnRows(sqlmock.NewRows([]string{"photo_id"}).AddRow(photoId))
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if currPhotoId, err := uTest.UploadProfilePhoto(photoUrl.String()); !(currPhotoId == photoId && err == expectBehaviour[iter]) {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}

		err = nil

		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}

	}

}

func TestUserRepositoryRealisation_UploadSettings(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{nil, errors.New("smth wrong")}
	expectBehaviour := []error{nil, _errors.FailSendToDB}
	userData := models.User{
		Login:           "123123",
		Telephone:       "123123",
		Email:           "123123",
		Name:            "123123",
		Password:        "123123",
		Surname:         "123123",
		Date:            "123123",
		Photo:           2,
		CryptedPassword: []byte("asdasd"),
	}

	for iter, _ := range expectBehaviour {
		uId := rand.Int()
		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectExec(`  update users set login \= \$1, phone \= \$2, mail \= \$3, name \= \$4, surname \= \$5, birthdate \= \$6, password \= \$7\:\:bytea , photo_id \= \$8 WHERE u_id\=\$9 `).WithArgs(userData.Login, userData.Telephone, userData.Email, userData.Name, userData.Surname, userData.Date, userData.CryptedPassword, userData.Photo, uId).WillReturnError(errs[iter])
		} else {
			mock.ExpectExec(`  update users set login \= \$1, phone \= \$2, mail \= \$3, name \= \$4, surname \= \$5, birthdate \= \$6, password \= \$7\:\:bytea , photo_id \= \$8 WHERE u_id\=\$9  `).WithArgs(userData.Login, userData.Telephone, userData.Email, userData.Name, userData.Surname, userData.Date, userData.CryptedPassword, userData.Photo, uId).WillReturnResult(sqlmock.NewResult(1, 1))
		}
		mock.ExpectCommit()

		var err error
		tx, _ := db.Begin()

		if err = uTest.UploadSettings(uId, userData); err != expectBehaviour[iter] {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}

		err = nil

		switch err {
		case nil:
			err = tx.Commit()
			if err != nil {
				fmt.Println(err)
			}
		default:
			tx.Rollback()
		}
	}

}

func TestUserRepositoryRealisation_GetUserDataByLogin(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{nil, errors.New("smth wrong")}
	expectBehaviour := []error{nil, _errors.FailReadToVar}
	user := models.User{
		Login:     "123123",
		Telephone: "123123",
		Email:     "123123",
		Name:      "123123",
		Surname:   "123123",
		Date:      "123123",
		Photo:     2,
	}

	for iter, _ := range expectBehaviour {
		login := uuid.NewV4()
		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectQuery(`  SELECT login, phone, mail, name, surname, birthdate , photo_id FROM users WHERE mail\=\$1 `).WithArgs(login).WillReturnError(errs[iter])
			user = models.User{}
		} else {
			mock.ExpectQuery(`   SELECT login, phone, mail, name, surname, birthdate , photo_id FROM users WHERE mail\=\$1 `).WithArgs(login).WillReturnRows(sqlmock.NewRows([]string{"login", "phone", "mail", "name", "surname", "birthdate", "photo_id "}).AddRow(user.Login, user.Telephone, user.Email, user.Name, user.Surname, user.Date, user.Photo))
		}
		mock.ExpectCommit()

		var err error
		tx, _ := db.Begin()

		if currentUser, err := uTest.GetUserDataByLogin(login.String()); !(err == expectBehaviour[iter] && currentUser.Name == user.Name) {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}

		err = nil

		switch err {
		case nil:
			err = tx.Commit()
			if err != nil {
				fmt.Println(err)
			}
		default:
			tx.Rollback()
		}
	}

}

func TestUserRepositoryRealisation_GetUserDataById(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{nil, errors.New("smth wrong")}
	expectBehaviour := []error{nil, _errors.FailReadToVar}
	pass := make([]byte, 8)
	cr.Read(pass)
	user := models.User{
		Login:           "123123",
		Telephone:       "123123",
		Email:           "123123",
		Name:            "123123",
		Surname:         "123123",
		Date:            "123123",
		Photo:           2,
		CryptedPassword: pass,
	}

	for iter, _ := range expectBehaviour {
		uId := rand.Int()
		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectQuery(`  SELECT login, phone, mail, name, surname, birthdate , photo_id , password FROM users WHERE u_id\=\$1  `).WithArgs(uId).WillReturnError(errs[iter])
			user = models.User{}
		} else {
			mock.ExpectQuery(`  SELECT login, phone, mail, name, surname, birthdate , photo_id , password FROM users WHERE u_id\=\$1 `).WithArgs(uId).WillReturnRows(sqlmock.NewRows([]string{"login", "phone", "mail", "name", "surname", "birthdate", "photo_id ", "password"}).AddRow(user.Login, user.Telephone, user.Email, user.Name, user.Surname, user.Date, user.Photo, user.CryptedPassword))
		}
		mock.ExpectCommit()

		var err error
		tx, _ := db.Begin()

		if currentUser, err := uTest.GetUserDataById(uId); !(err == expectBehaviour[iter] && currentUser.Name == user.Name) {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}

		err = nil

		switch err {
		case nil:
			err = tx.Commit()
			if err != nil {
				fmt.Println(err)
			}
		default:
			tx.Rollback()
		}
	}

}

func TestUserRepositoryRealisation_GetIdByLogin(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{nil, _errors.FailReadToVar}
	exceptBehaviour := []error{nil, _errors.FailReadToVar}

	for iter, _ := range exceptBehaviour {
		login := uuid.NewV4()
		uId := rand.Int()

		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectQuery(` select users\.u_id from users where users\.login \= \$1 `).WithArgs(login.String()).WillReturnError(errs[iter])
			uId = 0
		} else {
			mock.ExpectQuery(` select users\.u_id from users where users\.login \= \$1 `).WithArgs(login.String()).WillReturnRows(sqlmock.NewRows([]string{"u_id"}).AddRow(uId))
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if currentId, err := uTest.GetIdByLogin(login.String()); !(currentId == uId && err == exceptBehaviour[iter]) {
			t.Error(iter, err, exceptBehaviour[iter])
			return
		}

		err = nil
		switch err {
		case nil:
			tx.Commit()

		default:
			tx.Rollback()
		}
	}
}

func TestUserRepositoryRealisation_GetUserLoginById(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uTest := NewUserRepositoryRealisation(db)

	errs := []error{nil, _errors.FailReadToVar}
	exceptBehaviour := []error{nil, _errors.FailReadToVar}

	for iter, _ := range exceptBehaviour {
		uLog := uuid.NewV4()
		login := uLog.String()
		uId := rand.Int()

		mock.ExpectBegin()
		if errs[iter] != nil {
			mock.ExpectQuery(`  SELECT login FROM Users WHERE u_id \= \$1  `).WithArgs(uId).WillReturnError(errs[iter])
			login = ""
		} else {
			mock.ExpectQuery(`  SELECT login FROM Users WHERE u_id \= \$1  `).WithArgs(uId).WillReturnRows(sqlmock.NewRows([]string{"login"}).AddRow(login))
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if currentLogin, err := uTest.GetUserLoginById(uId); !(currentLogin == login && err == exceptBehaviour[iter]) {
			t.Error(iter, err, exceptBehaviour[iter])
			return
		}

		err = nil
		switch err {
		case nil:
			tx.Commit()

		default:
			tx.Rollback()
		}
	}
}
