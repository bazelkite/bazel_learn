package main

import (
	"bazel_learn/lib/utils"
	"fmt"
)

func main() {
	message := utils.Farewell("Buildkite")
	fmt.Println(message)
	fmt.Println("This is the goodbye service!")
}
