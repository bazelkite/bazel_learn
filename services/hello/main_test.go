package main

import (
	"bazel_learn/lib/utils"
	"testing"
)

func TestGreet(t *testing.T) {
	result := utils.Greet("Test")
	expected := "Hello, Test!"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
