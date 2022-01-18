package main

import "fmt"

// ZY, 18 Jan 2022, 7.43pm
// Learning Go through JetBrains Academy free lessons
//Control statements theory

func main() {
	number := 4
	if number%2 == 0 { // checks if the number is even
		fmt.Println("The number", number, "is even")
	}

	//It is important to know that you can extend an if statement with an else statement and execute another block of code if the condition is false. Let's take a look at an updated version of the code snippet above:
	number2 := 5
	if number2%2 == 0 { // checks if the number is even
		fmt.Println("The number", number2, "is even")
	} else {
		fmt.Println("The number", number2, "is odd")
	}

	//In contrast to the if-else statement that checks if the condition is either true or false, the else if statement can help us check several possible conditions. Let's evaluate the else if statement in action in the code snippet below:
	var score int = 77 // Your test score goes here.

	if score > 90 {
		fmt.Println("Your grade is A. Congratulations!")
	} else if score > 80 {
		fmt.Println("Your grade is B. Good enough.")
	} else if score > 70 {
		fmt.Println("Your grade is C. Could've done better!")
	} else if score > 60 {
		fmt.Println("Your grade is D. Study more next time!")
	} else {
		fmt.Println("Your grade is F. Terrible! you failed the test!")
	}

	//So far we've seen the if, else and else if statements. They work great for simple evaluation cases, but when we try to use them with a large number of conditions, the code becomes difficult to read. For situations like this, the switch statement provides a simple and organized way to choose between multiple conditions. Let's check out the switch statement in action, when it's used in a computer game programmed in Go:
	var selection int = 2 // Your selection goes here.

	switch selection {
	case 1:
		fmt.Println("Starting a new game!")
	case 2:
		fmt.Println("Loading a saved game.")
	case 3:
		fmt.Println("Quit the game.")
	default:
		fmt.Println("Invalid selection. Try again.")
	}
	//	Finally, if you are familiar with other programming languages such as Java or the C family of languages, you may have seen the break statement used at the end of each case body. In Go, the break statement is provided automatically, so there is no need for us to hard code it within our switch statement. In simple terms, the function of the break statement is to step out of the switch construct after executing the code block within the case.

	//The versatility of the switch statement

	//In Go, we have versatile switch statements. This means we can match a variable to a single value, multiple values, match it based on expressions and logic, or even use it without an expression, which would be the equivalent of a multiline if-else statement. Let's take a look at using the switch statement without an expression:
	var number3 int = 7 // Your number goes here.

	switch { // equivalent to switch true: this switch statement will always get executed.
	case number3%2 == 0:
		fmt.Println("The number", number3, "is even")
	case number3%2 != 0:
		fmt.Println("The number", number3, "is odd")
	}

	//	In this example, the switch statement will always get executed within our program. This switch statement will evaluate each case as a boolean expression; since this example is very similar to the previously explained code snippets, you can easily infer the output of the program based on the value of the number variable.

	//An important detail is that all case statements when working with the switch construct must be unique. Every constant, even in a list of multiple values, is checked against the entire switch statement, and if there is a duplicate value, an error will be reported when trying to execute our program. Let's take a look at an invalid switch statement:
	letter := "a"

	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("The letter", letter, "is a vowel")
	case "a", "e", "i", "o", "u":
		fmt.Println("The letter", letter, "is a vowel")
	default:
		fmt.Println("The letter", letter, "is a consonant")
	}

}
