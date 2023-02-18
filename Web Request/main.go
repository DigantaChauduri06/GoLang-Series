package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const URL = "http://localhost:8000/postform"

type User struct {
	Name string `json:"userName"`
	Email string
	Password string `json:"-"`
	IsValid bool
	Courses []string
}

func main() {
	users := []User{
		{"Diganta", "diganta.12001517@lpu.in", "123abc", true, []string{"Reactjs","Nodejs", "Golang"}},
		{"Kajal", "kajal.322251@lpu.in", "456yzu", true, []string{"Soft skills","Mongodb", "Python"}},
		{"Khusi", "dontknow.255644@lpu.in", "KMP211", true, []string{"Reactjs","Nodejs", "Golang"}},
		{"IAMMAD", "meMad.12001517@lpu.in", "211@ddd", true, []string{"Angular","FastAPI", "C++"}},
	}
	// encode, err := json.Marshal(users)
	encode, err := json.MarshalIndent(users, "", "\t")
	CheckError(err)

	log.Printf("%s",encode)
	JSON := []byte(encode)
	isValid := json.Valid(JSON)
	var myUser []User
	if isValid {
		// log.Println("Yes it is valid")
		err := json.Unmarshal(JSON,&myUser)
		CheckError(err)
		log.Printf("%#v\n",myUser)

	} else {
		log.Println("JSON is not valid")
	}


}


func PostFormReq() {
	data := url.Values{}
	data.Add("name", "Diganta")
	data.Add("email", "diganta.12001517@lpu.in")
	data.Add("password", "diganta@06")
	data.Add("isVarified", "true")
	PostForm(URL,data)
}

func POSTREQ() {
	data := strings.NewReader(`
		{
			"Courses" : "ReactJS",
			"Price" : "$99",
			"courseURL" : "https://www.udemy.com"
		}
	`)
	PostRequest(URL, data)
}

func GETREQ() {
	res, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	databytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// content := string(databytes)
	// fmt.Println(content)
	var resBuilder strings.Builder
	byteCount, _  := resBuilder.Write(databytes)
	fmt.Println(byteCount)
	fmt.Println(resBuilder.String())
}