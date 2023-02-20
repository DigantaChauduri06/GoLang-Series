package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// -- Models

type Course struct {
	CourseId string `json:"cid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	LastName string `json:"lastname"`
}

// fake db
var Courses []Course

// -- Middlewares
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080",r))
}


// -- Controllers
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to LCO 😊😊😊</h1>"))
}


func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get The Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}