package commands

import (
	"fmt"
	"os"
)

var Table map[string]cliCommand

func init() {
	Table = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
	}
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp() error {

	fmt.Println("Usage:")
	fmt.Println()
	for k, com := range Table {
		fmt.Printf("%s: %v\n", k, com.Description)
	}

	return nil
}
