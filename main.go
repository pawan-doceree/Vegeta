package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for {
		fmt.Println("Welcome to the dice game")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Println("Enter your choice (1 or 2):")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, Please enter 1 or 2.")
			continue
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
		die1 := rand.IntN(6) + 1
		die2 := rand.IntN(6) + 1

		fmt.Printf("You rolled a %d and a %d\n", die1, die2)

		fmt.Println("total:", die1+die2)

		fmt.Print("Do you want to roll again? (y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no.")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thank for plaing! Bye")

		}
	}
}
