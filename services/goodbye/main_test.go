package main

import (
	"bazel_learn/lib/utils"
	"testing"
)

func TestFarewell(t *testing.T) {
	result := utils.Farewell("Test")
	expected := "Goodbye, Test!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
