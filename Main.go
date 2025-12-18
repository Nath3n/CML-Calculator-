package main

//What you build

//A terminal calculator that supports: trying

//add, subtract, multiply, divide done

//Takes input via command-line args or stdin

//You’ll learn

//os.Args

//Functions

//Error handling

//Why recruiters like it

//Shows basic Go fluency + clean logic
import "fmt"

func main() {
	a := 0
	b := 0
	operation := ""
	fmt.Print("Hello welcome to the CML calculator!\n")

	fmt.Print("select two numbers please enter number 1:")
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println("Please enter a valid number")
		return
	}
	fmt.Print("select two numbers please enter number 2:")
	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println("Please enter a valid number")
		return
	}
	fmt.Print("Alright what operation do you want to do with this? (- / + % *)")
	if _, err := fmt.Scanln(&operation); err != nil {
		fmt.Println("Enter a valid operation")
		return
	}

	switch operation {
	case "-":
		fmt.Println(a, " - ", b, " = ", a-b)
	case "+":
		fmt.Println(a, " + ", b, " = ", a+b)
	case "*":
		fmt.Println(a, " * ", b, " = ", a*b)
	case "/":
		if b == 0 {
			fmt.Println("error: division by zero is not supported")

			return
		}
		fmt.Println("Result", a/b)

	case "%":
		fmt.Println(a, " % ", b, " = ", a%b)
	default:
		fmt.Println("invalid operation")
	}
}
