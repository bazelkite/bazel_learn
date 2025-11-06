package main

import (
	"fmt"
)

func farewell(name string) string {
	return fmt.Sprintf("Goodbye, %s!", name)
}

func main() {
	message := farewell("Buildkite")
	fmt.Println(message)
	fmt.Println("This is the goodbye service!")
}
