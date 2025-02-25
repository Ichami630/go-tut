package main //the greetings.go and the main.go files needs to belong to thesame package inorder to share variables and functions between the two files

import "fmt"

var points = []int{10, 24, 4}

func sayHi(n string) {
	fmt.Println("Hello", n)
}
func showScore() {
	fmt.Println("Your score for maths was:", score) //accessing the score variable from the main.go file
}
