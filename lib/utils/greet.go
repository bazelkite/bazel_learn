package utils

import "fmt"

// Greet returns a greeting message
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Farewell returns a farewell message
func Farewell(name string) string {
	return fmt.Sprintf("Goodbye, %s!", name)
}
