package main

import (
	"encoding/json"

	DataBase "./database"
	ET "./errors"
	AP "./json-answers"

	//"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var FileMaxSize = int64(5 * 1024 * 1024)

var post = DataBase.Post{PostName: "Test Post Name", PostText: "Test Post Text", PostPhoto: "https://picsum.photos/200/300?grayscale"}

type DataHandler struct {
	dataBase   DataBase.DataInterface
	cookieBase DataBase.CookieInterface
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

func SetData(data []interface{}, jsonType []string, w *http.ResponseWriter) {

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

	json.NewEncoder(*w).Encode(&AP.JsonStruct{Body: answer})
	(*w).WriteHeader(http.StatusOK)

}

func SetErrors(err []string, status int, w *http.ResponseWriter) {
	(*w).WriteHeader(status)
	json.NewEncoder(*w).Encode(&AP.JsonStruct{Err: err})
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

		SetErrors([]string{ET.AlreadyExistError}, http.StatusConflict, &w)
		return
	}

	data := DataBase.NewMetaData(login, meta.Name, meta.Phone, meta.Password, meta.Date, make([]byte, 0))
	err, info := (dh.dataBase).AddUser(login, *data)

	fmt.Println("login err is : ", err)

	cookie := (dh.cookieBase).SetCookie(login)
	SetCookie(&w, cookie)

	w.WriteHeader(http.StatusOK)

	fmt.Println("sucsessfully registered user :", info)

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
		SetErrors([]string{ET.WrongLogin}, http.StatusUnauthorized, &w)
		return
	} else {
		fmt.Println("OK")
	}

	cookie := (dh.cookieBase).SetCookie(login)
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

func (dh DataHandler) PhotoUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Print("=============PhotoUpload=============\n")

	cookie, err := r.Cookie("session_id")

	fmt.Println(r)

	if err == http.ErrNoCookie {
		SetErrors([]string{ET.CookieExpiredError}, http.StatusUnauthorized, &w)
		return
	}

	if login, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {

		err := r.ParseMultipartForm(FileMaxSize)

		if err != nil {
			fmt.Println(err)
			SetErrors([]string{ET.FileError}, http.StatusBadRequest, &w)
			return
		}

		fmt.Println(r)
		file, _, err := r.FormFile("uploadedFile")

		fmt.Println(r.Body)
		defer r.Body.Close()
		fmt.Println(file, err)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			SetErrors([]string{ET.FileError}, http.StatusBadRequest, &w)
			return
		}
		defer file.Close()

		JsonUserData := dh.dataBase.GetUserDataLogin(login)
		photoByte, fileErr := ioutil.ReadAll(file)

		if fileErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			SetErrors([]string{ET.FileSavingError}, http.StatusBadRequest, &w)
			return
		}

		JsonUserData.Photo = photoByte
		dh.dataBase.EditUser(login, JsonUserData)

		sendData := make([]interface{}, 1)

		sendData[0] = DataBase.MetaData{}

		SetData(sendData, []string{"user"}, &w)

	} else {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		SetErrors([]string{ET.WrongCookie}, http.StatusUnauthorized, &w)
		return

	}
}

func (dh DataHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Print("=============Profile=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		SetErrors([]string{ET.CookieExpiredError}, http.StatusUnauthorized, &w)
		return
	}

	if login, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {

		sendData := make([]interface{}, 1)

		sendData[0] = (dh.dataBase).GetUserDataLogin(login)

		SetData(sendData, []string{"user"}, &w)

	} else {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		SetErrors([]string{ET.WrongCookie}, http.StatusUnauthorized, &w)
		return
	}

}

func (dh DataHandler) GetNews(w http.ResponseWriter, r *http.Request) {
	fmt.Print("=============Profile=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		SetErrors([]string{ET.CookieExpiredError}, http.StatusUnauthorized, &w)
		return
	}

	if _, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {

		sendData := make([]interface{}, 1)

		sendData[0] = []DataBase.Post{post, post, post, post, post}

		SetData(sendData, []string{"feed"}, &w)

	} else {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		SetErrors([]string{ET.WrongCookie}, http.StatusUnauthorized, &w)
		return
	}

}

