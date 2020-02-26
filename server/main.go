package main

import (
	DataBase "./database"
	AP "./json-answers"
	ET "./errors"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"

	// "time"
)

type DataHandler struct {
	dataBase   DataBase.DataInterface
	cookieBase DataBase.CookieInterface
}

type JsonStruct struct {
	Body interface{} `json:"body,omitempty"`
	Err  []string    `json:"err,omitempty"`
}

func getDataFromJson(userData JsonStruct) (data map[string]interface{}, errConvert error) {

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

func makeCorsHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Content-Type", "application/json")
}

func (dh DataHandler) Register(w http.ResponseWriter, r *http.Request) {

	// тут получение данных с сервера
	fmt.Print("=============REGISTER=============\n")
	makeCorsHeaders(&w)

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var userData JsonStruct
	err := decoder.Decode(&userData)
	mapData, convertionError := getDataFromJson(userData)

	if err != nil || convertionError != nil {
		//SetErrors([]string{ET.DecodeError}, http.StatusBadRequest, &w)
		return
	}

	login := mapData["email"].(string)
	password := mapData["password"].(string)


	if !dh.dataBase.CheckUser(login) || password != dh.dataBase.GetPasswordByLogin(login){
		fmt.Println("Doesn't exit")
		return
	}



	fmt.Println(login)
	fmt.Println(password)

	json.NewEncoder(w).Encode(&JsonStruct{Body: "Authorised"})
	(w).WriteHeader(http.StatusOK)

}

func (dh DataHandler) Login(w http.ResponseWriter, r *http.Request) {

	fmt.Print("=============Login=============\n")
	makeCorsHeaders(&w)

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var userData JsonStruct
	decoder.Decode(&userData)
	mapData, convertionError := getDataFromJson(userData)

	if  convertionError != nil {
		 SetErrors([]string{ET.DecodeError}, http.StatusBadRequest, &w)
		return
	}

	login := mapData["login"].(string)
	password := mapData["password"].(string)


	if !dh.dataBase.CheckUser(login) || password != dh.dataBase.GetPasswordByLogin(login){
		fmt.Println("Doesn't exit")
		return
	}



	fmt.Println(login)
	fmt.Println(password)

	json.NewEncoder(w).Encode(&JsonStruct{Body: "Authorised"})
	(w).WriteHeader(http.StatusOK)



}

func main() {
	fmt.Print("main")
	server := mux.NewRouter()
	db := DataBase.NewDataBase()
	cb := DataBase.NewCookieBase()

	api := &(DataHandler{dataBase: db, cookieBase: cb})
	DataBase.FillDataBase(db)

	server.HandleFunc("/registration", api.Register)
	server.HandleFunc("/login", api.Login)

	http.ListenAndServe(":3001", server)

}
