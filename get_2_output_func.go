package main 
import "fmt"

func getNumbers(num1 int, num2 int) (int, int) {
	 
	sum:= num1 + num2

	multiply:= num1 * num2
	
	return sum, multiply	 

}
func main() {

	a:= 10
	b:= 20

	p , q := getNumbers(a, b)

	fmt.Println("Sum:", p)
	fmt.Println("Multiply:", q)

}	
