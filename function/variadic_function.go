package main
import "fmt"

func main(){

	bestFinsher := computeBestFinsherInTheWorldChampionish(12,3,1,4,7,54,64,9,65)
	fmt.Println(bestFinsher)
}

// variadic function is a function that accepts unknown number of parameters as slice
func computeBestFinsherInTheWorldChampionish(finishers ...int) int {
	bestFinisher := finishers[0]

	for _,i:= range finishers{
		if i < bestFinisher{
			bestFinisher = i
		}
	}
	return bestFinisher
}