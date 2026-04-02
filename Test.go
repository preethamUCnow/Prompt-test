package Prompt_test

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var choice string

	for {
		fmt.Println("\n--- User Info Program ---")

		fmt.Print("Enter your name: ")
		fmt.Scanln(&name)

		fmt.Print("Enter your age: ")
		fmt.Scanln(&age)
		
		fmt.Print("\nDo you want to continue? (yes/no): ")
		fmt.Scanln(&choice)

		if choice == "no" {
			fmt.Println("Goodbye! 👋")
			break
		}
	}
}
