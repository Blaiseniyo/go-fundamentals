package main

import(
	"fmt"
	"reflect"
	"strconv"
)

// Global declaration of variables

// var (
// 	name,city string
// 	age,dob int
// 	number = "5"
// )

func main(){

	// Local declaration of varaible
	name := "Peter"
	city := "Kigali" 
	age  := 45
	dob  := 1878
	number := "5"
	fmt.Println("My name is:",name,"and I stay at ", city)
	fmt.Println("I was born in", dob, "and so I am ",age,"old")

	// Check variable Type
	fmt.Println("Type of name variable is:",reflect.TypeOf(name))
	fmt.Println("Type of name dob is", reflect.TypeOf(dob))
	
	// change string to int
	iNumber,err := strconv.Atoi(number)

	if err == nil{
		total := iNumber + age
		fmt.Println("Total is: ", total)
	}
	
}