package main

/*
Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.

Suppose the following input is supplied to the program: 8
Then, the output should be: 40320
*/

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Exercise 002")

	// reading input from console
	var input int
	fmt.Print("Please enter a number : ")
	_, err := fmt.Scanln(&input)

	// checking for error and low 0
	if err != nil {
		log.Fatal("Please enter a number")
	}

	result, err := Ex002(input)

	if err != nil {
		log.Fatalf("Error for input %v: %v", input, err)
	}

	fmt.Printf("Factorial of %d = %d", input, result)
}

// Ex002 returns a factorial of input
// if input == 0 returns 1 as per definition
// if input < 0 returns 0 and an error
func Ex002(input int) (uint64, error) {

	if input < 0 {
		return 0, fmt.Errorf("Cannot find factorial of a negative number %v", input)
	}

	if input == 0 {
		return 1, nil
	}

	res, err := Ex002(input - 1)

	if err != nil {
		return res, err
	}

	return uint64(input) * res, nil

}
