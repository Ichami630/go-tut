package main

/*Pointers in Go
A pointer is a variable that stores the memory location of anothe variable
Pointers help us modify variables outside the variable scope and efficiently manage memory
Use an *(asterisk) to declare a pointer and &(ampersand) to get the address of the variable
*/

import "fmt"

func updateName(x *string) { //creates a pointer variable of type string
	*x = "brandon" //dereferencing the pointer to modify the original value
	/*Now what does this mean?
	Instead of changing x itself, which is just an address,we go to that memory location
	We modify the actuall value which is stored at that memory location
	*/
}

func changeNum(x *int) {
	*x = 100
}

func main() {
	name := "ichami" //declare normal variables
	num := 10

	//declaring a pointer to an integer
	var ptr *int

	ptr = &num //assigning the address of num to the ptr
	changeNum(ptr)

	// updateName(name)

	// fmt.Println("memory address of name:", &name)

	m := &name //pointer m that stores the memory address of the variable name
	// fmt.Println("memory address:", m)

	// fmt.Println("value stored at pointer memory address m:", m, "is the original value of name:", *m) //the * is used to derefference a pointer
	//Dereferencing a pointer means goining into to the memory address of that pointer and getting the value which the pointer points to
	fmt.Println(name) //original value of name remains unchanged
	updateName(m)
	fmt.Println("now the original value of the var name changes to the updated name as:", name)
}
