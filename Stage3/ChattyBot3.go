package main

import (
	"fmt"
	"strconv"
)

// ZY, 18 Jan 2022, 8.56am
// Learning Go through JetBrains Academy free lessons
//Intro to simple chatty bot 3

func main() {
	fmt.Println("Hello! My name is robot.")
	fmt.Println("I was created in 2022.")
	fmt.Println("Please, remind me your name.")

	// reading a name
	var name string
	fmt.Scan(&name)

	fmt.Println("What a great name you have, " + name + "!")
	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")

	// reading all remainders
	var remainder3, remainder5, remainder7, age int
	fmt.Scan(&remainder3)
	fmt.Scan(&remainder5)
	fmt.Scan(&remainder7)
	age = (remainder3*70 + remainder5*21 + remainder7*15) % 105

	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}
