package main

// ZY, 16 Jan 2022, 9.21pm
// Learning Go through JetBrains Academy free lessons
// Built string URL

func makeURL() (redirectURL string) {
	// Concatenate response parts into the redirectURL variable.
	// Parts are provided below:
	var (
		protocol   = "https://"
		domain     = "hyperskill.org/"
		path       = "golang-track?"
		parameters = "gems=1500"
	)
	redirectURL = protocol + domain + path + parameters
	return
}
