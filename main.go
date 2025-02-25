package main

import "fmt"

//Loops are essential in programming, and in Go,
//the for loop is the only looping construct. However,
//it can be used in different ways to handle various scenarios.

func main() {
	// 1 basic for loop
	for x := 0; x <= 5; x++ {
		fmt.Println("number:", x)
	}

	//2.1 infinite loop (do while)
	count := 1
	for {
		fmt.Println("loop running", count)
		count++
		if count > 3 {
			break //stop the loop when count > 3
		}
	}
	//2.2 infinite loop (while loop)
	n := 2
	for n > 0 {
		fmt.Println("countdown", n)
		n-- //decrement n in each iteration
	}

	//3 looping over array and slices
	nums := []int{10, 23, 30}
	for index, value := range nums { //similar to for in in js
		fmt.Printf("index: %d,value: %d\n", index, value)
	}
	//similarly, ignore the index if you dont need it
	for _, value := range nums {
		fmt.Printf("value: %d\n", value)
	}

	//4 looping over a map (key-value-pairs)
	person := map[string]int{"ichami": 10, "feli": 5}
	for key, value := range person {
		fmt.Printf("%s is %d years old\n", key, value)
	}

	//5 nested loops
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i:%d,j:%d\n", i, j)
		}
	}
}
