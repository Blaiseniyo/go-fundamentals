package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	// switch cases has an implicit break statement added to them, it automatically stops the codes execution to the next statment outside the switch statement
	
	switch tmpNum := random(); tmpNum {

	case 2,4,6,8:
		fmt.Println("The returned value is an even number - ",tmpNum)
	case 1,3,5,7,9:
		fmt.Println("The returned value is an odd number - ", tmpNum)
		// To check the next case statement add fallthrough statement
		 fallthrough
	default:
		fmt.Println("Hit the default Block by use of fallthrough statement ")
		
	}
}

func random() int{
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}