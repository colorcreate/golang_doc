package test

import (
	"testing"

	"golang_doc/go-tdd/sub"
)

func TestAddTwoNumbers(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Error("Failed to add two numbers")
	} else {
		t.Log("success")
	}
}

func TestAddMultipleNumbers(t *testing.T) {
	result := Add(1, 2, 3, 4)
	if result != 10 {
		t.Error("Failed to add multiple numbers")
	}
}

func TestSubFromOtherPackage(t *testing.T) {
	result := sub.Sub(2, 3)
	if result != -1 {
		t.Error("failed to substract two numbers")
	}
	if testing.Short() {
		t.Skip("Skipping test because short flag is enabled")
	}
	if testing.Verbose() {
		t.Log("Hello there!")
	}
}
