package main

import (
	"fmt"
	"strings"
)


func main(){

	author := "Peter Paul"
	course := "getting started with docker"

	fmt.Println(author,course)
	fmt.Println(conveter(author,course))
}


func conveter(author,course string) (string, string){
	author = strings.ToUpper(author)
	course = strings.Title(course)

	return author,course
}