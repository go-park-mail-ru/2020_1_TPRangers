package main

import (
	DB "./database"
	ET "./errors"
	AP "./json-answers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

var FileMaxSize = int64(5 * 1024 * 1024)

type DataHandler struct {
	dataBase   DB.DataInterface
	cookieBase DB.CookieInterface
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
		Name:     "session_id",
		Value:    cookieValue,
		Expires:  time.Now().Add(12 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(*w, &cookie)
}

func SetData(authFlag bool, answerData DB.MetaData, w *http.ResponseWriter) {

	answer := make(map[string]interface{})
	answer["isAuth"] = authFlag
	answer["data"] = answerData

	json.NewEncoder(*w).Encode(&AP.JsonStruct{Body: answer})
	(*w).WriteHeader(http.StatusOK)

}

func SetErrors(err []string, status int, w *http.ResponseWriter) {
	(*w).WriteHeader(status)
	json.NewEncoder(*w).Encode(&AP.JsonStruct{Err: err})
}

func (dh DataHandler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Print("=============REGISTER=============\n")

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var newUser AP.JsonStruct
	err := decoder.Decode(&newUser)
	mapData, convertionError := getDataFromJson(newUser)

	if err != nil || convertionError != nil {
		fmt.Print(err, " ", convertionError, "\n")
		SetErrors([]string{ET.DecodeError}, http.StatusBadRequest, &w)
		return
	}

	login := mapData["email"].(string)

	data := DB.NewMetaData(mapData["name"].(string), mapData["phone"].(string), mapData["password"].(string), mapData["date"].(string), make([]byte, 0))
	if err, info := (dh.dataBase).AddUser(login, *data); err == nil {

		SetData(true, info, &w)

		cookie := (dh.cookieBase).SetCookie(login)
		SetCookie(&w, cookie)

	} else {
		SetErrors([]string{ET.AlreadyExistError}, http.StatusBadRequest, &w)
		return
	}
}

func (dh DataHandler) Login(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Login=============\n")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var userData AP.JsonStruct
	err := decoder.Decode(&userData)
	mapData, convertionError := getDataFromJson(userData)

	if err != nil || convertionError != nil {
		fmt.Print(err, " ", convertionError, "\n")
		SetErrors([]string{ET.DecodeError}, http.StatusBadRequest, &w)
		return
	}

	login := mapData["login"].(string)
	password := mapData["password"].(string)

	if existFlag := (dh.dataBase).CheckUser(login); existFlag {

		if passFlag := (dh.dataBase).CheckAuth(login, password); passFlag == nil {

			SetData(true, (dh.dataBase).GetUserDataLogin(login), &w)
			SetCookie(&w, (dh.cookieBase).SetCookie(login))

		} else {
			SetErrors([]string{ET.WrongPassword}, http.StatusBadRequest, &w)
			return
		}

	} else {
		SetErrors([]string{ET.DoesntExistError}, http.StatusBadRequest, &w)
		return
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

		SetData(true, dh.dataBase.GetUserDataLogin(login), &w)

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
		err := decoder.Decode(&newMeta)
		mapData, convertionError := getDataFromJson(newMeta)

		if err != nil || convertionError != nil {
			SetErrors([]string{ET.DecodeError}, http.StatusBadRequest, &w)
			return
		}

		newData := *DB.NewMetaData(mapData["name"].(string), mapData["phone"].(string), mapData["password"].(string), mapData["date"].(string), make([]byte, 0))
		newData = DB.MergeData(dh.dataBase.GetUserDataLogin(login), newData)
		dh.dataBase.EditUser(login, newData)

		SetData(true, newData, &w)

	} else {
		SetErrors([]string{ET.WrongCookie}, http.StatusBadRequest, &w)
		return
	}

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

		var freeData DB.MetaData
		SetData(true, freeData, &w)

	} else {
		SetErrors([]string{ET.WrongCookie}, http.StatusBadRequest, &w)
		return

	}
}

func (dh DataHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		fmt.Print(err, "\n")
		SetErrors([]string{ET.CookieExpiredError,}, http.StatusBadRequest, &w)
		return
	}

	var freeData DB.MetaData
	SetData(false, freeData, &w)

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)
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
	fmt.Print("data: ", db.IdMeta, " \n", db.UserId, " \n")

	api := &(DataHandler{dataBase: db, cookieBase: cb})

	server.HandleFunc("/feed", api.Feed)

	server.HandleFunc("/profile", api.Profile).Methods("GET")
	server.HandleFunc("/settings", api.SettingsGet).Methods("GET")

	server.HandleFunc("/register", api.Register).Methods("POST")
	server.HandleFunc("/login", api.Login).Methods("POST")
	server.HandleFunc("/settings", api.SettingsPost).Methods("POST")

	server.HandleFunc("/logout", api.Logout).Methods("DELETE")

	server.HandleFunc("/settings", api.PhotoUpload).Methods("PUT")

	fmt.Print("hosted at 3001")
	http.ListenAndServe(":3001", server)

}
