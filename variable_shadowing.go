package main

import "fmt"

var a = 10

func main() {

	age := 30

	if age > 18 {
		a := 47
		fmt.Println(a) // This will print 47
	}

	fmt.Println(a) // This will print 10
}
