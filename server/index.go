package main


import (
	// "encoding/json"
	"fmt"
	// "log"
	"net/http"
	// "sync"
)




func makeCorsHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Content-Type", "application/json")
}

func main() {
	fmt.Println("main")

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("login")
		makeCorsHeaders(&w)
		w.Write([]byte("{}"))
	})

	http.ListenAndServe(":3001", nil)
}