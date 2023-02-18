package main

import (
	"fmt"
	"os"
)

func WriteItJSON(URL string) {
	fs, err := os.Create("LCO.json")
	CheckError(err)
	defer fs.Close()
	content := GetResponse(URL)
	_, err2 := fs.WriteString(content)
	CheckError(err2)
	fmt.Println("Write Completed !")
}