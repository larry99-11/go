package main

import (
	"bufio" // import bufio package
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	cliPrompt := "Pokedex >"
	scanner := bufio.NewScanner(os.Stdin)

	// we need to create an infinate loop to keep our prompt running
	for {
		fmt.Print(cliPrompt)
		//cleaned := cleanInput()

		//  To read input from the user, we need to use bufio.NewScanner. We'll scan the input, parse it, and then execute the corresponding command.
		// wait for user input
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		// basically if we recieve no text continue the loop right back at the top
		if len(cleaned) == 0 {
			continue
		}

	}

}

func cleanInput(str string) []string { // returns a slice of strings
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
