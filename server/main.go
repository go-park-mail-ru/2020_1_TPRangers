package main

import (
	DB "./database"
	AP "./json-answers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type DataHandler struct {
	dataBase   DB.DataInterface
	cookieBase DB.CookieInterface
}

var FileMaxSize = int64(5 * 1024 * 1024)

func makeCorsHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Content-Type", "application/json")
}

func (dh DataHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Print("=============REGISTER=============\n")
	makeCorsHeaders(&w)
	data := DB.NewMetaData("xd", "xd", "xd", make([]byte, 2))
	login := "nikita"
	// test data


	if err, info := (dh.dataBase).AddUser(login, *data); err != nil {
		http.Error(w, `{"err":"неправильные данные!"}`, 401)
		return
	} else {
		answerData := make(map[string]interface{})
		answerData["isAuth"] = true
		answerData["data"] = info
		json.NewEncoder(w).Encode(&AP.JsonStruct{Body: answerData})
		cValue := (dh.cookieBase).SetCookie(login)
		cookie := http.Cookie{
			Name:    "session_id",
			Value:   cValue,
			Expires: time.Now().Add(12 * time.Hour),
		}
		http.SetCookie(w, &cookie)
	}
}

func (dh DataHandler) Login(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Login=============\n")
	makeCorsHeaders(&w)
}

func (dh DataHandler) Profile(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Profile=============\n")
	makeCorsHeaders(&w)
}

func (dh DataHandler) Feed(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Feed=============\n")
	makeCorsHeaders(&w)

}

func (dh DataHandler) Settings(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Settings=============\n")
	makeCorsHeaders(&w)

}

func (dh DataHandler) PhotoUpload(w http.ResponseWriter, r *http.Request){
	fmt.Print("=============PhotoUpload=============\n")
	makeCorsHeaders(&w)
	cookie, err:= r.Cookie("session_id")

	if err != nil {
		http.Error(w, `{"err":"истёкшие куки!"}`, 401)
		return
	}

	if login , flag := dh.cookieBase.GetUser(cookie.Value); flag != nil {

		r.ParseMultipartForm(FileMaxSize)
		file, header , err := r.FormFile("my_file")
		if err != nil {
			http.Error(w, `{"err":"неверный формат файла!"}`, 401)
			return
		}
		defer file.Close()

		userData := dh.dataBase.GetUserDataLogin(login)

		photoByte , fileErr := ioutil.ReadAll(file)

		if fileErr!= nil{
			http.Error(w, `{"err":"файл не может быть сохранён!"}`, 500)
			return
		}

		userData.Photo = photoByte

		dh.dataBase.EditUser(login,userData)
		size := r.Header.Get("Content-Length")
		//  отправка фотки
		w.Header().Set("Content-Disposition", "attachment; filename="+ header.Filename)
		w.Header().Set("Content-Type", http.DetectContentType([]byte(header.Filename)))
		w.Header().Set("Content-Length", size)
		io.Copy(w,file)

	} else {

		http.Error(w, `{"err":"неверная сессия!"}`, 401)
		return

	}
}



func main() {
	fmt.Print("main")
	server := mux.NewRouter()
	db := DB.NewDataBase()
	cb := DB.NewCookieBase()
	DB.FillDataBase(db)
	api := &(DataHandler{dataBase: db, cookieBase: cb})

	server.HandleFunc("/feed", api.Feed)
	server.HandleFunc("/profile", api.Profile)
	server.HandleFunc("/register", api.Register)
	server.HandleFunc("/login", api.Login)
	server.HandleFunc("/settings", api.Settings)

	http.ListenAndServe(":3001", server)

}
