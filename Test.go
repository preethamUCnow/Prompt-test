package Prompt_test

import (
	"fmt"
)

func main() {
	var name string

	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	fmt.Println("Hello,", name, "! Welcome to Go 🚀")
}
