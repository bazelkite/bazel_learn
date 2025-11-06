package main

import "testing"

func TestGreet(t *testing.T) {
	result := greet("Test")
	expected := "Hello, Test!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
