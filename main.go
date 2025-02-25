package main

import "fmt"

/*Maps in Go
A map is a built-in data type used to store key value pairs,similar to dictonaries in python and hash tables in other languages
*/

func main() {
	//create an empty map using make
	studenAges := make(map[string]int)
	//add key value pairs
	studenAges["ichami"] = 10
	studenAges["feli"] = 4
	fmt.Println(studenAges)

	//declare and initialise a map using the map keyword
	salaries := map[string]float64{
		"ichami": 2500000.00,
		"feli":   2000000.00,
	}
	//add new element to a map
	salaries["betty"] = 50000.00

	//updating an existing map element
	salaries["betty"] = 100000.00

	fmt.Println(salaries)

	//accessing values in map
	fmt.Println(studenAges["ichami"]) //output 10

	fmt.Println(studenAges["john"]) //accessing a non-existing element in map returns 0

	//checking if a key exist
	age, exist := studenAges["icham"]
	if exist {
		fmt.Println("Age:", age)
	} else {
		fmt.Println("No age found for ichami")
	}

	//delete an element
	delete(studenAges, "ichami") //remove ichami
	fmt.Println(studenAges)

	//iterating over a map
	for name, salary := range salaries {
		fmt.Println(name, " ends ", salary, " monthly")
	}

	//finding the length of a map
	fmt.Println("total employees:", len(salaries))

	//nested maps
	users := map[string]map[string]string{
		"ichami": {
			"age":  "10",
			"city": "Buea",
		},
		"feli": {
			"age":  "4",
			"city": "douala",
		},
	}

	//iterating over nested maps
	fmt.Println(users)
	for user, details := range users {
		fmt.Println("user:", user)
		for key, value := range details {
			fmt.Println(" ", key, ":", value)
		}
	}
}
