package main

import "log"

// func main() {
// 	start := time.Now()
// 	done := make(chan bool)
// 	defer close(done)
// 	defer func ()  {
// 		fmt.Println(time.Since(start))
// 	}()

// 	badNinjas := []string{"Andy", "Johnny", "Kutty", "Tutuy"}

// 	for _, ninja := range badNinjas {
// 		go attack(ninja, done)
// 		done <- false
// 	}
// 	// time.Sleep(time.Second*2)
// 	// Better way to do this use chanels
// 	fmt.Println(<-done)
// }

// func attack(ninja string, done chan bool) {
// 	fmt.Printf("Attacking on %v with stars \n", ninja)
// 	time.Sleep(time.Second)
// 	done <- true
// }

func main() {
	chanel := make(chan string)
	go func () {
		chanel <- "First Message"
	}()
	log.Println(<-chanel)
}