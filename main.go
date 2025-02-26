package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// get input
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, err
}

// function to dynamically create a new bill
func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new bill: ")
	// name, _ := reader.ReadString('\n') //store the entered value once the user presses enter
	// name = strings.TrimSpace(name)     //removes extra spaces
	name, _ := getInput("Create a new bill: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill - ", name)

	return b
}

// func to allow users choose an option
func promptOption(b Bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose an option (a - Add a bill b - Save a bill c - add a tip): ", reader)

	//switch in Go
	switch opt {
	case "a":
		name, _ := getInput("Enter item name: ", reader)
		price, _ := getInput("Enter the price for the item: ", reader)
		fmt.Println(name, price)
	case "b":
		fmt.Println("you choosed b")
	case "c":
		tip, _ := getInput("Enter the tip amount ($): ", reader)
		fmt.Println(tip)
	default:
		fmt.Println("invalid option...please try again")
		promptOption(b)
	}
}

func main() {
	myBill := createBill()
	promptOption(myBill)
	fmt.Println(myBill)

}
