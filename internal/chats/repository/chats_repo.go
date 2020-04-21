package repository

import (
	"database/sql"
	"fmt"
	"main/internal/models"
	"main/internal/tools/errors"
	"strconv"
)

type ChatRepositoryRealisation struct {
	chatDB *sql.DB
}

func NewChatRepositoryRealisation(db *sql.DB) ChatRepositoryRealisation {
	return ChatRepositoryRealisation{chatDB: db}
}

func (CR ChatRepositoryRealisation) CreateNewChat(chatPhoto, chatName string, users []int) error {

	if len(users) < 0 {
		return nil
	}

	if len(users) <=2 {
		privateValues := make([]interface{},0)
		privateValues = append(privateValues , users[0])

		insertRow := ""
		chatsValues := make([]interface{},0)
		chatsValues = append(chatsValues , users[0])

		if len(users) == 1 {
			insertRow = "INSERT INTO ChatsUsers (u_id,pch_id) VALUES ($1,$2)"
			privateValues = append(privateValues , users[0])
		} else {
			insertRow = "INSERT INTO ChatsUsers (u_id,pch_id) VALUES ($1,$3),($2,$4)"
			chatsValues = append(chatsValues , users[1])
			privateValues = append(privateValues , users[1])
		}

		row := CR.chatDB.QueryRow("INSERT INTO PrivateChats (fu_id,ch_id) VALUES ($1,$2) RETURNING ch_id" , privateValues...)

		chatId := 0
		err := row.Scan(&chatId)

		if err != nil {
			return err
		}

		chatsValues = append(chatsValues , chatId)
		if len(users) == 2 {
			chatsValues = append(chatsValues , chatId)
		}

		_ , err = CR.chatDB.Exec(insertRow, chatsValues...)
		return err
	}


	chatId := 0
	if chatPhoto == "" {
		row := CR.chatDB.QueryRow("INSERT INTO GroupChats (u_id,name) VALUES ($1,$2) RETURNING ch_id", users[0], chatName)
		err := row.Scan(&chatId)

		if err != nil {
			return err
		}
	} else
	{

		photoId := 0
		photoRow := CR.chatDB.QueryRow("INSERT INTO Photos (url) VALUES($1) RETURNING photo_id", chatPhoto)
		err := photoRow.Scan(&photoId)

		if err != nil {
			return err
		}

		row := CR.chatDB.QueryRow("INSERT INTO GroupChats (u_id,name) VALUES ($1,$2) RETURNING ch_id", users[0], chatName)
		err = row.Scan(&chatId)

		if err != nil {
			return err
		}
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

	_, err := CR.chatDB.Exec("INSERT INTO ChatsUsers (gch_id,u_id) VALUES "+insertRow, insertValues...)

	if err != nil {
		return err
	}

	return nil

}

func (CR ChatRepositoryRealisation) ExitChat(chatId int64, userId int) error {

	_ , err := CR.chatDB.Exec("DELETE FROM ChatsUsers WHERE u_id = $1 AND gch_id = $2", userId , chatId)

	return err
}

func (CR ChatRepositoryRealisation) GetChatMessages(chatId int64, userId int) (models.Chat, []models.Message, error) {

	isInChat := 0

	row := CR.chatDB.QueryRow("SELECT u_id FROM ChatsUsers WHERE u_id = $1 AND ch_id = $2", userId, chatId)

	err := row.Scan(&isInChat)

	if err != nil || isInChat != userId {
		fmt.Println(err, isInChat, userId)
		return models.Chat{}, nil, errors.NotExist
	}

	row = CR.chatDB.QueryRow("SELECT CH.ch_id , PH.url , CH.name , COUNT(CU.u_id) AS u_count FROM Chats CH LEFT JOIN Photos PH ON(PH.photo_id=CH.photo_id) INNER JOIN ChatsUsers CU ON(CU.ch_id=CH.ch_id) WHERE CH.ch_id = $1 GROUP BY CH.ch_id , PH.url , CH.name ", chatId)

	chatInfo := new(models.Chat)
	err = row.Scan(&chatInfo.ChatId, &chatInfo.ChatPhoto, &chatInfo.ChatName, &chatInfo.ChatCounter)

	if err != nil {
		fmt.Println("4" , err)
		return models.Chat{}, nil, err
	}

	lastMsgRow := CR.chatDB.QueryRow("SELECT lst_msg_id FROM ChatsUsers WHERE ch_id = $1 AND u_id = $2" , chatInfo.ChatId , userId)
	lastMsgId := int64(0)

	err = lastMsgRow.Scan(&lastMsgId)

	if err != nil {
		fmt.Println(err)
		return models.Chat{}, nil, err
	}

	if chatInfo.ChatCounter == 2 {
		name := ""
		surname := ""
		row = CR.chatDB.QueryRow("SELECT U.name , U.surname , P.url FROM ChatsUsers CU INNER JOIN Users U ON(U.u_id=CU.u_id) INNER JOIN Photos P ON(P.photo_id=u.photo_id) WHERE CU.u_id != $1", userId)
		err = row.Scan(&name, &surname, &chatInfo.ChatPhoto)

		chatInfo.IsGroupChat = false

		if err != nil {
			fmt.Println("3" , err)
			return models.Chat{}, nil, err
		}

		chatInfo.ChatName = name + " " + surname
	} else {
		chatInfo.IsGroupChat = true
	}

	messages := make([]models.Message, 0)

	var msgRows *sql.Rows

	defer func() {
		if msgRows != nil {
			msgRows.Close()
		}
	}()

	iter := 0
	if lastMsgId == 0 {
		iter = 1
		msgRows, err = CR.chatDB.Query("SELECT M.msg_id , M.txt, M.send_time ,U.login , U.name , U.surname ,P.url FROM Messages M INNER JOIN Users U ON(U.u_id=M.u_id) INNER JOIN Photos P ON(P.photo_id = U.photo_id) WHERE M.ch_id = $1 AND M.del_stat = TRUE  ORDER BY M.msg_id DESC  LIMIT 30", chatId)
	} else {
		iter = 2
		msgRows, err = CR.chatDB.Query("SELECT M.msg_id , M.txt, M.send_time ,U.login , U.name , U.surname ,P.url FROM Messages M INNER JOIN Users U ON(U.u_id=M.u_id) INNER JOIN Photos P ON(P.photo_id = U.photo_id) WHERE M.ch_id = $1 AND M.del_stat = TRUE AND M.msg_id < $2 ORDER BY M.msg_id DESC LIMIT 30", chatId, lastMsgId)
	}

	if err != nil {
		fmt.Println(iter, err)
		return models.Chat{}, nil, err
	}

	for msgRows.Next() {

		msg := new(models.Message)
		msgId := int64(0)

		err = msgRows.Scan(&msgId, &msg.Text, &msg.Time, &msg.AuthorUrl, &msg.AuthorName, &msg.AuthorSurname, &msg.AuthorPhoto)

		messages = append(messages, *msg)

	}

	return *chatInfo, messages, nil
}

func (CR ChatRepositoryRealisation) GetAllChats(userId int) ([]models.Chat, error) {



}
