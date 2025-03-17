package greetings

import (
	"regexp"
	"testing"
)

func Test_hello(t *testing.T) {
	name := "Gladys"
	msg := Hello(name)
	want := regexp.MustCompile(`\b` + name + `\b`)
	if !want.MatchString(msg) {
		t.Errorf(`Hello("Gladys") = %q, want match for %#q, nil`, msg, want)
	}
}

func Test_helloSecondVersion(t *testing.T) {
	name := "Gladys"
	msg := HelloSecondVersion(name)
	want := regexp.MustCompile(`\b` + name + `\b`)
	if !want.MatchString(msg) {
		t.Errorf(`Hello("Gladys") = %q, want match for %#q, nil`, msg, want)
	}
}

func Test_helloThirdVersion(t *testing.T) {
	name := "Gladys"
	msg := HelloThirdVersion(name)
	want := regexp.MustCompile(`\b` + name + `\b`)
	if !want.MatchString(msg) {
		t.Errorf(`Hello("Gladys") = %q, want match for %#q, nil`, msg, want)
	}
}

func Test_helloFourthVersion(t *testing.T) {
	name := "Gladys"
	want := "Gladys"
	msg := HelloFourthVersion(name)
	if want != msg {
		t.Errorf(`Hello("Gladys") = %q, want match for %#q, nil`, msg, want)
	}
}
