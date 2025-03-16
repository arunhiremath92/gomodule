package greetings

import (
	"testing"
	"regexp"
)


func test_hello(t *testing.T){
	name:= "Gladys"
	msg := Hello(name)
	want := regexp.MustCompile(`\b`+name+`\b`)
	if !want.MatchString(msg) {
        t.Errorf(`Hello("Gladys") = %q, want match for %#q, nil`, msg, want)
    }
}
