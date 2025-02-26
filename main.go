package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	opt, _ := getInput("Choose an option (a - Add a bill s - Save a bill t - add a tip): ", reader)

	//switch in Go
	switch opt {
	case "a":
		name, _ := getInput("Enter item name: ", reader)
		price, _ := getInput("Enter the price for the item: ", reader)
		//now see that the price is pass as "4.99" and is not a number
		//parse float
		p, err := strconv.ParseFloat(price, 64) //convert string to float with the help of the strconv package
		if err != nil {
			fmt.Println("The price must be a number")
		}
		//now add item to items map
		b.addItem(name, p)
		fmt.Println(name, price)
		promptOption(b)
	case "s":
		fmt.Println("you choosed to save the bill:", b)
	case "t":
		tip, _ := getInput("Enter the tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64) //convert string to float with the help of the strconv package
		if err != nil {
			fmt.Println("The tip must be a number")
		}
		//now add new tip
		b.updateTips(t)
		fmt.Println(t)
		promptOption(b)
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
