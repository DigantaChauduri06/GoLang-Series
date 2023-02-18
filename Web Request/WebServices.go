package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetResponse(url string) string {
	res, err := http.Get(url)
	CheckError(err)
	defer res.Body.Close()
	fmt.Println("Status code ", res.StatusCode, "\nContent Length ", res.ContentLength)
	content, err := io.ReadAll(res.Body)
	CheckError(err)
	// fmt.Println("Content is ", content) // In byte
	// To get it in the format of string we need to format of string
	// fmt.Println(string(content))
	// We can use String builder also
	var response strings.Builder
	len, err := response.Write(content)
	CheckError(err)
	fmt.Println(response.String(), "\n", "And the length is ", len)
	return response.String()
}


func PostRequest(url  string, data io.Reader) {
	res, err := http.Post(url,"application/json", data)
	CheckError(err)
	defer res.Body.Close()
	content, err2 := io.ReadAll(res.Body)
	CheckError(err2)
	fmt.Println("Content After post req ", string(content))
}


func PostForm(url string, data url.Values) {
	res, err := http.PostForm(url,data)
	CheckError(err)
	defer res.Body.Close()
	content,err2 := io.ReadAll(res.Body)
	CheckError(err2)
	fmt.Println(string(content))
}