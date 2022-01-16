package main

import (
	"fmt"
)

// ZY, 16 Jan 2022, 2.23pm
// Learning Go through JetBrains Academy free lessons
// Input/Output simple program

func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ") // Writing a request message to the stdout
	fmt.Scan(&name)                // Reading from the stdin into the name variable
	fmt.Println("")                // Going to the next line by writing /n to the stdout

	fmt.Print("Enter your age: ") // Writing a request message to the stdout
	fmt.Scan(&age)                // Reading from the stdin into the age variable
	fmt.Println("")               // Going to the next line by writing /n to the stdout

	fmt.Println(name, age) // Writing to the stdout the values of name and
	// age variables that you have entered

	//Write a program that reads two integers, adds them up, and prints the result.
	var firstInt, secondInt int

	fmt.Print("Enter your integers: ")
	fmt.Scan(&firstInt, &secondInt)

	sum := firstInt + secondInt
	fmt.Println(sum)

	fmt.Print("First line\n\n")
	fmt.Print("Third line\n\n")
	fmt.Print("Fifth line")

}
