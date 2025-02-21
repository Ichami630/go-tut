package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// fmt.Print("prints without new line")
	// fmt.Print("you can see from here \n")
	// fmt.Println("prints with a newline at the end")

	name := "ichami"
	age := 12
	height := 1.3
	isGoFun := true
	pi := 3.14159

	//using fmt.Printf for formatted strings %_ = a format specify (%v,%d,%s etc)
	fmt.Printf("my name is %s, im %d and %.2f meters tall \n", name, age, height)

	//storing formatted text in a variable
	message := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(message)

	//special formatting verbs
	fmt.Printf("print type of name: %T\n", name) //%T prints the type of the variable
	fmt.Printf("boolean value: %t\n", isGoFun)   //%t prints a boolean (true/false)
	fmt.Printf("scientific notation: %e\n", pi)  //%e prints scientific notation
	fmt.Printf("adds quote %q\n", name)          //%q adds quotes around a string

	//printing a struct
	P := Person{name: "ichami", age: 10}
	fmt.Printf("struct values %v\n", P)   //prints struct values
	fmt.Printf("struct details %+v\n", P) //prints fields plus values

	//using backticks for multiline prints
	msg := `Hello ichami
	Welcome to Go programming
	THis string spans multiple lines`
	fmt.Println(msg)
}
