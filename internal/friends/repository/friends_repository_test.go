package repository

import (
	"database/sql"
	errors "errors"
	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"
	"main/internal/models"
	_error "main/internal/tools/errors"
	"math/rand"
	"testing"
)

func TestFriendRepositoryRealisation_AddFriend(t *testing.T) {

	db, mock, _ := sqlmock.New()

	firstFriend := []int{1, 2, 3}
	secondFriend := []int{5, 2, 4}
	customErr := errors.New("smth happend")
	errs := []error{nil, nil, customErr}
	expectBehaviour := []error{nil, _error.FailAddFriend, customErr}

	fRepo := NewFriendRepositoryRealisation(db)

	for iter, _ := range expectBehaviour {
		mock.ExpectBegin()

		if firstFriend[iter] != secondFriend[iter] {
			if errs[iter] == nil {
				mock.ExpectExec(` INSERT INTO Friends \(u_id , f_id\) VALUES \(\$1 , \$2\) , \(\$2 , \$1\) `).WithArgs(firstFriend[iter], secondFriend[iter]).WillReturnResult(sqlmock.NewResult(2, 2))
			} else {
				mock.ExpectExec(` INSERT INTO Friends \(u_id , f_id\) VALUES \(\$1 , \$2\) , \(\$2 , \$1\) `).WithArgs(firstFriend[iter], secondFriend[iter]).WillReturnError(errs[iter])
			}
		}
		mock.ExpectCommit()

		tx, err := db.Begin()
		if err != nil {
			return
		}
		err = fRepo.AddFriend(firstFriend[iter], secondFriend[iter])

		if err != expectBehaviour[iter] {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}
		err = nil

		switch err {
		case nil:
			err = tx.Commit()
			if err != nil {
				return
			}
		default:
			err = tx.Rollback()
			if err != nil {
				return
			}
		}

	}

}

func TestFriendRepositoryRealisation_DeleteFriend(t *testing.T) {
	db, mock, _ := sqlmock.New()

	firstFriend := []int{1, 2, 3}
	secondFriend := []int{5, 2, 4}
	customErr := errors.New("smth happend")
	errs := []error{nil, nil, customErr}
	expectBehaviour := []error{nil, nil, customErr}

	fRepo := NewFriendRepositoryRealisation(db)

	for iter, _ := range expectBehaviour {
		mock.ExpectBegin()

		if errs[iter] == nil {
			mock.ExpectExec(`  DELETE FROM Friends WHERE \(\(u_id \= \$1 AND f_id \= \$2\) OR \(u_id \= \$2 AND f_id \= \$1\)  `).WithArgs(firstFriend[iter], secondFriend[iter]).WillReturnResult(sqlmock.NewResult(0, 2))
		} else {
			mock.ExpectExec(`  DELETE FROM Friends WHERE \(\(u_id \= \$1 AND f_id \= \$2\) OR \(u_id \= \$2 AND f_id \= \$1\)  `).WithArgs(firstFriend[iter], secondFriend[iter]).WillReturnError(errs[iter])
		}

		mock.ExpectCommit()

		tx, err := db.Begin()
		err = fRepo.DeleteFriend(firstFriend[iter], secondFriend[iter])

		if err != expectBehaviour[iter] {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}
		err = nil

		switch err {
		case nil:
			err = tx.Commit()
			if err != nil {
				return
			}
		default:
			err = tx.Rollback()
			if err != nil {
				return
			}
		}

	}
}

func TestFriendRepositoryRealisation_CheckFriendship(t *testing.T) {
	db, mock, _ := sqlmock.New()

	firstFriend := []int{1, 2, 3}
	secondFriend := []int{5, 2, 4}
	customErr := errors.New("smth happend")
	errs := []error{nil, sql.ErrNoRows, customErr}
	expectBehaviour := []error{nil, nil, customErr}
	friendshipStatus := []bool{true, false, false}

	fRepo := NewFriendRepositoryRealisation(db)

	for iter, _ := range expectBehaviour {
		mock.ExpectBegin()

		if errs[iter] == nil {
			mock.ExpectQuery(`  SELECT f_id FROM friends WHERE u_id\=\$1 AND f_id\=\$2  `).WithArgs(firstFriend[iter], secondFriend[iter]).WillReturnRows(sqlmock.NewRows([]string{"f_id"}).AddRow(secondFriend[iter]))
		} else {
			mock.ExpectQuery(`   SELECT f_id FROM friends WHERE u_id\=\$1 AND f_id\=\$2   `).WithArgs(firstFriend[iter], secondFriend[iter]).WillReturnError(errs[iter])
		}

		mock.ExpectCommit()

		tx, err := db.Begin()
		fStatus, err := fRepo.CheckFriendship(firstFriend[iter], secondFriend[iter])

		if err != expectBehaviour[iter] || fStatus != friendshipStatus[iter] {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}
		err = nil

		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}

	}
}

