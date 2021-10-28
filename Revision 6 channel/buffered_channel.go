package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int, 2) //buffered channel
	go option(a)
	time.Sleep(3 * time.Second)
	for v := range a {
		fmt.Println(v)
		time.Sleep(2 * time.Second)

	}
}
func option(a chan int) {
	for i := 0; i < 8; i++ {
		a <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(a)
}
