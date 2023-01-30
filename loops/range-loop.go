package main

import "fmt"

func main(){
	districts :=[]string{"Kicukiro","Gasabo","Nyarugenge","Kamonyi","Rwamagana"}

	provinceDistricts := []string{"Kamonyi","Rwamagana"}

	for _,i:= range districts{
		for _,j := range provinceDistricts{
			if i == j{
				fmt.Println(i," is a district outside Kigali")
			}
		}
	}
}