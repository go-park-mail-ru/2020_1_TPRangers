package database

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type DataInterface interface {
	AddUser(string, MetaData) (error, MetaData)
	GetUserDataLogin(string) MetaData
	GetUserDataId(int64) MetaData
	DeleteUser(string) error
	EditUser(string, MetaData)
	CheckUser(string) bool
	CheckAuth(string, string) error
	GetPasswordByLogin(string) string
}

type CookieInterface interface {
	SetCookie(string) string
	CheckCookie(string, string) bool
	GetUser(string) (string, error)
}

type CookieData struct {
	CookieSession map[string]string

	mutex sync.Mutex
}

type Post struct {
	PostName  string
	PostText  string
	PostPhoto string
}

//  хэшировать пароли
type DataBase struct {
	UserId map[string]int64
	IdMeta map[int64]MetaData

	UserCounter int64
	mutex       sync.Mutex
}
type MetaData struct {
	Email string
	Username  string
	Photo     []byte
	Telephone string
	Password  string
	Date      string
}

func NewMetaData(login, name, tel, pass, date string, photo []byte) *MetaData {
	return &MetaData{login,name, photo, tel, pass, date}
}

func MergeData(dataLeft, dataRight MetaData) MetaData {
	dataLeft.Email = dataRight.Email
	dataLeft.Password = dataRight.Password
	dataLeft.Username = dataRight.Username
	dataLeft.Telephone = dataRight.Telephone
	dataLeft.Date = dataRight.Date

	return dataLeft

}

func NewDataBase() *DataBase {

	return &DataBase{UserId: make(map[string]int64), IdMeta: make(map[int64]MetaData), UserCounter: 0}

}

func NewCookieBase() *CookieData {
	return &CookieData{CookieSession: make(map[string]string)}
}

func (db *CookieData) SetCookie(login string) string {

	db.mutex.Lock()
	info, _ := uuid.NewV4()
	cookie := info.String()
	db.CookieSession[cookie] = login

	db.mutex.Unlock()

	return cookie

}

func (db *DataBase) GetPasswordByLogin(login string) string {
	return db.IdMeta[db.UserId[login]].Password
}

func (db *CookieData) GetUser(cookie string) (string, error) {

	if val, _ := db.CookieSession[cookie]; val != "" {
		return db.CookieSession[cookie], nil

	}
	return "", errors.New("неверные куки!")
}

func (db *CookieData) CheckCookie(cookie string, login string) bool {

	if val, flag := db.CookieSession[cookie]; val == login && flag {
		return true
	}
	return false

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

func (db *DataBase) DeleteUser(login string) error {

	if !db.CheckUser(login) {
		return errors.New("such user doesnt exist!")
	}

	delete(db.IdMeta, db.UserId[login])
	delete(db.UserId, login)

	return nil

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

func FillDataBase(dataInterface DataInterface) {

	sliceMail := []string{"asdasd@yandex.ru", "123@yandex.ru", "znajderko@yandex.ru"}

	for _, val := range sliceMail {
		defData := NewMetaData(val,"TEST", "88005553535", "TEST", "00.00.2000", make([]byte, 16))
		dataInterface.AddUser(val, *defData)
	}

}
