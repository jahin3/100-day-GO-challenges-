package main 

import "fmt"

func division(num1 int , num2 int) int {
	return num1 / num2
}		

func main() {
	result1 := division(10, 2)
	result2 := division(result1, 5)
	result3 := division(100, 4)
	
	
    fmt.Println(result3)
    fmt.Println(result1)
	fmt.Println(result2)
}	