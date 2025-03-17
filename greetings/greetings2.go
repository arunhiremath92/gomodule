package greetings

import "fmt"

func HelloFourthVersion(name string) string {
    message := fmt.Sprintf("%v", name)
    return message
}
