package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
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

		if err := command.callback(cfg); err != nil {
			fmt.Println(err)
		}
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
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: callbackHelp,
		},
		"map": {
			name: "map",
			description: "Lists the next page of location areas",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "Lists the previous page of location areas",
			callback: callbackMapb,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: callbackExit,
		},
	}
}
