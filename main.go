package main

import "fmt"

func multiply(num1 float64, num2 float64) float64 {
    return num1 * num2
}

func main() {
    result1 := multiply(5.0, 3.0	)
	result2 := multiply(result1, 2.0)	
	result3 := multiply(32.3, 4.5)

	
     fmt.Println(result3)
	 fmt.Println(result1)
	 fmt.Println(result2)
}
