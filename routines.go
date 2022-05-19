package main

import (
	"fmt"
)

var int p  int y 

func main() {
	//for i = 0; i <= 0; i++ {
	go func() { //A
		p := 1
		fmt.Println("y", y)
	}()

	go func() { //B
		y := 1
		fmt.Println("p", p)
	}()
}