func TestFriendRepositoryRealisation_GetIdByLogin(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uniqueLogin := uuid.NewV4()

	logins := []string{uniqueLogin.String(), uniqueLogin.String(), uniqueLogin.String()}
	customErr := errors.New("smth happend")
	errs := []error{nil, sql.ErrNoRows, customErr}
	expectBehaviour := []error{nil, sql.ErrNoRows, customErr}
	ids := []int{3, 0, 0}
	expectId := []int{3, 0, 0}

	fRepo := NewFriendRepositoryRealisation(db)

	for iter, _ := range expectBehaviour {
		mock.ExpectBegin()

		if errs[iter] == nil {
			mock.ExpectQuery(`   SELECT users\.u_id FROM users WHERE users\.login \= \$1  `).WithArgs(logins[iter]).WillReturnRows(sqlmock.NewRows([]string{"u_id"}).AddRow(ids[iter]))
		} else {
			mock.ExpectQuery(`    SELECT users\.u_id FROM users WHERE users\.login \= \$1  `).WithArgs(logins[iter]).WillReturnError(errs[iter])
		}

		mock.ExpectCommit()

		tx, err := db.Begin()
		id, err := fRepo.GetIdByLogin(logins[iter])

		if err != expectBehaviour[iter] || id != expectId[iter] {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}
		err = nil

		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}

	}
}

func TestFriendRepositoryRealisation_GetFriendIdByLogin(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uniqueLogin := uuid.NewV4()

	logins := []string{uniqueLogin.String(), uniqueLogin.String(), uniqueLogin.String()}
	customErr := errors.New("smth happend")
	errs := []error{nil, sql.ErrNoRows, customErr}
	expectBehaviour := []error{nil, sql.ErrNoRows, customErr}
	ids := []int{3, 0, 0}
	expectId := []int{3, 0, 0}

	fRepo := NewFriendRepositoryRealisation(db)

	for iter, _ := range expectBehaviour {
		mock.ExpectBegin()

		if errs[iter] == nil {
			mock.ExpectQuery(`   SELECT u_id FROM users WHERE login \= \$1   `).WithArgs(logins[iter]).WillReturnRows(sqlmock.NewRows([]string{"u_id"}).AddRow(ids[iter]))
		} else {
			mock.ExpectQuery(`     SELECT u_id FROM users WHERE login \= \$1   `).WithArgs(logins[iter]).WillReturnError(errs[iter])
		}

		mock.ExpectCommit()

		tx, err := db.Begin()
		id, err := fRepo.GetFriendIdByLogin(logins[iter])

		if err != expectBehaviour[iter] || id != expectId[iter] {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}
		err = nil

		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}

	}
}

func TestFriendRepositoryRealisation_GetUserLoginById(t *testing.T) {
	db, mock, _ := sqlmock.New()
	uniqueLogin := uuid.NewV4()

	logins := []string{uniqueLogin.String(), "", ""}
	customErr := errors.New("smth happend")
	errs := []error{nil, sql.ErrNoRows, customErr}
	expectBehaviour := []error{nil, sql.ErrNoRows, customErr}
	ids := []int{rand.Int(), rand.Int(), rand.Int()}
	expectLogin := []string{uniqueLogin.String(), "", ""}

	fRepo := NewFriendRepositoryRealisation(db)

	for iter, _ := range expectBehaviour {
		mock.ExpectBegin()

		if errs[iter] == nil {
			mock.ExpectQuery(`    SELECT login FROM Users WHERE u_id \= \$1   `).WithArgs(ids[iter]).WillReturnRows(sqlmock.NewRows([]string{"login"}).AddRow(logins[iter]))
		} else {
			mock.ExpectQuery(`      SELECT login FROM Users WHERE u_id \= \$1  `).WithArgs(ids[iter]).WillReturnError(errs[iter])
		}

		mock.ExpectCommit()

		tx, err := db.Begin()
		login, err := fRepo.GetUserLoginById(ids[iter])

		if err != expectBehaviour[iter] || login != expectLogin[iter] {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}
		err = nil

		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}

	}
}

