package main

// ZY, 16 Jan 2022, 9.15pm
// Learning Go through JetBrains Academy free lessons
// Split string URL

import "fmt"

//Imagine that you are currently on http://hyperskill.org/golang-track, and somehow you've found an enormous bug. However, our fixer, robot Beep-Boop, doesn't have a clue how to get the page to fix it.
func main() {
	fmt.Println(notBugButFeature())
}

//Your task is to split the URL http://hyperskill.org/golang-track, with the use of the protocol, domain, and path variables, so that the poor robot stops chilling out.
func notBugButFeature() (string, string, string) {
	// Declare three variables protocol, domain, path
	var (
		protocol = "http://" // put protocol here
		domain   = "hyperskill.org/"
		path     = "golang-track" // put path to resource here

	)
	// Sending values to fixer. DO NOT EDIT
	return protocol, domain, path
}
