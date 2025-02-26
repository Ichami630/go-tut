package main

/*Structs and custom types in Go
Go doesn't have classes, but structs allow you to create custom data types that can group related fields together
A struct is a collection of related fields

-Struct with pointers
*By default, Go passes copy of struct to functions. To modify a struct, use pointers
*/
import "fmt"

// define a struct
type Person struct {
	name string
	age  int
}

type Car struct {
	brand string
	model string
	year  int
}

type Book struct {
	title  string
	author string
}

// method for car struct
/*A method is tied to a specific struct type and can access it fields without exlicitly passing them
Receiver parameter (c Car)
c is an instance of Car pass implicitly
*/
func (c Car) displayInfo() {
	fmt.Printf("Car: %s %s (%d)\n", c.brand, c.model, c.year)
}

// method using pointer receiver
func (b *Book) updateTitle(newTitle string) {
	b.title = newTitle
}

// custom types
// you can define a custom type using type
type Age int

func main() {
	//create a struct instance
	p1 := Person{name: "Ichami", age: 10}
	//access fields
	fmt.Println("Name:", p1.name, "\nAge:", p1.age)

	//struct methods
	c1 := Car{brand: "Benz", model: "ML350", year: 2025}
	c1.displayInfo()

	b1 := Book{title: "Golang", author: "ichami"}
	fmt.Println("before:", b1)
	b1.updateTitle("Golang new edition")
	fmt.Println("After:", b1)

	var myAge Age = 10
	fmt.Println(myAge)
}
