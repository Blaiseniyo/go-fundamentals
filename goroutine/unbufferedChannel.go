package main

import (
	"fmt"
	"time"
)

func getData( ch chan string) {
	msg := <-ch
	fmt.Println(msg)
}

func main() {

	ch := make(chan string)

	go getData(ch)

	ch <- "Hello"  // blocking
	time.Sleep(3)
}