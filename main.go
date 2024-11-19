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
	callback	func() error
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		Prompt()
		reader.Scan()

		text := strings.ToLower(reader.Text())
		if command, exist := getCommands()[text]; exist {
			command.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
	Prompt()
}

func Prompt() {
	fmt.Print("Pokedex > ")
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println(" ")
	for _, command := range getCommands() {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
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
    }
}

