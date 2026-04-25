package main 

import ("fmt")

func main() {

	isPretty := false

    if isPretty {
		fmt.Println("This is not pretty.")
	} else if !isPretty {
		fmt.Println("This is pretty.")
	} else {			
		fmt.Println("This is neither pretty nor not pretty.")
	}	

}