package main 

import "fmt"

func main()	{
	a := 18

	switch {
	case a >= 18:
		fmt.Println("You are eligible to vote")
	case a < 18:
		fmt.Println("You are not eligible to vote")
	default:
		fmt.Println("You are just eligible to vote")
	}	
}