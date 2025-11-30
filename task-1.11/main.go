package main

import (
	"fmt"
	"time"
)

// var wg sync.WaitGroup

// func worker(id int) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d начал работу\n", id)
// }

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "один"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "два"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Получено:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Получено:", msg2)
		}
	}
}
