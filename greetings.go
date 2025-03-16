package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}

func HelloSecondVersion(name string) string {
    message := fmt.Sprintf("This was second Version of Hello, and welcome to Go Lang Modules %v", name)
    return message
}
