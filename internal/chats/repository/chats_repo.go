package repository

import (
	"database/sql"
	"fmt"
	"main/models"
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
	if len(users) <= 2 {
		privateValues := make([]interface{}, 0)
		privateValues = append(privateValues, users[0])

		insertRow := ""
		chatsValues := make([]interface{}, 0)
		chatsValues = append(chatsValues, users[0])

		if len(users) == 1 {
			insertRow = "INSERT INTO ChatsUsers (u_id,pch_id) VALUES ($1,$2)"
			privateValues = append(privateValues, users[0])
		} else {
			insertRow = "INSERT INTO ChatsUsers (u_id,pch_id) VALUES ($1,$3),($2,$4)"
			chatsValues = append(chatsValues, users[1])
			privateValues = append(privateValues, users[1])
		}
		fmt.Println(privateValues)
		row := CR.chatDB.QueryRow("INSERT INTO PrivateChats (fu_id,su_id) VALUES ($1,$2) RETURNING ch_id", privateValues...)

		chatId := 0
		err := row.Scan(&chatId)

		if err != nil {
			return err
		}

		chatsValues = append(chatsValues, chatId)
		if len(users) == 2 {
			chatsValues = append(chatsValues, chatId)
		}

		_, err = CR.chatDB.Exec(insertRow, chatsValues...)
		return err
	}

	chatId := 0
	if chatPhoto == "" {
		row := CR.chatDB.QueryRow("INSERT INTO GroupChats (u_id,name) VALUES ($1,$2) RETURNING ch_id", users[0], chatName)
		err := row.Scan(&chatId)

		if err != nil {
			return err
		}
	} else {

		photoId := 0
		photoRow := CR.chatDB.QueryRow("INSERT INTO Photos (url) VALUES($1) RETURNING photo_id", chatPhoto)
		err := photoRow.Scan(&photoId)

		if err != nil {
			return err
		}

		row := CR.chatDB.QueryRow("INSERT INTO GroupChats (u_id,name,photo_id) VALUES ($1,$2,$3) RETURNING ch_id", users[0], chatName, photoId)
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

	_, err := CR.chatDB.Exec("DELETE FROM ChatsUsers WHERE u_id = $1 AND gch_id = $2", userId, chatId)

	return err
}

func (CR ChatRepositoryRealisation) GetPrivateChatMessages(chatId int64, userId int) (models.ChatInfo, []models.Message, error) {
	chatInfo := new(models.ChatInfo)
	chatInfo.IsGroupChat = false
	chatInfo.ChatId = strconv.FormatInt(chatId, 10)

	chatQuery := CR.chatDB.QueryRow("SELECT U.name,U.surname,U.login,P.url FROM PrivateChats PCH INNER JOIN Users U ON CASE "+
		"WHEN PCH.su_id != $2 AND PCH.su_id=U.u_id THEN 1 "+
		"WHEN PCH.fu_id != $2 AND PCH.fu_id=U.u_id THEN 1 "+
		"ELSE 0 "+
		"END = 1 "+
		"INNER JOIN Photos P ON(P.photo_id=U.photo_id) "+
		"WHERE PCH.ch_id = $1 GROUP BY U.name,U.surname,U.login,P.url", chatId, userId)

	err := chatQuery.Scan(&chatInfo.PrivateName, &chatInfo.PrivateSurname, &chatInfo.PrivateUrl, &chatInfo.ChatPhoto)

	if err != nil {
		fmt.Println(err)
		return models.ChatInfo{}, nil, err
	}

	msgQuery, err := CR.chatDB.Query("SELECT M.u_id , M.txt, M.send_time,U.name,U.surname,U.login,P.url FROM Messages M INNER JOIN Users U ON(U.u_id=M.u_id) INNER JOIN Photos P ON(P.photo_id=U.photo_id) WHERE M.pch_id = $1 ORDER BY M.send_time DESC", chatId)

	if err != nil {
		return models.ChatInfo{}, nil, err
	}

	msgs := make([]models.Message, 0)

	for msgQuery.Next() {

		var uId *int
		msg := new(models.Message)

		err := msgQuery.Scan(&uId, &msg.Text, &msg.Time, &msg.AuthorName, &msg.AuthorSurname, &msg.AuthorUrl, &msg.AuthorPhoto)

		if *uId != userId {
			msg.IsMe = false
		} else {
			msg.IsMe = true
		}

		if err != nil {
			return models.ChatInfo{}, nil, err
		}

		msgs = append(msgs, *msg)
	}

	return *chatInfo, msgs, nil
}

func (CR ChatRepositoryRealisation) GetGroupChatMessages(chatId int64, userId int) (models.ChatInfo, []models.Message, error) {
	chatInfo := new(models.ChatInfo)
	chatInfo.IsGroupChat = true
	chatInfo.ChatId = "c" + strconv.FormatInt(chatId, 10)

	chatQuery := CR.chatDB.QueryRow("SELECT GC.name,P.url,COUNT(CU.u_id) FROM GroupChats GC INNER JOIN Photos P ON(P.photo_id=GC.photo_id) INNER JOIN ChatsUsers CU ON(CU.gch_id=GC.ch_id) WHERE GC.ch_id = $1 GROUP BY GC.name,P.url", chatId)

	err := chatQuery.Scan(&chatInfo.ChatName, &chatInfo.ChatPhoto, &chatInfo.ChatCounter)

	if err != nil {
		fmt.Println(err)
		return models.ChatInfo{}, nil, err
	}

	msgQuery, err := CR.chatDB.Query("SELECT M.u_id , M.txt, M.send_time,U.name,U.surname,U.login,P.url FROM Messages M INNER JOIN Users U ON(U.u_id=M.u_id) INNER JOIN Photos P ON(P.photo_id=U.photo_id) WHERE M.gch_id = $1 ORDER BY M.send_time DESC", chatId)

	if err != nil {
		fmt.Println(err, "here")
		return models.ChatInfo{}, nil, err
	}

	msgs := make([]models.Message, 0)

	for msgQuery.Next() {

		var uId *int
		msg := new(models.Message)

		err := msgQuery.Scan(&uId, &msg.Text, &msg.Time, &msg.AuthorName, &msg.AuthorSurname, &msg.AuthorUrl, &msg.AuthorPhoto)

		if *uId != userId {
			msg.IsMe = false
		} else {
			msg.IsMe = true
		}

		if err != nil {
			return models.ChatInfo{}, nil, err
		}

		msgs = append(msgs, *msg)
	}

	return *chatInfo, msgs, nil
}

func (CR ChatRepositoryRealisation) GetAllChats(userId int) ([]models.Chat, error) {

	queryRow := "WITH MESS AS " +
		"( " +
		"SELECT max(M.msg_id) as msg_id, GC.ch_id as cg_id FROM ChatsUsers CU " +
		"LEFT JOIN PrivateChats PC ON(PC.ch_id=CU.pch_id) " +
		"LEFT JOIN Users U ON CASE " +
		"WHEN PC.fu_id != $1 AND PC.fu_id= U.u_id  THEN 1 " +
		"WHEN PC.su_id != $1 AND PC.su_id = U.u_id THEN 1 " +
		"WHEN PC.fu_id = $1 AND PC.su_id = $1 AND U.u_id=PC.fu_id THEN 1 " +
		"ELSE 0 " +
		"END = 1 " +
		"LEFT JOIN GroupChats GC ON(GC.ch_id=CU.gch_id) " +
		"LEFT JOIN Messages M ON CASE " +
		"WHEN M.pch_id = PC.ch_id THEN 1 " +
		"WHEN M.gch_id = GC.ch_id THEN 1 " +
		"ELSE 0 " +
		"END = 1 " +
		"WHERE CU.u_id = $1 GROUP BY CU.cu_id,PC.ch_id,GC.ch_id " +
		") " +
		"SELECT PC.ch_id,GC.ch_id,GC.name,U.name,U.surname,U.login,CP.url,P.url,max(M.send_time),M.txt FROM ChatsUsers CU " +
		"LEFT JOIN PrivateChats PC ON(PC.ch_id=CU.pch_id) " +
		"LEFT JOIN Users U ON CASE " +
		"WHEN PC.fu_id != $1 AND PC.fu_id= U.u_id  THEN 1 " +
		"WHEN PC.su_id != $1 AND PC.su_id = U.u_id THEN 1 " +
		"WHEN PC.fu_id = $1 AND PC.su_id = $1 AND U.u_id=PC.fu_id THEN 1 " +
		"ELSE 0 " +
		"END = 1 " +
		"LEFT JOIN GroupChats GC ON(GC.ch_id=CU.gch_id) " +
		"LEFT JOIN Messages M ON CASE " +
		"WHEN M.pch_id = PC.ch_id THEN 1 " +
		"WHEN M.gch_id = GC.ch_id THEN 1 " +
		"ELSE 0 " +
		"END = 1 " +
		"INNER JOIN MESS ON CASE " +
		"WHEN MESS.msg_id=M.msg_id THEN 1 " +
		"WHEN MESS.msg_id IS NULL AND M.msg_id IS NULL THEN 1 " +
		"ELSE 0 " +
		"END = 1 " +
		"LEFT JOIN Users UM ON(UM.u_id=M.u_id) LEFT JOIN Photos P ON(P.photo_id=UM.photo_id) " +
		"LEFT JOIN Photos CP ON CASE " +
		"WHEN CU.pch_id != 0 AND U.photo_id = CP.photo_id THEN 1 " +
		"WHEN CU.gch_id != 0 AND GC.photo_id = CP.photo_id THEN 1 " +
		"ELSE 0 " +
		"END = 1 " +
		"WHERE CU.u_id = $1 GROUP BY PC.ch_id,GC.ch_id,GC.name,U.name,U.surname,U.login,CP.url,P.url,M.txt ORDER BY max(M.send_time) DESC"

	chatRow, err := CR.chatDB.Query(queryRow, userId)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	chats := make([]models.Chat, 0)

	for chatRow.Next() {

		var (
			isPrivate *int64
			isGroup   *int64
			chatName  *string
			prName    *string
			prSurname *string
			prLogin   *string
		)

		chat := new(models.Chat)
		err = chatRow.Scan(&isPrivate, &isGroup, &chatName, &prName, &prSurname, &prLogin, &chat.ChatPhoto, &chat.LastMessageAuthorPhoto, &chat.LastMessageTime, &chat.LastMessageTxt)
		if isPrivate == nil {
			chat.IsGroupChat = true
			chat.ChatId = "c" + strconv.FormatInt(*isGroup, 10)
			chat.ChatName = *chatName
		} else {
			chat.IsGroupChat = false
			chat.ChatId = strconv.FormatInt(*isPrivate, 10)
			chat.PrivateName = *prName
			chat.PrivateSurname = *prSurname
			chat.PrivateUrl = *prLogin
		}

		chats = append(chats, *chat)
	}

	return chats, nil
}
