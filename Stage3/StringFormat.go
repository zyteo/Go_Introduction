package main

import "fmt"

// ZY, 17 Jan 2022, 8.28pm
// Learning Go through JetBrains Academy free lessons
//String formatting theory

//String formatting helps us save time to solve complicated tasks.

//For example, if we need to print the hex value of a decimal number 100000 (which is 186a0) to the console, normally we wouldn't input this value by hand. We can solve this task more conveniently with string formatting.

//In this topic, we'll learn how to use string formatting in Go. Overall, there are two ways to do this:

//Printing a formatted string with the Printf() function;
//Formatting and returning the string by using the Sprintf() function instead of printing it to the console.

func main() {

	fmt.Printf("kitty1\nkitty2")
	fmt.Printf("\n")
	fmt.Printf("kitty1\tkitty2")
	// kitty1
	// kitty2
	// kitty1  kitty2

	//However, in most cases, we pass two or more arguments tofmt.Printf(). The first one is the string template that contains the text that we want to format, together with one or more verbs (or format specifiers) that specify how we want the string template to be formatted. All remaining arguments are variables that store the value of what we want to format.

	//For example, look at the following code snippet:

	s := "Golang"
	a := "young"
	fmt.Printf("We are %s hackers, we are so %s", s, a)

	//Formatting strings with different verbs

	//In this section, we are going to consider different format specifiers that we can use to format strings.

	//Let's begin with numeric values. Take a look at the following examples:
	// Using %d to print an int value
	fmt.Printf("%d", 36) // 36

	//When using this verb, we can specify the precision field length by a decimal number. To do this, we need to put a dot and then the precision length, like in the example below. If we don't specify a special value, Printf() will use the default value of the width and the precision length:
	// Using %f to print a float value with default width and precision
	fmt.Printf("%f", 20.66) // 20.660000

	// Using %f with the precision length 1
	fmt.Printf("%.1f", 20.33) // 20.3

	//Now, this is how we format strings, characters, and boolean values:

	// The %c formatter is convenient for formatting characters
	fmt.Printf("%c", 't') // t

	// The %s formatter is great to format a string
	fmt.Printf("%s", "This is a random string") // This is a random string

	// The %t formatter is suitable for Boolean values
	fmt.Printf("%t", false) // false

	//	And this is how we format the width of string values:
	// Formatting the width of a string
	fmt.Printf("|%8s|", "pikachu") // | pikachu| (this action added one additional space where the string begins)

	//Finally, we can use the %% verb to format a raw "%" string:

	fmt.Printf("%%") // %

	//Explicit argument indexes

	//Now, sometimes we might need to use a variable more than once or format multiple variables in a different order. We can do it by using [] square brackets.

	//The number inside the square brackets [] between the % character and the verb denote the replacement of the corresponding variable when formatting the string. This means we can use this symbol to express how many times a variable appears or change the order of the variables when formatting a string:

	z := "First variable"
	b := "Second variable"
	fmt.Printf("%[1]s | %[1]s \n", z) // First variable | First variable
	fmt.Printf("%[2]s | %[1]s", z, b) // Second variable | First variable

	//Default verb %v

	//It is important that we learn about one more verb — the %v verb. We can use it in many cases, such as printing the UNICODE number of an emoji or non-English characters, printing the raw form of an array or a slice in Go, or formatting complex numbers. In short, you can format all objects in Go as strings in some way, but in case we don't know how to format them without errors using a built-in verb, we can always pass them to %v. This is why we call it the default verb:
	fmt.Printf("%v", '😄') // 128516 (UNICODE number)

	fmt.Printf("%v", []int{1, 2, 3, 4, 5, 6}) // [1 2 3 4 5 6]

	fmt.Printf("%v", [3]int{9, 8, 7}) // [9 8 7]

	fmt.Printf("%v", 1+5i) // 1+5i (this is a complex number)

	//Other functions in the fmt package: Sprintf(), Print(), Sprint(), Println(), Sprintln()

	//We can create and return strings using Sprintf(). This function creates formatted strings without directly printing them:
	//first := fmt.Sprintf("%d", 500) // 'first' variable now has the value of 500

	//binaryVariable := fmt.Sprintf("%b", 500) // 'binaryVariable' variable now has the value of 111110100

	//Print(), Sprint(), Println(), Sprintln() use the default verb %v, so we don't need to add any additional string templates or verbs with them.

	//All these functions can take more than one argument, and the string format in this case will be concatenated.

	//If the function name has the prefix S, it means these functions will return the formatted string as the data object instead of directly printing it to the console. And if the function name has the suffix ln, it means the function will add an additional newline at the end of the formatted string.

	//The Print() function uses the same style of arguments as Printf():
	fmt.Print("Hello") // Hello

	x := "World"

	fmt.Print("Hello, ", x) // Hello, World

	//	And here's an example of a function that ends with the ln suffix:
	c := fmt.Sprintln("if you try to print s,",
		"it will automatically print a new line at the end of the string")

	fmt.Print(c, "Newline")

	/*
	   Output:
	   if you try to print s, it will automatically print a new line at the end of the string
	   Newline
	*/

	//Formatting multiline string with backticks

	//Additionally, we can use raw string literals by wrapping string sequences between the back quotes `...`.

	//This technique can help us easily write multiline formatted strings. Notice that special characters like \n newline will not be interpreted within raw string literals:

	fmt.Printf(`%s 
and
the
brave
new\n
world\n`, "Go")

}
