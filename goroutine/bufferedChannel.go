package main

import "fmt"

func main() {

	mychnl := make(chan string, 5) // buffered channel
	mychnl <- "Hello"
	mychnl <- "Team"
	mychnl <- "Great"
	mychnl <- "work"

	fmt.Println(<- mychnl)
	fmt.Println(<- mychnl)
	fmt.Println("Length of the channel is: ", len(mychnl))
	fmt.Println("Capacity of the channel is: ", cap(mychnl))

	// for msgs := range mychnl {
		
	// 	fmt.Println("Length of the channel is: ", msgs)
	// }
	
}
