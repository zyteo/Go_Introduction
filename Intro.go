package main

import "fmt"

// ZY, 15 Jan 2022, 8pm
// Learning Go through JetBrains Academy free lessons

// Intro to Go
func main() {
	fmt.Println("Hello Go!")
	fmt.Println(MakeMeCompile())
}

// Comments lesson
// The MakeMeCompile() function will be automatically call within main() function
func MakeMeCompile() string {

	// x is a normal variable
	x := "You should fix this piece of code to make it compile"
	// We trusting you can do it!
	return x
}
