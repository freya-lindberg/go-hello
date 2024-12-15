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
		t.Fatal(`Hello("Freya Lindberg") = %q %v, want match for %#q, nil`, msg, e, want)
	}
}
