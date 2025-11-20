package main

import (
	"bazel_training_example/lib/utils"
	"fmt"
)

func main() {
	message := utils.Greet("Buildkite")
	fmt.Println(message)
	fmt.Println("This is the hello service!")
}
