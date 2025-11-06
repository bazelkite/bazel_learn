package main

import (
	"fmt"
)

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	message := greet("Buildkite")
	fmt.Println(message)
	fmt.Println("This is the hello service!")
}
