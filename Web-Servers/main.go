package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	
    // r.HandleFunc("/", )
	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080",r))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome To Golang -- A new way of programming</h1>"))
}
