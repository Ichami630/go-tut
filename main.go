package main

/*Go is a pass-by-value language
This means that when we pass a variable to a function, Go creates a copy of the value and pasess that copy
The original value remains unchanged unlesss you use pointers
*/

import "fmt"

func updateName(x string) {
	x = "brandon" //this modifies the copy
}

// one way is to return a value in the function and set the value  of the variable to be that of the  returned function
func updateAge(x int) int {
	x = 12
	return x
}

func updateMenu(x map[string]float64) {
	x["coffee"] = 10.99
}
func main() {
	//Group A type: int,float,string,arrays,bool,struct
	name := "ichami"
	age := 10

	updateName(name) //passing name copy
	age = updateAge(age)

	fmt.Println(name) //name remains unchanged
	fmt.Println(age)  //age will be the updated value in the function

	//Group B types: slices,maps,functions
	menu := map[string]float64{
		"pie":       4.99,
		"ice-cream": 9.99,
	}

	updateMenu(menu)
	fmt.Println(menu) //see that coffee will be added to the menu from the updateMenu func
	/*why was the update possible for group B type?
	In Go, maps are referenced types, unlike struct,int,etc.
	This means that when you pass a map (group A type) to a function, you are passing a reference to the underlying data, not a copy of the map
	*/
}
