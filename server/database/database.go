package database

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	AddUser(string, MetaData)
	GetUserDataByLogin(string) MetaData
	GetUserDataById(int64) MetaData
	DeleteUser(string) error
	EditUser(string, MetaData)
	CheckUser(string) bool
	CheckAuth(string, string) error
	GetPasswordByLogin(string) string
}

type SessionRepository interface {
	SetCookie(string) string
	CheckCookie(string, string) bool
	GetUserByCookie(string) (string, error)
}

type SessionData struct {
	mutex         sync.RWMutex
	CookieSession map[string]string
}

type Post struct {
	PostName  string
	PostText  string
	PostPhoto string
}

//  хэшировать пароли
type DataBase struct {
	mutex  sync.RWMutex
	userId map[string]int64
	idMeta map[int64]MetaData

	UserCounter int64
}

type MetaData struct {
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	Photo     string `json:"photo,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Password  string `json:"password,omitempty"`
	Date      string `json:"date,omitempty"`
}

func NewMetaData(login, name, tel, pass, date, photo string) *MetaData {
	return &MetaData{login, name, photo, tel, pass, date}
}

func NewDataBase() *DataBase {

	return &DataBase{userId: make(map[string]int64), idMeta: make(map[int64]MetaData), UserCounter: 0}

}

func NewCookieSession() *SessionData {
	return &SessionData{CookieSession: make(map[string]string)}
}

func (db *SessionData) SetCookie(login string) string {

	db.mutex.Lock()
	info, _ := uuid.NewV4()
	cookie := info.String()
	db.CookieSession[cookie] = login

	db.mutex.Unlock()

	return cookie

}

func (db *SessionData) GetUserByCookie(cookie string) (string, error) {

	db.mutex.RLock()

	userName, sessionExist := db.CookieSession[cookie]

	db.mutex.RUnlock()

	if sessionExist {
		return userName, nil
	}

	return "", errors.New("неверные куки!")
}

func (db *SessionData) CheckCookie(cookie string, login string) bool {

	db.mutex.RLock()
	val, flag := db.CookieSession[cookie]
	db.mutex.RUnlock()

	return val == login && flag
}

func (db *DataBase) GetPasswordByLogin(login string) string {

	db.mutex.RLock()
	password := db.idMeta[db.userId[login]].Password
	db.mutex.RUnlock()

	return password
}

func (db *DataBase) AddUser(login string, data MetaData) {

	db.mutex.Lock()
	db.userId[login] = db.UserCounter
	db.idMeta[db.UserCounter] = data
	db.UserCounter++
	db.mutex.Unlock()

}

func (db *DataBase) GetUserDataByLogin(login string) MetaData {

	db.mutex.RLock()
	data := db.idMeta[db.userId[login]]
	db.mutex.RUnlock()

	return data

}

func (db *DataBase) GetUserDataById(id int64) MetaData {

	db.mutex.RLock()
	data := db.idMeta[id]
	db.mutex.RUnlock()

	return data

}

func (db *DataBase) DeleteUser(login string) error {

	existFlag := db.CheckUser(login)
	db.mutex.Lock()
	var err error
	if !existFlag {
		err = errors.New("такого пользователя не существует!")
	} else {
		delete(db.idMeta, db.userId[login])
		delete(db.userId, login)
	}

	db.mutex.Unlock()

	return err

}

func (db *DataBase) EditUser(login string, data MetaData) {
	db.mutex.Lock()
	db.idMeta[db.userId[login]] = data
	db.mutex.Unlock()
}

func (db *DataBase) CheckUser(login string) bool {

	db.mutex.RLock()
	_, doesExist := db.userId[login]
	db.mutex.RUnlock()

	return doesExist
}

func (db *DataBase) CheckAuth(login, password string) error {

	if !db.CheckUser(login) {
		err := errors.New("неправильные данные!")
		return err
	}

	db.mutex.RLock()
	var err error
	if db.idMeta[db.userId[login]].Password != password {
		err = errors.New("неправильные данные!")
	}
	db.mutex.RUnlock()
	return err
}

func FillDataBase(dataInterface UserRepository) {

	sliceMail := []string{"asdasd@yandex.ru", "123@yandex.ru", "znajderko@yandex.ru"}

	for _, val := range sliceMail {
		defData := NewMetaData(val, "TEST", "88005553535", "TEST", "00.00.2000", "./fileWay")
		dataInterface.AddUser(val, *defData)
	}

}
