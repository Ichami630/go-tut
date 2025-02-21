package main

//Arrays in Go

import "fmt"

func main() {
	//arrays in Go have a fixed length
	//declaring an array in go
	//elements are initialised with 0 values
	//access elements using index

	var ages [3]int = [3]int{10, 45, 20}
	var nums = [3]int{1, 2, 3} //alternate way to create an array in Go

	nums[0] = 23

	fmt.Println(ages)      //prints the values of an array
	fmt.Println(len(nums)) //prints the length of an array

	//use '...' to let Go determine the size of the array
	colors := [...]string{"red", "yellow", "blue"}
	fmt.Println("colors:", colors)

	//multidimensional arrays
	matrix := [2][3]int{ //creates a multidimensional array of 2 rows 3 columns
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("Multidimensional aray(2x3):", matrix)

	//looping through an array
	for i := 0; i < len(colors); i++ {
		fmt.Println("Color at index ", i, ":", colors[i])
	}

	//slices in Go (use arrays under the hood)

	//Donot have a fixed size
	//slices are declared using [] without the size

	scores := []int{1, 2, 3, 4} //declaring a new slice, also Go infers the length for us
	scores[2] = 10

	fmt.Println(scores, len(scores)) //scores slice length is 4

	//slice functions
	//1 Append (add new element at the end of the slice)

	//appending new values
	//this actually returns a slice for us
	scores = append(scores, 20)      //returns a new slice scores now of length 5
	fmt.Println(scores, len(scores)) //the length now changes from 4 to 5

	//2 copy (simply copy a slice)
	source := []int{10, 20, 30}
	dest := make([]int, len(source)) // Create slice with same length
	copy(dest, source)               // Copy data

	fmt.Println("source:", source, "dest:", dest)

	//slice ranges(creating a slice from an array)
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice1 := arr1[1:4] //including index 1 to 3 (excluding index 4)

	fmt.Println("slice1:", slice1) //output [2,3,4]
	slice2 := arr1[:2]             //from index 0 to 2 output [1,2]
	fmt.Println("slice2:", slice2)

	slice3 := arr1[1:]             //from index 0 to the last index of arr1
	fmt.Println("slice3:", slice3) //expected output [2,3,4,5]

	//multidimensional slices (slices of slices (flexible))

	grid := [][]int{ //slices allows rows of different lengths unlike arrays
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
	}

	fmt.Println("slices of slices(grid):", grid)

}
