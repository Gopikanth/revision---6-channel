package main

import (
	"fmt"
	"time"
)

func other(b chan bool) {
	fmt.Println("start")
	time.Sleep(12 * time.Second)
	fmt.Println("Middle")
	b <- true

}

func main() {
	b := make(chan bool)
	fmt.Println("process started")
	go other(b)
	<-b
	fmt.Println("Process done")
}
