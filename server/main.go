package main

import (
	DB "./database"
	ET "./errors"
	AP "./json-answers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

type DataHandler struct {
	dataBase   DB.DataInterface
	cookieBase DB.CookieInterface
}

var FileMaxSize = int64(5 * 1024 * 1024)

func SetCookie(w *http.ResponseWriter, cookieValue string) {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(12 * time.Hour),
	}
	http.SetCookie(w, &cookie)
}

func (dh DataHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Print("=============REGISTER=============\n")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var newUser AP.JsonStruct
	err := decoder.Decode(newUser)
	mapData := newUser.Body.(map[string]interface{})

	if err != nil {
		// посмотреть номер ошибки
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.DecodeError})
		return
	}

	if err, info := (dh.dataBase).AddUser(mapData["login"].(string), mapData["data"].(DB.MetaData)); err == nil {

		answerData := make(map[string]interface{})
		answerData["isAuth"] = true
		answerData["data"] = info

		json.NewEncoder(w).Encode(&AP.JsonStruct{Body: answerData})
		SetCookie(&w, (dh.cookieBase).SetCookie(mapData["login"].(string)))

		w.WriteHeader(http.StatusOK)
	} else {
		// посмотреть номер ошибки
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.AlreadyExistError})
		return
	}
}

func (dh DataHandler) Login(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Login=============\n")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var loginUser AP.JsonStruct
	err := decoder.Decode(loginUser)
	mapData := loginUser.Body.(map[string]interface{})

	if err != nil {
		// посмотреть номер ошибки
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.DecodeError})
		return
	}

	if existFlag := (dh.dataBase).CheckUser(mapData["login"].(string)); existFlag {

		answerData := make(map[string]interface{})
		answerData["isAuth"] = true
		answerData["data"] = (dh.dataBase).GetUserDataLogin(mapData["login"].(string))

		json.NewEncoder(w).Encode(&AP.JsonStruct{Body: answerData})
		SetCookie(&w, (dh.cookieBase).SetCookie(mapData["login"].(string)))

		w.WriteHeader(http.StatusOK)
	} else {
		// посмотреть номер ошибки
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.DoesntExistError})
		return
	}

}

func (dh DataHandler) Profile(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Profile=============\n")

}

func (dh DataHandler) Feed(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Feed=============\n")

}

func (dh DataHandler) SettingsPost(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============SettingsPost=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.CookieExpiredError})
		return
	}

	if login, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {
		answerData := make(map[string]interface{})
		answerData["data"] = dh.dataBase.GetUserDataLogin(login)
		answerData["isAuth"] = true
		json.NewEncoder(w).Encode(&AP.JsonStruct{Body: answerData})

		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.WrongCookie})
		return
	}

}

func (dh DataHandler) SettingsGet(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============SettingsGet=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.CookieExpiredError})
		return
	}

	if login, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var newMeta DB.MetaData
		err := decoder.Decode(newMeta)

		if err != nil {
			// посмотреть номер ошибки
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.DecodeError})
			return
		}

		newMeta = DB.MergeData(dh.dataBase.GetUserDataLogin(login), newMeta)
		dh.dataBase.EditUser(login, newMeta)

		answerData := make(map[string]interface{})
		answerData["data"] = newMeta
		answerData["isAuth"] = true

		json.NewEncoder(w).Encode(&AP.JsonStruct{Body: answerData})
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.WrongCookie})
		return
	}

}

func (dh DataHandler) PhotoUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Print("=============PhotoUpload=============\n")

	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.CookieExpiredError})
		return
	}

	if login, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {

		r.ParseMultipartForm(FileMaxSize)

		file, _, err := r.FormFile("uploadedFile")
		if err != nil {
			w.WriteHeader(http.StatusInsufficientStorage)
			json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.FileError})
			return
		}
		defer file.Close()

		userData := dh.dataBase.GetUserDataLogin(login)
		photoByte, fileErr := ioutil.ReadAll(file)

		if fileErr != nil {
			w.WriteHeader(http.StatusInsufficientStorage)
			json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.FileSavingError})
			return
		}

		userData.Photo = photoByte
		dh.dataBase.EditUser(login, userData)

		w.WriteHeader(http.StatusOK)
		answerData := make(map[string]interface{})
		answerData["isAuth"] = true
		json.NewEncoder(w).Encode(&AP.JsonStruct{Body: answerData})

	} else {

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.WrongCookie})
		return

	}
}

func SetCorsMiddlware(r *mux.Router) mux.MiddlewareFunc {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			(w).Header().Set("Access-Control-Allow-Origin", "*")
			(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			(w).Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization")
			(w).Header().Set("Access-Control-Allow-Credentials", "true")
			(w).Header().Set("Content-Type", "*")
			next.ServeHTTP(w, req)
		})
	}

}

func main() {
	fmt.Print("main")
	server := mux.NewRouter()

	server.Use(SetCorsMiddlware(server))
	db := DB.NewDataBase()
	cb := DB.NewCookieBase()
	DB.FillDataBase(db)
	api := &(DataHandler{dataBase: db, cookieBase: cb})

	server.HandleFunc("/feed", api.Feed)
	server.HandleFunc("/profile", api.Profile).Methods("GET")
	server.HandleFunc("/register", api.Register).Methods("POST")
	server.HandleFunc("/login", api.Login).Methods("POST")
	server.HandleFunc("/settings", api.SettingsPost).Methods("POST")
	server.HandleFunc("/settings", api.SettingsGet).Methods("GET")
	server.HandleFunc("/uploadphoto", api.PhotoUpload).Methods("POST")

	http.ListenAndServe(":3001", server)

}
