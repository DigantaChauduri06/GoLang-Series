package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	urls := []string{
		"https://github.com/DigantaChauduri06",
		"https://www.netflix.com",
		"https://chat.openai.com/chat",
		"https://nodejs.org/asfsdf",
	}
	// wg.Add(len(urls)+1)
	for _, url := range urls {
		wg.Add(1)
		// go CheckPing(url)
		go func (url string)  {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer wg.Done()
			defer res.Body.Close()
			log.Printf("%s ping = %d \n", url,res.StatusCode)

		}(url)
	}
	wg.Wait()
	fmt.Println("I Am Here")
}

func CheckPing(url string) (int) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer wg.Done()
	defer res.Body.Close()
	return res.StatusCode
}

// func example() int {
//     defer fmt.Println("defer 2")
//     defer fmt.Println("defer 1")
//     fmt.Println("in example")
//     return 42
// }

// func main() {
//     result := example()
//     fmt.Println("result:", result)
// }
