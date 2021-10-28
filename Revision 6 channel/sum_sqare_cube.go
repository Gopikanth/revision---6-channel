package main

import "fmt"

func square(number int, squaree chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squaree <- sum

}
func cube(number int, cubeee chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeee <- sum

}
func main() {
	number := 5
	s := make(chan int)
	c := make(chan int)
	go square(number, s)
	go cube(number, c)
	squar, cubee := <-s, <-c
	fmt.Println(squar + cubee)
}
