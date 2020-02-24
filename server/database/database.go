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
}

type CookieInterface interface {
	SetCookie(string) string
	CheckCookie(string, string) bool
	GetUser(string) (string,error)
}

type CookieData struct {
	CookieSession map[string]string

	mutex sync.Mutex
}

//  хэшировать пароли
type DataBase struct {
	UserId map[string]int64
	IdMeta map[int64]MetaData

	UserCounter int64
	mutex       sync.Mutex
}
type MetaData struct {
	Username  string
	Photo     []byte
	Telephone string
	Password  string
	//Birthday time
}

func NewMetaData(name, tel, pass string, photo []byte) *MetaData {
	return &MetaData{name, photo, tel, pass}
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

func (db *CookieData) GetUser(cookie string) (string , error){

	if _ , flag := db.CookieSession[cookie] ; flag{
		return "",errors.New("неверные куки!")
	}

	return db.CookieSession[cookie] , nil
}

func (db *CookieData) CheckCookie(cookie string, login string) bool {

	if val, flag := db.CookieSession[cookie] ; val == login && flag{
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
	defData := NewMetaData("TEST", "88005553535", "TEST", make([]byte, 16))

	for _, val := range sliceMail {
		dataInterface.AddUser(val, *defData)
	}

}
