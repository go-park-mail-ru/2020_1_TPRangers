package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	DataBase "./database"
	ET "./errors"
	AP "./json-answers"
	"github.com/gorilla/mux"
)

var FileMaxSize = int64(5 * 1024 * 1024)

var post = DataBase.Post{PostName: "Test Post Name", PostText: "Test Post Text", PostPhoto: "https://picsum.photos/200/300?grayscale"}

type DataHandler struct {
	dataBase      DataBase.UserRepository
	cookieSession DataBase.SessionRepository
}

func getDataFromJson(jsonType string, r *http.Request) (data interface{}, errConvert error) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	switch jsonType {

	case "reg", "data":
		data = new(AP.JsonUserData)
		decoder.Decode(&data)
	case "log":
		data = new(AP.JsonRequestLogin)
		decoder.Decode(&data)
	}

	errConvert = nil
	return
}

func SetCookie(w *http.ResponseWriter, cookieValue string) {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(12 * time.Hour),
	}
	http.SetCookie(*w, &cookie)

	fmt.Println("cookie value: ", cookie.Value)
}

func SetData(data []interface{}, jsonType []string, w http.ResponseWriter) {

	answer := make(map[string]interface{})

	for i, val := range jsonType {
		switch val {

		case "user":
			answer[val] = data[i].(DataBase.MetaData)
		case "feed":
			answer[val] = data[i].([]DataBase.Post)

		// TODO : change case
		case "login":
			answer[val] = data[i].(string)

		}
	}

	json.NewEncoder(w).Encode(&AP.JsonStruct{Body: answer})
	(w).WriteHeader(http.StatusOK)

}

func SetErrors(err error, status int, w http.ResponseWriter) {
	(w).WriteHeader(status)
	json.NewEncoder(w).Encode(&AP.JsonStruct{Err: err.Error()})
}

func (dh DataHandler) Register(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============REGISTER=============\n")
	regData, convertionError := getDataFromJson("reg", r)
	meta := regData.(*AP.JsonUserData)

	if convertionError != nil {
		return
	}

	login := meta.Email
	println(login)

	if dh.dataBase.CheckUser(login) {
		fmt.Println("already registered!")

		SetErrors(errors.New(ET.AlreadyExist), http.StatusConflict, w)
		return
	}

	data := DataBase.NewMetaData(login, meta.Name, meta.Phone, meta.Password, meta.Date, "default photo way")

	(dh.dataBase).AddUser(login, *data)

	cookie := (dh.cookieSession).SetCookie(login)
	SetCookie(&w, cookie)

	w.WriteHeader(http.StatusOK)

	fmt.Println("sucsessfully registered user :", *data)
}

func (dh DataHandler) Login(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Login=============\n")
	mapData, convertionError := getDataFromJson("log", r)
	meta := mapData.(*AP.JsonRequestLogin)

	if convertionError != nil {
		return
	}

	login := meta.Login
	password := meta.Password

	if dh.dataBase.CheckAuth(login, password) != nil {
		fmt.Println("Doesn't exit")
		SetErrors(errors.New(ET.WrongLogin), http.StatusUnauthorized, w)
		return
	} else {
		fmt.Println("OK")
	}

	cookie := (dh.cookieSession).SetCookie(login)
	SetCookie(&w, cookie)
	w.WriteHeader(http.StatusOK)

}

func (dh DataHandler) Logout(w http.ResponseWriter, r *http.Request) {

	fmt.Println("===================LOGOUT===================")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)
}

func (dh DataHandler) Profile(w http.ResponseWriter, r *http.Request) {
	fmt.Print("=============Profile=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		SetErrors(errors.New(ET.CookieExpired), http.StatusUnauthorized, w)
		return
	}

	if login, flag := dh.cookieSession.GetUserByCookie(cookie.Value); flag == nil {

		sendData := make([]interface{}, 1)

		sendData[0] = (dh.dataBase).GetUserDataByLogin(login)

		SetData(sendData, []string{"user"}, w)

	} else {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		SetErrors(errors.New(ET.InvalidCookie), http.StatusUnauthorized, w)
		return
	}

}

