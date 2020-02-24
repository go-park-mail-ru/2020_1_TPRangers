package database

import (
	"errors"
	"sync"
	"github.com/satori/go.uuid"
)

type DataInterface interface {
	AddUser(MetaData) (error, MetaData)
	GetUserDataLogin(string) MetaData
	GetUserDataId(int64) MetaData
	DeleteUser(string) error
	EditUser(string, MetaData)
	CheckUser(string) bool
}

type MetaData struct {
	Username  string
	Photo     []byte
	Telephone string
	Password  string
	//Birthday time
}

type Result struct {
	Body interface{} `json:"body,omitempty"`
	Err  string      `json:"err,omitempty"`
}

func NewMetaData(name, tel, pass string, photo []byte) MetaData {
	return MetaData{name, photo, tel, pass}
}

type DataBase struct {
	UserId        map[string]int64
	IdMeta        map[int64]MetaData
	CookieSession map[string]string

	UserCounter int64
	mutex       sync.Mutex
}

func NewDataBase() DataBase {

	return DataBase{UserId: make(map[string]int64), IdMeta: make(map[int64]MetaData), UserCounter: 0}

}

func (db *DataBase) SetCookie(login string) string {

	db.mutex.Lock()
	info , _  := uuid.NewV4()
	cookie := info.String()
	db.CookieSession[login] = cookie
	db.mutex.Unlock()

	return cookie

}

func (db *DataBase) CheckCookie(cookie string, login string) bool {

	return db.CookieSession[login] == cookie

}





func (db *DataBase) AddUser(login string, data MetaData) (error, MetaData) {

	if db.CheckUser(login) {
		return errors.New(`{"error":"такой пользователь уже был зарегестрирован!"}`), MetaData{}
	}

	db.mutex.Lock()
	db.UserId[login] = db.UserCounter
	db.IdMeta[db.UserCounter] = data
	db.UserCounter++
	db.mutex.Unlock()
	return nil, data

}



func (db *DataBase) GetUserDataLogin(login string) MetaData {

	return db.IdMeta[db.UserId[login]]

}

func (db *DataBase) GetUserDataId(id int64) MetaData {

	return db.IdMeta[id]

}

func (db *DataBase) DeleteUser(login string) {

	delete(db.IdMeta, db.UserId[login])
	delete(db.UserId, login)

}

func (db *DataBase) EditUser(login string, data MetaData) {
	db.mutex.Lock()
	db.IdMeta[db.UserId[login]] = data
	db.mutex.Unlock()
}

func (db *DataBase) CheckUser(login string) bool {

	if _, flag := db.UserId[login]; flag {
		return true
	}
	return false
}

func (db *DataBase) CheckAuth(login, password string) error {

	if !db.CheckUser(login) {
		return errors.New(`{"error":"неправильные данные!"}`)
	}

	if db.IdMeta[db.UserId[login]].Password != password {
		return errors.New(`{"error":"неправильные данные!"}`)
	}

	return nil
}
