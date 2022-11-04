package main

import (
	"testing"
)

//exits if invalid path
func TestInvalidPath(t *testing.T) {
	_, err := RenameFiles("abcd")
	if err == nil {
		t.Fatalf("The error wasn't trown")
	}
}

// passing valid pass returns 0
func TestValidPath(t *testing.T) {
	count, err := RenameFiles(".")
	if err != nil && count != 0 {
		t.Fatalf("%s", err)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {

}
