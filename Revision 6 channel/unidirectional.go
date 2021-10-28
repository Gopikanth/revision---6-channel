package main

import "fmt"

func other(a chan<- int) {
	a <- 4
}

func main() {
	a := make(chan int)
	go other(a)
	fmt.Println(<-a)

}
