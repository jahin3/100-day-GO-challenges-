package main

import "fmt"

func add(a, b int) {

	fmt.Println(a + b)
}

func subtract(a, b int) {

	fmt.Println(a - b)
}

func babluCat(name string) {

	fmt.Println(name)
}
func main() {

	add(4, 7)
	subtract(32, 16)
	babluCat("Bablu")
}
