package main

import (
	DataBase "./database"
	ET "./errors"
	AP "./json-answers"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"

	// "time"
)

var FileMaxSize = int64(5 * 1024 * 1024)

type DataHandler struct {
	dataBase   DataBase.DataInterface
	cookieBase DataBase.CookieInterface
}

func getDataFromJson(userData AP.JsonStruct) (data map[string]interface{}, errConvert error) {

	defer func() {

		if err := recover(); err != nil {
			data = make(map[string]interface{})
			errConvert = errors.New("decode err")
		}

	}()

	return userData.Body.([]interface{})[0].(map[string]interface{}), nil
}

func SetCookie(w *http.ResponseWriter, cookieValue string) {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(12 * time.Hour),
	}
	http.SetCookie(*w, &cookie)
}

func SetData(data []interface{}, jsonType []string, w *http.ResponseWriter) {

	answer := make(map[string]interface{})

	for i, val := range jsonType {
		switch val {

		case "isAuth":
			answer[val] = data[i].(bool)
		case "userData":
			answer[val] = data[i].(DataBase.MetaData)

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

	// тут получение данных с сервера
	fmt.Print("=============REGISTER=============\n")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var userData AP.JsonStruct
	decoder.Decode(&userData)
	mapData, convertionError := getDataFromJson(userData)

	if convertionError != nil {
		return
	}

	login := mapData["email"].(string)

	data := DataBase.NewMetaData(mapData["name"].(string), mapData["phone"].(string), mapData["password"].(string), mapData["date"].(string), make([]byte, 0))
	if err, info := (dh.dataBase).AddUser(login, *data); err == nil {

		sendData := make([]interface{}, 2)

		sendData[0] = true
		sendData[1] = info

		SetData(sendData, []string{"isAuth", "userData"}, &w)

		cook := (dh.cookieBase).SetCookie(login)
		cookie := http.Cookie{
			Name:     "session_id",
			Value:    cook,
			Expires:  time.Now().Add(12 * time.Hour),
		}

		fmt.Println(cook)

		http.SetCookie(w, &cookie)

	} else {
		SetErrors([]string{ET.AlreadyExistError}, http.StatusBadRequest, &w)
		return
	}

}

func (dh DataHandler) Login(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Login=============\n")
	fmt.Println(*r)
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var userData AP.JsonStruct
	decoder.Decode(&userData)
	mapData, convertionError := getDataFromJson(userData)

	if convertionError != nil {
		return
	}

	login := mapData["login"].(string)
	password := mapData["password"].(string)

	if !dh.dataBase.CheckUser(login) || password != dh.dataBase.GetPasswordByLogin(login) {
		fmt.Println("Doesn't exit")
		return
	}

	fmt.Println(login)
	fmt.Println(password)

	json.NewEncoder(w).Encode(&AP.JsonStruct{Body: "Authorised"})
	(w).WriteHeader(http.StatusOK)

}

func (dh DataHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		SetErrors([]string{ET.CookieExpiredError,}, http.StatusBadRequest, &w)
		return
	}

	(w).WriteHeader(http.StatusOK)
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)
}

func (dh DataHandler) PhotoUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Print("=============PhotoUpload=============\n")

	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		SetErrors([]string{ET.CookieExpiredError}, http.StatusBadRequest, &w)
		return
	}

	if login, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {

		r.ParseMultipartForm(FileMaxSize)

		file, _, err := r.FormFile("uploadedFile")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			SetErrors([]string{ET.FileError}, http.StatusBadRequest, &w)
			return
		}
		defer file.Close()

		userData := dh.dataBase.GetUserDataLogin(login)
		photoByte, fileErr := ioutil.ReadAll(file)

		if fileErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			SetErrors([]string{ET.FileSavingError}, http.StatusBadRequest, &w)
			return
		}

		userData.Photo = photoByte
		dh.dataBase.EditUser(login, userData)

		sendData := make([]interface{}, 2)

		sendData[0] = true
		sendData[1] = DataBase.MetaData{}

		SetData(sendData, []string{"isAuth", "userData"}, &w)

	} else {
		SetErrors([]string{ET.WrongCookie}, http.StatusBadRequest, &w)
		return

	}
}

func SetCorsMiddleware(r *mux.Router) mux.MiddlewareFunc {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, PUT, DELETE")
			(w).Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
			(w).Header().Set("Access-Control-Allow-Credentials", "true")
			(w).Header().Set("Content-Type", "*")

			next.ServeHTTP(w, req)
		})
	}

}

func (dh DataHandler) Profile(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Profile=============\n")

}

func (dh DataHandler) Feed(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Feed=============\n")

}

func (dh DataHandler) SettingsGet(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============SettingsGET=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		SetErrors([]string{ET.CookieExpiredError,}, http.StatusBadRequest, &w)
		return
	}

	if login, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {

		sendData := make([]interface{}, 2)

		sendData[0] = true
		sendData[1] = (dh.dataBase).GetUserDataLogin(login)

		SetData(sendData, []string{"isAuth", "userData"}, &w)

	} else {
		SetErrors([]string{ET.WrongCookie}, http.StatusBadRequest, &w)
		return
	}

}

func (dh DataHandler) SettingsPost(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============SettingsPOST=============\n")
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		w.WriteHeader(http.StatusBadRequest)
		SetErrors([]string{ET.CookieExpiredError}, http.StatusBadRequest, &w)
		return
	}

	if login, flag := dh.cookieBase.GetUser(cookie.Value); flag == nil {

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		var newMeta AP.JsonStruct
		decoder.Decode(&newMeta)
		mapData, convertionError := getDataFromJson(newMeta)

		if convertionError != nil {
			return
		}

		newData := *DataBase.NewMetaData(mapData["name"].(string), mapData["phone"].(string), mapData["password"].(string), mapData["date"].(string), make([]byte, 0))
		newData = DataBase.MergeData(dh.dataBase.GetUserDataLogin(login), newData)
		dh.dataBase.EditUser(login, newData)

		sendData := make([]interface{}, 2)

		sendData[0] = true
		sendData[1] = newData

		SetData(sendData, []string{"isAuth", "userData"}, &w)

	} else {
		SetErrors([]string{ET.WrongCookie}, http.StatusBadRequest, &w)
		return
	}

}

func main() {
	fmt.Print("main")
	server := mux.NewRouter()
	db := DataBase.NewDataBase()
	cb := DataBase.NewCookieBase()

	server.Use(SetCorsMiddleware(server))

	api := &(DataHandler{dataBase: db, cookieBase: cb})
	DataBase.FillDataBase(db)

	server.HandleFunc("/feed", api.Feed)

	server.HandleFunc("/profile", api.Profile).Methods("GET", "OPTIONS")
	server.HandleFunc("/settings", api.SettingsGet).Methods("GET", "OPTIONS")

	server.HandleFunc("/registration", api.Register).Methods("POST", "OPTIONS")
	server.HandleFunc("/login", api.Login).Methods("POST", "OPTIONS")
	server.HandleFunc("/settings", api.SettingsPost).Methods("POST", "OPTIONS")

	server.HandleFunc("/login", api.Logout).Methods("DELETE", "OPTIONS")

	server.HandleFunc("/settings", api.PhotoUpload).Methods("PUT", "OPTIONS")

	http.ListenAndServe(":3001", server)

}
