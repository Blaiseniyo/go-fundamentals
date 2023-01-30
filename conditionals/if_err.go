package main

import (
	"fmt"
	"os"
)


func main(){

	_,err := os.Open("./text1.txt")

	if err != nil{
		fmt.Println("returned error code - ",err)
	}else{
		fmt.Println("returned error code - ",err)
	}
}