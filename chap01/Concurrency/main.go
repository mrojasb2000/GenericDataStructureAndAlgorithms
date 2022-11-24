package main

import (
	"fmt"
	"time"
)

func regularFunction() {
	fmt.Println("Just entered regularFunction()")
	time.Sleep(10 * time.Second)
}

func gorotineFunction() {
	fmt.Println("Just entered goroutineFunction()")
	time.Sleep(5 * time.Second)
	fmt.Println("goroutineFunction finished its work")
}

func main() {
	go gorotineFunction()
	fmt.Println("In main one line below goroutineFunction()")
	regularFunction()
	fmt.Println("In main one line below regularFunction()")
}
