package main

import "fmt"

/*Functions in Go
A function in Go is declared using the "func" keyword
*/

func sayHello(n string) {
	fmt.Printf("Hello %s\n", n)
}

// functions that return a value
func Square(num int) int {
	return num * num
}

// functions returning multiple values
func divide(divident, divisor int) (int, int) {
	quotient := divident / divisor
	remainder := divident % divisor

	return quotient, remainder
}

// named return values
func calculateRect(length, width int) (area int) {
	area = length * width //no need for return area
	return
}

//variatic functions (multiple arguments)
//a variable can accept variable number of arguments using ...

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// function as a parameter(higher order function)
func applyOperations(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}
func add(x, y int) int {
	return x + y
}

// recursive functions
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	sayHello("ichami")
	square := Square(3)
	fmt.Printf("Square root: %d\n", square)

	q, r := divide(10, 3)
	fmt.Println("quotient:", q, "remainder:", r)

	fmt.Println("Area of the rectangle is:", calculateRect(7, 3))

	fmt.Println("sum of 1,2,3,4,5 is:", sum(1, 2, 3, 4, 5))

	fmt.Println("result:", applyOperations(2, 3, add))

	//anonymouse (lambda functions)
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("mutiply:", multiply(2, 3))

	fmt.Println("\nfactorial:", factorial(4))
}
