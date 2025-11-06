package utils

import "testing"

func TestGreet(t *testing.T) {
	result := Greet("World")
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestFarewell(t *testing.T) {
	result := Farewell("World")
	expected := "Goodbye, World!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
