package main

import "fmt"

// string data types
var name = "ichami"      //type inference
var city string = "buea" //explicit type declaration
var occupation string

// numbers in go
// 1 interger number types
var ageOne int = 10
var ageTwo = 20

//bits and int
// var numOne int8 = 25 //all signed(negative and positive) 8 bit intergers range -127 to 128
// var numTwo int32 = 2323 //Range: -2147483648 through 2147483647
// var numThree uint8 = 10 //uint8 is the set of all unsigned 8-bit integers. Range: 0 through 255. no negative numbers

// float (float32 and float64)
var one float32 = 1.3
var two float64 = 343234.3 //float64 has a higher precision than float32
func main() {
	// country := "cameroon" //short declaration, only allowed inside functions
	// fmt.Println("Country:", country)
	fmt.Println(name, city, occupation)

	ageThree := 30
	fmt.Println(ageOne, ageTwo, ageThree)
	three := 134.4 //will be auto inferred as float64
	fmt.Println(one, two, three)
}
