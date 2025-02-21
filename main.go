package main

import "fmt"

var name = "ichami"      //type inference
var city string = "buea" //explicit type declaration
var occupation string

func main() {
	country := "cameroon" //short declaration, only allowed inside functions
	fmt.Println("hello, go, ichami is coming for you!")
	fmt.Println("Country:", country)
}
