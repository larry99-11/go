package main

import (
	"bufio" // import bufio package
	"fmt"
	"os"
)

func main() {

	cliPrompt := "Pokedex >"
	userInput := bufio.NewScanner(os.Stdin)

	// we need to create an infinate loop to keep our prompt running
	for {
		fmt.Println(cliPrompt)

		//  To read input from the user, we need to use bufio.NewScanner. We'll scan the input, parse it, and then execute the corresponding command.
		// wait for user input
		if userInput.Scan() {
			//get the input text
			input := userInput.Text()

			// we have our options avalible for the user
			if input == "exit" {
				fmt.Println("Goodbye!")
				break
			} else if input == "help" {
				fmt.Println("Welcome to the Pokedex!")
				fmt.Println("Usage:")
				fmt.Println("help: Displays a help message")
				fmt.Println("exit: Exit the Pokedex")
			} else {
				fmt.Println("unknown command: ", input)
			}

		}

	}

	// next piece: getting user input and handling commands.

	//  To read input from the user, we need to use bufio.NewScanner. We'll scan the input, parse it, and then execute the corresponding command.
}
