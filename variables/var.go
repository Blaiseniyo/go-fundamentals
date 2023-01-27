package main

import(
	"fmt"
	"reflect"
	"strconv"
)

var (
	name,city string
	age,dob int
	number = "5"
)

func main(){
	fmt.Println("My name is:",name,"and I stay at ", city)
	fmt.Println("I was born in", dob, "and so I am ",age,"old")

	fmt.Println("Type of name variable is:",reflect.TypeOf(name))
	fmt.Println("Type of name dob is", reflect.TypeOf(dob))
	
	iNumber,err := strconv.Atoi(number)

	if err == nil{
		total := iNumber + age
		fmt.Println("Total is: ", total)
	}
	
}