package repository

import (
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestChatRepositoryRealisation_CreateNewChat(t *testing.T) {
	db, mock, _ := sqlmock.New()
	cRepo := NewChatRepositoryRealisation(db)
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO PrivateChats \(fu_id,su_id\) VALUES \(\$1,\$2\) RETURNING ch_id`).WithArgs(0, 0).WillReturnRows(sqlmock.NewRows([]string{"ch_id"}).AddRow(nil))
	mock.ExpectCommit()

	tx, _ := db.Begin()
	if err := cRepo.CreateNewChat("", "", []int{0}); err == nil {
		fmt.Println(err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func TestChatRepositoryRealisation_ExitChat(t *testing.T) {
	customErr := errors.New("smth happend")
	db, mock, _ := sqlmock.New()
	cRepo := NewChatRepositoryRealisation(db)

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM ChatsUsers WHERE u_id \= \$1 AND gch_id \= \$2`).WithArgs(0, int64(0)).WillReturnError(customErr)
	mock.ExpectCommit()

	tx, _ := db.Begin()
	if err := cRepo.ExitChat(int64(0), 0); err != customErr {
		fmt.Println(err)
		tx.Rollback()
	} else {
		tx.Commit()
	}

}
