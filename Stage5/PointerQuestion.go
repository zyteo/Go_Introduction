package main

// ZY, 1 Feb 2022, 11.32am
// Learning Go through JetBrains Academy free lessons
//Pointers Question

//We've got the p pointer of type int as an argument in our process function.
//
//Assign 10 to the integer variable pointed to by the p pointer.

func solve(p *int) {
	// Put your code here
	*p = 10
}
