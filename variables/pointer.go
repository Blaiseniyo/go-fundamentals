package main

import(
	"fmt"

)

func main(){
	name := "Peter"
	course := "Getting started with Kubernrtes"

	fmt.Println("Hi",name,"Your course is ",course)
	updateCourse(&course) 
	fmt.Println("Hi",name,"Your course is ",course)
	
}

func updateCourse(course *string) *string{
	*course = "Getting started with Docker"
	return course
}