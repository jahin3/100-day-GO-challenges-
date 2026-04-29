package main

import (
	"fmt"
)	

var (
	a = 20
	b = 30
)	

func add(x int, y int)  {
     res:= x + y
	printNum(res)
}

func main() {		
   add(a, b)

}

func printNum(num int) {
	fmt.Println("The result is ", num)
}