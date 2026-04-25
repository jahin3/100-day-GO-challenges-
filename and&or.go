package main 

import ("fmt")

func main() {

	age := 18
	sex:= "male"

	if age == 20 && sex == "male" {
		fmt.Println("You are ready to get married")
	} else if age == 18 || sex == "male" {
		fmt.Println("You are eligible to vote but not ready to get married")	
	} else {
		fmt.Println("You are not eligible to vote and not ready to get married")
	}
}