package main

import (
	"fmt"
)

// ZY, 16 Jan 2022, 11.17am
// Learning Go through JetBrains Academy free lessons
// Theory: Input/Output

func main() {
	fmt.Print("Hello, World!") // Writing "Hello, World!" to the stdout

	//	When you run a program, the program output (except error messages) goes to the standard output or stdout.
	//For multi-line printing, use the special symbol \n to escape a character and show where a new line begins:
	fmt.Print("Hello, World!\n")
	fmt.Print("This is my first Go program!\n")

	//In addition, there is a special function from this package that appends \n for us. We can rewrite the example above using the Println function:
	fmt.Println("Hello, World!")                // Writing "Hello, World!"\n to the stdout
	fmt.Println("This is my first Go program!") // Writing "This is my first Go program!"\n to the stdout

	//Printing different variables
	//Now, Print and Println functions from the fmt package can print not only strings. As you progress, you will be able to format output in a sophisticated manner. But let's not be getting ahead of ourselves:
	fmt.Println(true)        // Printing type bool,   will print true
	fmt.Println(1023493)     // Printing type int,    will print 1023493
	fmt.Println("my string") // Printing type string, will print my string
	fmt.Println(12.4)        // Printing type float,  will print 12.4
	fmt.Println('A')         // Printing type rune,   will print 65
	//	The rune type is an alias for the int32 type, so it is printed as an integer by default.

	//	Print functions can also take several arguments and display them with space separation:
	var boolean = true
	var integer = 1023493
	var str = "my string"
	var float = 12.4
	var character = 'A'

	fmt.Println(boolean, integer, str, float, character) // Will print true 1023493 my string 12.4 65

	//	So far, we have only worked with sending data to the stdout. Now, let's see how we can obtain data from the standard input or stdin. The fmt.Scan function will help us with that. The function arguments are pointers to variables where they will store the input data. It stores successive space-separated values in the successive arguments. New lines are also counted as spaces:
	var name string

	fmt.Scan(&name)   // Reading from the stdin into the name variable
	fmt.Println(name) // Writing to the stdout the name you've entered
	// on the previous step
}
