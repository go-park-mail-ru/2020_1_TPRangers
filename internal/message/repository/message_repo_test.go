package repository

//import(
//	"database/sql"
//	"fmt"
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/golang/mock/gomock"
//	"main/internal/models"
//	"main/internal/tools/errors"
//	"math/rand"
//	"testing"
//)
//
//func TestMessageRepositoryRealisation_AddNewMessage(t *testing.T) {
//	db, mock, _ := sqlmock.New()
//	MessageTest := NewMessageRepositoryRealisation("" , "" , db)
//	msgs := []models.Message{
//		models.Message{
//			ChatId:        "1",
//			ChatPhoto:     "",
//			ChatName:      "",
//			AuthorName:    "",
//			AuthorSurname: "",
//			AuthorUrl:     "",
//			AuthorPhoto:   "",
//			Text:          "1234",
//			Time:          "",
//			Sticker:       "",
//			IsMe:          false,
//		},
//		models.Message{
//			ChatId:        "c1",
//			ChatPhoto:     "",
//			ChatName:      "",
//			AuthorName:    "",
//			AuthorSurname: "",
//			AuthorUrl:     "",
//			AuthorPhoto:   "",
//			Text:          "1234",
//			Time:          "",
//			Sticker:       "",
//			IsMe:          false,
//		},
//		models.Message{
//			ChatId:        "123c",
//			ChatPhoto:     "",
//			ChatName:      "",
//			AuthorName:    "",
//			AuthorSurname: "",
//			AuthorUrl:     "",
//			AuthorPhoto:   "",
//			Text:          "1234",
//			Time:          "",
//			Sticker:       "",
//			IsMe:          false,
//		},
//		models.Message{
//			ChatId:        "c12c",
//			ChatPhoto:     "",
//			ChatName:      "",
//			AuthorName:    "",
//			AuthorSurname: "",
//			AuthorUrl:     "",
//			AuthorPhoto:   "",
//			Text:          "1234",
//			Time:          "",
//			Sticker:       "",
//			IsMe:          false,
//		},
//		models.Message{
//			ChatId:        "c1",
//			ChatPhoto:     "",
//			ChatName:      "",
//			AuthorName:    "",
//			AuthorSurname: "",
//			AuthorUrl:     "",
//			AuthorPhoto:   "",
//			Text:          "1234",
//			Time:          "",
//			Sticker:       "1234",
//			IsMe:          false,
//		},
//	}
//	error := []string{"private","group","convertion","convertion","sticker"}
//
//	for iter , value := range msgs {
//		mock.ExpectBegin()
//		if error[iter] != "convertion" {
//			groupType := "gch"
//			if error[iter] == "private" {
//				groupType = "pch"
//			}
//
//			if error[iter] != "sticker" {
//				mock.ExpectQuery(`INSERT INTO Messages \(`+groupType+`_id,u_id,txt,send_time\) VALUES\(\$1,\$2,\$3,\$4\) RETURNING msg_id`).WithArgs()
//			}
//		}
//		mock.ExpectCommit()
//	}
//
//}


