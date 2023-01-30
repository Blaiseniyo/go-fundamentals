package main
import "fmt"

func main(){

	// you can declare the variables normaly outside the if statament

	// peterAge :=13
	// johnAge := 18

	// you can alse declare the variables inside the if statament 
	if peterAge,johnAge := 23,25;  peterAge > johnAge {
		fmt.Println("Peter is older than John")
	}else if johnAge == peterAge{
		fmt.Println("Peter and John are of the same Age")
	}else{
		fmt.Println("John is older than Peter")
	}

}