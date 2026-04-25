package main

import "fmt"

func main() {

    a:= 18

    if a > 18 {
		fmt.Println("You are eligible to vote")
	} else if a < 18 {
		fmt.Println("You are not eligible to vote")
	} else {
		fmt.Println("You are just eligible to vote")
	} 
}