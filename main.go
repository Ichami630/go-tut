package main

import "fmt"

func main() {
	myBill := newBill("ichami")

	fmt.Println(myBill.format())
}
