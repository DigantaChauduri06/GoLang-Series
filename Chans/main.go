package main

import (
	"fmt"

	"github.com/DigantaChauduri06/helpers"
)

func CalculateValue(intchan chan int) {
	randomValue := helpers.RandomValue()
	intchan <- randomValue
}

func main() {
	intchan := make(chan int)
	defer close(intchan)
	go CalculateValue(intchan)
	num := <-intchan
	fmt.Println(num)
}