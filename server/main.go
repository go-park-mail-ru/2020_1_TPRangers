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

var FileMaxSize = int64(5 * 1024 * 1024)

type DataHandler struct {
	dataBase   DB.DataInterface
	cookieBase DB.CookieInterface
}

func SetCookie(w *http.ResponseWriter, cookieValue string) {
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(12 * time.Hour),
	}
	http.SetCookie(*w, &cookie)
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
	// структура для реги
	// name
	// email
	// phone
	// date
	// password
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
	var userData AP.JsonStruct
	err := decoder.Decode(&userData)
	mapData := userData.Body.([]interface{})[0].(map[string]interface{})


	if err != nil {
		// посмотреть номер ошибки
		fmt.Print(err,"\n")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.DecodeError})
		return
	}

	login := mapData["login"].(string)
	password := mapData["password"].(string)


	if existFlag := (dh.dataBase).CheckUser(login); existFlag {

		if passFlag := (dh.dataBase).CheckAuth(login, password); passFlag == nil {
			answerData := make(map[string]interface{})
			answerData["isAuth"] = true
			answerData["data"] = (dh.dataBase).GetUserDataLogin(login)

			json.NewEncoder(w).Encode(&AP.JsonStruct{Body: answerData})
			SetCookie(&w, (dh.cookieBase).SetCookie(login))

			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&AP.JsonStruct{Err: ET.WrongPassword})
			return
		}

	} else {
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

func SetCorsMiddleware(r *mux.Router) mux.MiddlewareFunc {

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
	server := mux.NewRouter()
	server.Use(SetCorsMiddleware(server))

	db := DB.NewDataBase()
	cb := DB.NewCookieBase()
	DB.FillDataBase(db)
	fmt.Print("data: ", db.IdMeta , " \n", db.UserId , " \n" )

	api := &(DataHandler{dataBase: db, cookieBase: cb})

	server.HandleFunc("/feed", api.Feed)
	server.HandleFunc("/profile", api.Profile).Methods("GET")
	server.HandleFunc("/register", api.Register).Methods("POST")
	server.HandleFunc("/login", api.Login).Methods("POST")
	server.HandleFunc("/settings", api.SettingsPost).Methods("POST")
	server.HandleFunc("/settings", api.SettingsGet).Methods("GET")
	server.HandleFunc("/settings", api.PhotoUpload).Methods("PUT")
	fmt.Print("hosted at 3001")
	http.ListenAndServe(":3001", server)

}
