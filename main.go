package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

//standard libraries in Go
/*Go's standard library provides us with a rich set of build-in packages that helps with various
functionalities like networking,strings,file handling and more*/

/*
1 fmt - Formatted I/O (printing and scanning)
the fmt package is used for input and output operations
*/

func main() {
	// printing output
	name := "ichami"
	fmt.Println("name:", name)     //printing with a new line
	fmt.Printf("hello %s\n", name) //formatted printing

	// taking input
	var firstName string
	fmt.Print("Enter your name:")
	fmt.Scanln(&firstName) //takes user input
	fmt.Println("welcome,", name)

	/*2 strings - String manipulation
	The string package helps in handling and manipulating strings
	*/

	text := "Hello, Go is awesome"
	fmt.Println(strings.ToUpper(text))                    //converts to uppercase
	fmt.Println(strings.ToLower(text))                    //converts to lowercase
	fmt.Println(strings.Contains(text, "Go"))             //checks if Go is present (true)
	fmt.Println(strings.ReplaceAll(text, "Go", "Golang")) //replace substring
	fmt.Println(strings.Split(text, " "))                 //split into a splice

	/*3 math - Mathematical functions
	The math package provides common mathematical functions
	*/
	fmt.Println("square root of 16:", math.Sqrt(16))   // Square root
	fmt.Println("2 raise to power 3:", math.Pow(2, 3)) // 2^3
	fmt.Println("round 4.6:", math.Round(4.6))         // Round to nearest integer
	fmt.Println("ceil 4.2:", math.Ceil(4.2))           // Round up
	fmt.Println("floor of 4.9:", math.Floor(4.9))      // Round down

	/*4 time - working with date and time
	the time package handles date, time and duration
	*/
	now := time.Now()
	formatted := now.Format("2006-01-02 15:04:05") // Go uses this reference date
	fmt.Println("current time:", now, "formatted:", formatted)

	fmt.Println("Wait for 2 seconds...")
	time.Sleep(2 * time.Second) // Pauses for 2 seconds
	fmt.Println("Done!")

	//Go uses "2006-01-02 15:04:05" as it reference time

	/*5 rand - Generating random numbers
	the math/rand package is used to generate random numbers
	*/
	rand.Seed(time.Now().UnixNano()) //set a different seed each time
	fmt.Println(rand.Intn(100))      //generate a random number between 0 - 100
	//use rand.Seed(time.Now().UnixNano()), to ensure new values everytime

	/*6 sort */

	ages := []int{40, 34, 100, 2, 54}
	sort.Ints(ages)
	fmt.Println(ages) //sort the splice in desc order

	index := sort.SearchInts(ages, 54) //returns the index of the now sorted splice
	fmt.Println(index)                 //expected answer 3

	names := []string{"ichami", "mispa", "feli", "brandon"}
	sort.Strings(names)
	fmt.Println(names)                              //returns the sorted names in desc order of alphabet
	fmt.Println(sort.SearchStrings(names, "mispa")) //returns the index of name mispa in splice of names (output:3)

	/*7 io/ioutil - Reading and Writing files
	The io/ioutil helps in reading and writing files*/
	// content := []byte("Hello Go")
	// err := os.WriteFile("test.txt", content, 0664) //create a new file

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	data, err := os.ReadFile("test.txt") //read file content
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("file content:", data)
}