func TestFriendRepositoryRealisation_GetUserFriendsById(t *testing.T) {
	db, mock, _ := sqlmock.New()
	fRepo := NewFriendRepositoryRealisation(db)
	rowsNumber := 6

	errs := []error{nil, _error.FailReadFromDB, _error.FailReadToVar}
	expectBehaviour := []error{nil, _error.FailReadFromDB, _error.FailReadToVar}
	//expectLength := []int{rowsNumber, 0 , 0}

	for iter, _ := range errs {
		id := rand.Int()
		mock.ExpectBegin()
		if errs[iter] == _error.FailReadFromDB {
			mock.ExpectQuery(`  select name, url , login from friends F inner join users U on F\.f_id\=U\.u_id INNER JOIN photos P ON U\.photo_id\=P\.photo_id WHERE F\.u_id\=\$1 GROUP BY F\.u_id,F\.f_id,U\.u_id,P\.photo_id LIMIT \$2   `).WithArgs(id, rowsNumber).WillReturnError(errs[iter])
		} else {
			if errs[iter] != nil {
				mock.ExpectQuery(`  select name, url , login from friends F inner join users U on F\.f_id\=U\.u_id INNER JOIN photos P ON U\.photo_id\=P\.photo_id WHERE F\.u_id\=\$1 GROUP BY F\.u_id,F\.f_id,U\.u_id,P\.photo_id LIMIT \$2   `).WithArgs(id, rowsNumber).WillReturnRows(sqlmock.NewRows([]string{"name", "url", "login"}).AddRow(nil, "2", "3").AddRow(1, "2", "3").RowError(1, errs[iter]))
			} else {

				row := sqlmock.NewRows([]string{"name", "url", "login"})
				for i := 0; i < rowsNumber; i++ {
					uniqueLogin := uuid.NewV4()
					uniqueUrl := uuid.NewV4()
					uniqueName := uuid.NewV4()

					row.AddRow(uniqueName, uniqueUrl, uniqueLogin)
				}

				mock.ExpectQuery(`  select name, url , login from friends F inner join users U on F\.f_id\=U\.u_id INNER JOIN photos P ON U\.photo_id\=P\.photo_id WHERE F\.u_id\=\$1 GROUP BY F\.u_id,F\.f_id,U\.u_id,P\.photo_id LIMIT \$2 `).WithArgs(id, rowsNumber).WillReturnRows(row)
			}
			mock.ExpectCommit()
		}

		tx, err := db.Begin()
		_, err = fRepo.GetUserFriendsById(id, rowsNumber)

		if err != expectBehaviour[iter] {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}
		err = nil

		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}

}

func TestFriendRepositoryRealisation_GetUserFriendsByLogin(t *testing.T) {
	db, mock, _ := sqlmock.New()
	fRepo := NewFriendRepositoryRealisation(db)
	rowsNumber := 6

	errs := []error{nil, _error.FailReadFromDB, _error.FailReadToVar, nil}
	expectBehaviour := []error{nil, _error.FailReadFromDB, _error.FailReadToVar, _error.FailReadFromDB}
	userErr := []error{nil, nil, nil, errors.New("some err")}

	for iter, _ := range errs {
		login := uuid.NewV4()
		id := rand.Int()
		mock.ExpectBegin()

		if userErr[iter] == nil {
			mock.ExpectQuery(`   SELECT users\.u_id FROM users WHERE users\.login \= \$1    `).WithArgs(login.String()).WillReturnRows(sqlmock.NewRows([]string{"u_id"}).AddRow(id))

			if errs[iter] == _error.FailReadFromDB {
				mock.ExpectQuery(`  select name, url , login from friends F inner join users U on F\.f_id\=U\.u_id INNER JOIN photos P ON U\.photo_id\=P\.photo_id WHERE F\.u_id\=\$1 GROUP BY F\.u_id,F\.f_id,U\.u_id,P\.photo_id LIMIT \$2   `).WithArgs(id, rowsNumber).WillReturnError(errs[iter])
			} else {
				if errs[iter] != nil {
					mock.ExpectQuery(`  select name, url , login from friends F inner join users U on F\.f_id\=U\.u_id INNER JOIN photos P ON U\.photo_id\=P\.photo_id WHERE F\.u_id\=\$1 GROUP BY F\.u_id,F\.f_id,U\.u_id,P\.photo_id LIMIT \$2   `).WithArgs(id, rowsNumber).WillReturnRows(sqlmock.NewRows([]string{"name", "url", "login"}).AddRow(nil, "2", "3").AddRow(1, "2", "3").RowError(1, errs[iter]))
				} else {
					userFriends := make([]models.FriendLandingInfo, 6)
					row := sqlmock.NewRows([]string{"name", "url", "login"})
					for i := 0; i < rowsNumber; i++ {
						uniqueLogin := uuid.NewV4()
						uniqueUrl := uuid.NewV4()
						uniqueName := uuid.NewV4()
						userFriends[i].Login = uniqueLogin.String()
						userFriends[i].Photo = uniqueUrl.String()
						userFriends[i].Name = uniqueName.String()

						row.AddRow(uniqueName, uniqueUrl, uniqueLogin)
					}

					mock.ExpectQuery(` select name, url , login from friends F inner join users U on F\.f_id\=U\.u_id INNER JOIN photos P ON U\.photo_id\=P\.photo_id WHERE F\.u_id\=\$1 GROUP BY F\.u_id,F\.f_id,U\.u_id,P\.photo_id LIMIT \$2 `).WithArgs(id, rowsNumber).WillReturnRows(row)
				}
				mock.ExpectCommit()
			}
		} else {
			mock.ExpectQuery(` SELECT users\.u_id FROM users WHERE users\.login \= \$1 `).WithArgs(login.String()).WillReturnError(userErr[iter])
		}

		tx, err := db.Begin()
		_, err = fRepo.GetUserFriendsByLogin(login.String(), rowsNumber)

		if err != expectBehaviour[iter] {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}
		err = nil

		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}

}

