//when there is no go-routines to receive or sent this condition may leads to deadlock
package main

func main() {
	a := make(chan int)
	a <- 5
}

//the above program may lead to deadlock because there is no receiving go routine
