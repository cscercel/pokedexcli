package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		available_commands := getCommands()

		command, ok := available_commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		command.callback()
	}
}

func cleanInput(text string) []string {
	lower_case := strings.ToLower(text)
	split_string := strings.Fields(lower_case)

	return split_string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: callbackExit,
		},
	}
}
