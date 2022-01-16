package main

import "fmt"

// ZY, 16 Jan 2022, 9.25pm
// Learning Go through JetBrains Academy free lessons
//Intro to simple chatty bot 2

func main() {
	fmt.Println("Hello! My name is robot.")
	fmt.Println("I was created in 2022.")
	fmt.Println("Please, remind me your name.")

	// reading a name
	var name string
	fmt.Scan(&name)

	fmt.Println("What a great name you have, " + name + "!")
}
