package main

import (
	"fmt"
	// "time"
)

func sayHello() {
	fmt.Println("Hello, world")
}

func main() {
	go sayHello() // start a new goroutine

	fmt.Println("Main function")
	// time.Sleep()
}
