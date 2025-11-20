package main

import (
	"bazel_training_example/lib/utils"
	"fmt"
)

func main() {
	message := utils.Farewell("Buildkite")
	fmt.Println(message)
	fmt.Println("This is the goodbye service!")
}
