package main

import (
	"fmt"

	"github.com/LK4D4/testproject/testpkg"
)

func X() {
	fmt.Println("X")
}

func Y() {
	fmt.Println("Y")
}

func main() {
	testpkg.Function()
}
