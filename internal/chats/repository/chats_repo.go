package repository

import (
	"database/sql"
	"fmt"
	"main/internal/models"
	"strconv"
	"main/internal/tools/errors"
)

type ChatRepositoryRealisation struct {
	chatDB *sql.DB
}

func NewChatRepositoryRealisation(db *sql.DB) ChatRepositoryRealisation {
	return ChatRepositoryRealisation{chatDB: db}
}

func (CR ChatRepositoryRealisation) CreateNewChat(chatPhoto, chatName string, users []int) error {

	photoId := 0

	if chatPhoto != "" {
		row := CR.chatDB.QueryRow("INSERT INTO Photos (url , photos_likes_count) VALUES ($1,$2) RETURNING photo_id", chatPhoto, 0)

		err := row.Scan(&photoId)

		if err != nil {
			return err
		}

	}

	chatId := int64(0)

	var row *sql.Row

	if photoId == 0 {
		row = CR.chatDB.QueryRow("INSERT INTO Chats (name) VALUES ($1) RETURNING ch_id", chatName)
	} else {
		row = CR.chatDB.QueryRow("INSERT INTO Chats (name , photo_id) VALUES ($1,$2) RETURNING ch_id", chatName, photoId)
	}

	err := row.Scan(&chatId)

	if err != nil {
		return err
	}

	iter := 2
	insertRow := ""
	insertValues := make([]interface{}, 0)

	for i, _ := range users {

		insertRow += "($" + strconv.Itoa(iter-1) + ",$" + strconv.Itoa(iter) + "),"
		insertValues = append(insertValues, chatId, users[i])
		iter += 2

	}

	insertRow = insertRow[:len(insertRow)-1]

	_, err = CR.chatDB.Exec("INSERT INTO ChatsUsers (ch_id,u_id) VALUES "+insertRow, insertValues...)

	if err != nil {
		return err
	}

	return nil
}

func (CR ChatRepositoryRealisation) ExitChat(chatId int64 , userId int) error {
	_ , err := CR.chatDB.Exec("DELETE FROM ChatsUsers WHERE ch_id = $1 , u_id = $2" ,chatId , userId)

	return err
}

func (CR ChatRepositoryRealisation) GetChatMessages(chatId int64 , userId int) error {

	isInChat := 0

	row := CR.chatDB.QueryRow("SELECT u_id FROM ChatsUsers WHERE u_id = $1 AND ch_id = $2" , userId , chatId)

	err := row.Scan(&isInChat)

	if err != nil || isInChat == userId {
		fmt.Println(err , isInChat , userId)
		return errors.NotExist
	}

	row = CR.chatDB.QueryRow("SELECT CH.ch_id , PH.url , CH.name , COUNT(CU.u_id) AS u_count FROM Chats CH LEFT JOIN Photos PH ON(PH.photo_id=CH.photo_id) INNER JOIN ChatsUsers CU ON(CU.ch_id=CH.ch_id) WHERE CH.ch_id = $1" , chatId)

	chatInfo := new(models.Chat)
	err = row.Scan(&chatInfo.ChatId , &chatInfo.ChatPhoto , &chatInfo.ChatName , &chatInfo.ChatCounter)

	if err != nil {
		return err
	}

	if chatInfo.ChatCounter == 2 {
		row = CR.chatDB.QueryRow("SELECT U.name , U.surname , P.url FROM ChatsUsers CU INNER JOIN Users U ON(U.u_id=CU.u_id) INNER JOIN Photos P ON(P.photo_id=u.photo_id) WHERE CU.u_id != $1" , userId)

	}

}
