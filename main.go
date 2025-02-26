package main

import "fmt"

func main() {
	myBill := newBill("ichami")

	// fmt.Println("before updates:", myBill)
	myBill.updateTips(10)
	myBill.addItem("apple", 4.99)
	myBill.addItem("shawarma", 14.99)
	myBill.addItem("eru", 10.99)
	// fmt.Println("after updates:", myBill)

	fmt.Println(myBill.format())
}
