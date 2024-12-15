package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Freya Lindberg"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, e := Hello(name)

	if !want.MatchString(msg) || e != nil {
		t.Fatalf(`Hello("Freya Lindberg") = %q %v, want match for %#q, nil`, msg, e, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
