package main

import "fmt"

// ZY, 19 Jan 2022, 8.53am
// Learning Go through JetBrains Academy free lessons
//Loops question

func main() int {
	// put your code here
	for i := 2; i < 1024; i++ {
		fmt.Println(i)
	}

	//Print all odd numbers between 2 and 1023 using loops, each on a new line.
	for i := 2; i < 1024; i++ {
		if i%2 != 0 {
			fmt.Println(i)

		}
	}

	//	Write a program that:
	//
	//    Reads the integer from the stdin;
	//    Prints the value of its factorial.
	var input int
	fmt.Println("Pass an integer, and I'll calculate its factorial for you!")
	fmt.Scan(&input)

	func Factorial(n int) {
		multiplier := 1
	for i := n; i > 1; i-- {
		multiplier *= i
	}

		return multiplier

	}
	fmt.Println(Factorial(input))

	return 0
}
