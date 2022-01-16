package main

// ZY, 16 Jan 2022, 8.27pm
// Learning Go through JetBrains Academy free lessons
//Primitive types question

import (
	"fmt"
)

func main() {
	var r rune
	r = 'a'

	var i int
	i = 1

	var b bool
	b = true

	fmt.Printf("%c %d %t", r, i, b) // do not remove this line
}