func (dh DataHandler) Feed(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Feed=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		SetErrors([]string{ET.CookieExpiredError}, http.StatusUnauthorized, &w)
		return
	}

	if _, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {

		sendData := make([]interface{}, 1)

		sendData[0] = []DataBase.Post{post, post, post, post, post}

		SetData(sendData, []string{"feed"}, &w)

	} else {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		SetErrors([]string{ET.WrongCookie}, http.StatusUnauthorized, &w)
		return
	}

}

func (dh DataHandler) SettingsGet(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============SettingsGET=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		SetErrors([]string{ET.CookieExpiredError}, http.StatusUnauthorized, &w)
		return
	}

	if login, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {

		sendData := make([]interface{}, 1)

		sendData[0] = (dh.dataBase).GetUserDataLogin(login)

		SetData(sendData, []string{"user"}, &w)

	} else {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		SetErrors([]string{ET.WrongCookie}, http.StatusUnauthorized, &w)
		return
	}

}

func (dh DataHandler) SettingsPost(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============SettingsPOST=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		SetErrors([]string{ET.CookieExpiredError}, http.StatusUnauthorized, &w)
		return
	}

	if login, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {

		mapData, convertionError := getDataFromJson("data", r)
		meta := mapData.(*AP.JsonUserData)

		if convertionError != nil {
			return
		}

		newData := *DataBase.NewMetaData(login, meta.Name, meta.Phone, meta.Password, meta.Date, make([]byte, 0))
		newData = DataBase.MergeData(dh.dataBase.GetUserDataLogin(login), newData)
		dh.dataBase.EditUser(login, newData)

		sendData := make([]interface{}, 1)

		sendData[0] = newData

		SetData(sendData, []string{"user"}, &w)

		fmt.Println(w)

	} else {
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		SetErrors([]string{ET.WrongCookie}, http.StatusUnauthorized, &w)
		return
	}

}

func (dh DataHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Print("=============GETUSER=============\n")
	fmt.Println(r)
	login := r.Header.Get("X-User")

	sendData := make([]interface{}, 2)
	sendData[0] = (dh.dataBase).GetUserDataLogin(login)
	sendData[1] = []DataBase.Post{post, post, post, post, post}

	SetData(sendData, []string{"user", "feed"}, &w)

}

func main() {
	fmt.Print("main")
	server := mux.NewRouter()
	db := DataBase.NewDataBase()
	cb := DataBase.NewCookieBase()
	fmt.Printf("%T", post)
	server.Use(SetCorsMiddleware(server))

	api := &(DataHandler{dataBase: db, cookieBase: cb})
	DataBase.FillDataBase(db)

	server.HandleFunc("/api/v1/news", api.Feed).Methods("GET", "OPTIONS")            //
	server.HandleFunc("/api/v1/profileUser", api.GetProfile).Methods("GET", "OPTIONS")
	server.HandleFunc("/api/v1/profileNews", api.GetNews).Methods("GET", "OPTIONS")
	server.HandleFunc("/api/v1/settings", api.SettingsGet).Methods("GET", "OPTIONS") //
	server.HandleFunc("/api/v1/user", api.GetUser).Methods("GET", "OPTIONS")         //

	server.HandleFunc("/api/v1/registration", api.Register).Methods("POST", "OPTIONS") //
	server.HandleFunc("/api/v1/login", api.Login).Methods("POST", "OPTIONS")           //
	server.HandleFunc("/api/v1/settings", api.SettingsPost).Methods("POST", "OPTIONS") //

	server.HandleFunc("/api/v1/login", api.Logout).Methods("DELETE", "OPTIONS") //

	server.HandleFunc("/api/v1/settings", api.PhotoUpload).Methods("PUT", "OPTIONS") //

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
