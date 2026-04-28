package main 

import (
	
	"fmt"
 
    "example.com/mathlib"

)


var (
	a := 20
	b := 30
)

func main() {

	fmt.Println("Showing custom package")	
	
	mathlib.Add(4 , 7)
	fmt.Println(mathlib.Money)

}	