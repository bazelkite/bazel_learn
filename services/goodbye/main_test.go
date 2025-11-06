package main

import "testing"

func TestFarewell(t *testing.T) {
	result := farewell("Test")
	expected := "Goodbye, Test!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
