package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Command struct {
	name		string
	description	string
	callback	func(*Config) error
}

type Config struct {
	apiClient		Client
	previousPage	*string
	nextPage		*string
}

func replStart(config *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		Prompt()
		reader.Scan()

		text := strings.ToLower(reader.Text())
		if command, exist := getCommands()[text]; exist {
			command.callback(config)
		} else {
			fmt.Println("Unknown command")
		}
	}
	Prompt()
}

func Prompt() {
	fmt.Print("Pokedex > ")
}

func getCommands() map[string]Command {
	return map[string]Command {
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "exit": {
            name:        "exit",
            description: "Exits the Pokedex",
            callback:    commandExit,
        },
		"mapf": {
			name: 		"mapf",
			description:"Moves forward in the list of locations",
			callback: 	commandMapf,
		},
		"mapb": {
			name:		"mapb",
			description:"Moves backward in the list of locations",
			callback:	commandMapb,
		},
    }
}