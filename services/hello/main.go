package main

import (
	"bazel_learn/lib/utils"
	"fmt"
)

func main() {
	message := utils.Greet("Buildkite")
	fmt.Println(message)
	fmt.Println("This is the hello service!")
}
