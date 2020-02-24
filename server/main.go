package main

import (
	myDataB "./database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type DataHandler struct {
	dataBase *myDataB.DataBase
	//cookie
}




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
	data := myDataB.NewMetaData("xd", "xd", "xd", make([]byte, 2))
	login := "nikita"
	if err, info := dh.dataBase.AddUser(login, data); err != nil {
		http.Error(w,`{"error":"неправильные данные!"}` , 401)
		return
	} else {
		json.NewEncoder(w).Encode(&myDataB.Result{Body: info})
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

func main() {
	fmt.Print("main")
	server := mux.NewRouter()
	db := myDataB.NewDataBase()
	api := &(DataHandler{dataBase:&db})

	server.HandleFunc("/feed", api.Feed)
	server.HandleFunc("/profile", api.Profile)
	server.HandleFunc("/register",api.Register)
	server.HandleFunc("/login",api.Login)
	server.HandleFunc("/settings", api.Settings)

	http.ListenAndServe(":3001", server)

}
