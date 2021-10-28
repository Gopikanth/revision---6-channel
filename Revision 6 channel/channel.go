package main

import "fmt"

func first(a chan bool) {
	fmt.Printf("first function\n")
	a <- true //here we are sending data to a

}
func main() {
	a := make(chan bool)
	go first(a)
	<-a //here we are receving the data from a channel
	fmt.Println("main function")
}
