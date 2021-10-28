package main

import "fmt"

func going(ab chan int){
	for i:=0 ; i<8 ; i++{
      ab <- i
	}
	close(ab)
}

func main() {
	a := make(chan int)
	go going(a)
	for {
		for i := range a{
			fmt.Println(i)
		}
	}

}
