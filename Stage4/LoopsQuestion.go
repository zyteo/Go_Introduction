package main

import "fmt"

// ZY, 19 Jan 2022, 8.53am
// Learning Go through JetBrains Academy free lessons
//Loops question

func main() {
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

	fmt.Println(Factorial(input))

	//	Write a program that:
	//
	//    Reads the integer from the stdin;
	//    Reverses its digits;
	//    Prints the result.

	//Hint:You can "shift" an integer to the left or to the right by multiplying or dividing it by 10:
	//1234 * 10 = 12340
	//1234 / 10 = 123
	//Hint2:You can retrieve the last digit of an integer by taking the reminder of its division by 10:
	//1234 % 10 = 4

	var intReverse int
	fmt.Println("Pass an integer, and I'll reverse it!")
	fmt.Scan(&intReverse)
	var reversed string

	//	using a for loop, I create a new variable, and append to the variable as I get the new digits
	for intReverse != 0 {
		reversed = reversed + string(intReverse%10)
		intReverse = intReverse / 10
	}
	fmt.Println(reversed)

}

func Factorial(n int) int {
	multiplier := 1
	for i := n; i > 1; i-- {
		multiplier *= i
	}

	return multiplier
}
