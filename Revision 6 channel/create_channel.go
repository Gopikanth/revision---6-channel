package main

import "fmt"

func main() {
	var a chan int //one way of creating channel
	if a == nil {
		fmt.Println("This is Easy")
	}
	a = make(chan int) //other eay of creating channel
	fmt.Printf("This is too easy")
}