func TestFriendRepositoryRealisation_GetAllFriendsByLogin(t *testing.T) {
	db, mock, _ := sqlmock.New()
	fRepo := NewFriendRepositoryRealisation(db)
	rowsNumber := 6

	errs := []error{nil, _error.FailReadFromDB, _error.FailReadToVar, nil}
	expectBehaviour := []error{nil, _error.FailReadFromDB, _error.FailReadToVar, _error.FailReadFromDB}
	userErr := []error{nil, nil, nil, errors.New("some err")}

	for iter, _ := range errs {
		login := uuid.NewV4()
		id := rand.Int()
		mock.ExpectBegin()
		if userErr[iter] == nil {
			mock.ExpectQuery(`   SELECT users\.u_id FROM users WHERE users\.login \= \$1    `).WithArgs(login.String()).WillReturnRows(sqlmock.NewRows([]string{"u_id"}).AddRow(id))

			if errs[iter] == _error.FailReadFromDB {
				mock.ExpectQuery(`  select name, url , login , surname from friends F inner join users U on F\.f_id\=U\.u_id INNER JOIN photos P ON U\.photo_id\=P\.photo_id WHERE F\.u_id\=\$1 GROUP BY F\.u_id,F\.f_id,U\.u_id,P\.photo_id  `).WithArgs(id).WillReturnError(errs[iter])
			} else {
				if errs[iter] != nil {
					mock.ExpectQuery(`  select name, url , login , surname from friends F inner join users U on F\.f_id\=U\.u_id INNER JOIN photos P ON U\.photo_id\=P\.photo_id WHERE F\.u_id\=\$1 GROUP BY F\.u_id,F\.f_id,U\.u_id,P\.photo_id  `).WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"name", "url", "login", "surname"}).AddRow(nil, "2", "3", "4").AddRow(1, "2", "3", "4").RowError(1, errs[iter]))
				} else {
					row := sqlmock.NewRows([]string{"name", "url", "login", "surname"})
					for i := 0; i < rowsNumber; i++ {
						uniqueLogin := uuid.NewV4()
						uniqueUrl := uuid.NewV4()
						uniqueName := uuid.NewV4()
						uniqueSurname := uuid.NewV4()

						row.AddRow(uniqueName.String(), uniqueUrl.String(), uniqueLogin.String(), uniqueSurname.String())
					}

					mock.ExpectQuery(`  select name, url , login , surname from friends F inner join users U on F\.f_id\=U\.u_id INNER JOIN photos P ON U\.photo_id\=P\.photo_id WHERE F\.u_id\=\$1 GROUP BY F\.u_id,F\.f_id,U\.u_id,P\.photo_id  `).WithArgs(id).WillReturnRows(row)
				}
				mock.ExpectCommit()
			}
		} else {
			mock.ExpectQuery(` SELECT users\.u_id FROM users WHERE users\.login \= \$1 `).WithArgs(login.String()).WillReturnError(userErr[iter])
		}

		tx, err := db.Begin()
		_, err = fRepo.GetAllFriendsByLogin(login.String())

		if err != expectBehaviour[iter] {
			t.Error(iter, err, expectBehaviour[iter])
			return
		}
		err = nil

		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}

}
