package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("./someFile.txt")
	defer os.Remove("./someFile.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	text := "Something to write for today"
	length, x := io.WriteString(file,text)
	if x != nil {
		panic(x)
	}
	fmt.Println("Length is ", length)
	AppendFile(file, "Something been appended")
	ReadFile("./someFile.txt")

}

func ReadFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func AppendFile(filename io.Writer, appendedstring string) {
	_, err := io.WriteString(filename,appendedstring)
	if err != nil {
		panic(err)
	}
}