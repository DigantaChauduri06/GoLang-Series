package main

import (
	"fmt"
	"log"
	"time"
)

// [1, 2, 3, 4, 5]


func DelByIndex(arr []int,i int) []int {
	mySlice := []int{}
	mySlice = append(arr[:i], arr[i+1:]...)
	return mySlice
}

func Gen() {
		// arr := []int{1,2,3,4,5}

	// fmt.Println(delByIndex(arr, 3))
	var num = 122
	ReturnedPointer(&num)
	fmt.Println(num)
	fmt.Println(time.Now().Format("01-02-2006 15:04:05 Monday"))
	mySlices := make([]int, 6)
	log.Println(mySlices[2])
	mySlices = append(mySlices, 20)
	log.Println(mySlices)
}

func main () {
	mpp := make(map[string]int)

	mpp["KOLA"] = 1002
	delete(mpp, "KOLA")
	log.Println(mpp)
	val, ok := mpp["KOLA"]
	log.Println(ok, val)
}