func (dh DataHandler) Feed(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Feed=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		SetErrors(errors.New(ET.CookieExpired), http.StatusUnauthorized, w)
		return
	}

	if _, flag := dh.cookieSession.GetUserByCookie(cookie.Value); flag == nil {

		sendData := make([]interface{}, 1)

		sendData[0] = []DataBase.Post{post, post, post, post, post}

		SetData(sendData, []string{"feed"}, w)

	} else {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		SetErrors(errors.New(ET.InvalidCookie), http.StatusUnauthorized, w)
		return
	}

}

func (dh DataHandler) SettingsGet(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============SettingsGET=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		SetErrors(errors.New(ET.CookieExpired), http.StatusUnauthorized, w)
		return
	}

	login, err := dh.cookieSession.GetUserByCookie(cookie.Value)

	if err != nil {

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		SetErrors(errors.New(ET.InvalidCookie), http.StatusUnauthorized, w)
		return

	}

	sendData := make([]interface{}, 1)

	sendData[0] = (dh.dataBase).GetUserDataByLogin(login)

	SetData(sendData, []string{"user"}, w)

}

func (dh DataHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Print("=============GETUSER=============\n")
	login := r.Header.Get("X-User")

	sendData := make([]interface{}, 2)
	sendData[0] = (dh.dataBase).GetUserDataByLogin(login)
	sendData[1] = []DataBase.Post{post, post, post, post, post}

	SetData(sendData, []string{"user", "feed"}, w)

}

func (dh DataHandler) SettingsUpload(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============UploadSettings=============\n")

	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		SetErrors(errors.New(ET.CookieExpired), http.StatusUnauthorized, w)
		return
	}

	login, err := dh.cookieSession.GetUserByCookie(cookie.Value)

	if err != nil {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		SetErrors(errors.New(ET.InvalidCookie), http.StatusUnauthorized, w)
		return
	}

	err = r.ParseMultipartForm(FileMaxSize)

	uploadDataFlags := []string{"uploadedFile", "email", "password", "name", "phone", "date"}

	currentUserData := dh.dataBase.GetUserDataByLogin(login)

	for _, dataFlag := range uploadDataFlags {
		switch dataFlag {
		case "uploadedFile":
			if data := r.FormValue(dataFlag); data != "" {
				currentUserData.Photo = data
			}
		case "email":
			if data := r.FormValue(dataFlag); data != "" {
				currentUserData.Email = data
			}
		case "password":
			if data := r.FormValue(dataFlag); data != "" {
				currentUserData.Password = data
			}
		case "name":
			if data := r.FormValue(dataFlag); data != "" {
				currentUserData.Username = data
			}
		case "phone":
			if data := r.FormValue(dataFlag); data != "" {
				currentUserData.Telephone = data
			}
		case "date":
			if data := r.FormValue(dataFlag); data != "" {
				currentUserData.Date = data
			}

		}
	}

	dh.dataBase.EditUser(login, currentUserData)

	sendData := make([]interface{}, 1)

	sendData[0] = currentUserData

	SetData(sendData, []string{"user"}, w)
}

func main() {
	fmt.Print("main")
	server := mux.NewRouter()
	db := DataBase.NewDataBase()
	cb := DataBase.NewCookieSession()
	fmt.Printf("%T", post)
	server.Use(SetCorsMiddleware(server))

	api := &(DataHandler{dataBase: db, cookieSession: cb})
	DataBase.FillDataBase(db)

	server.HandleFunc("/api/v1/news", api.Feed).Methods("GET", "OPTIONS")
	server.HandleFunc("/api/v1/profile", api.Profile).Methods("GET", "OPTIONS")
	server.HandleFunc("/api/v1/settings", api.SettingsGet).Methods("GET", "OPTIONS")
	server.HandleFunc("/api/v1/user", api.GetUser).Methods("GET", "OPTIONS")

	server.HandleFunc("/api/v1/registration", api.Register).Methods("POST", "OPTIONS")
	server.HandleFunc("/api/v1/login", api.Login).Methods("POST", "OPTIONS")

	server.HandleFunc("/api/v1/login", api.Logout).Methods("DELETE", "OPTIONS")

	server.HandleFunc("/api/v1/settings", api.SettingsUpload).Methods("PUT", "OPTIONS")

	http.ListenAndServe(":3001", server)

}

func SetCorsMiddleware(r *mux.Router) mux.MiddlewareFunc {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			//TODO: убрать из корса
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, PUT, DELETE, POST")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Login, Set-Cookie, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, csrf-token, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Vary", "Cookie")

			if req.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, req)
		})
	}

}
