package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("routine started ")
	go first()
	go second()
	time.Sleep((2000 * time.Millisecond))
	fmt.Println("routine terminated")

}
func first() {
	for i := 0; i < 10; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Println("Publication")
	}

}
func second() {
	for j := 0; j < 10; j++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("mana")

	}
}
