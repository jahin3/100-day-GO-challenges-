package main 
import "fmt"

func printWelcomeMessage()  {
    fmt.Println("Welcome to the application!")
}

func getUserName() string {
    var name string
    fmt.Println("Please enter your name => ")
    fmt.Scanln(&name)
    return name
}

func getTwoNumbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Please enter first number => ")
	fmt.Scanln(&num1)
	fmt.Println("Please enter second number => ")
	fmt.Scanln(&num2)
	return num1, num2
}

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func display(name string, sum int) {

	fmt.Println("Hello " , name)
	fmt.Println("Summation => ", sum)

}

func printGoodbyeMessage() {
	fmt.Println("Thank you for using the application. Goodbye!")
	fmt.Println("Have a great day!")
}

func main() {
    printWelcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum := add(num1, num2)
	display(name, sum)
	printGoodbyeMessage()	
}