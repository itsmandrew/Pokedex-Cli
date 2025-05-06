package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cmd "github.com/itsmandrew/Pokedex-Cli/cmds"
)

const COMMAND_KEY string = "Pokedex > "

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	for {

		fmt.Printf("%s", COMMAND_KEY)

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprint(os.Stderr, "shit broken")
			}
			fmt.Println("\nGoodbye")
		}

		line := strings.TrimSpace(scanner.Text())
		val, ok := cmd.Table[line]
		if !ok {
			fmt.Println("Unknown command:", line)
			continue
		}

		if err := val.Callback(); err != nil {
			fmt.Println("Issue with callback:", err)
			os.Exit(1)
		}
	}
}

func cleanInput(text string) []string {

	lower := strings.ToLower(text)

	return strings.Fields(lower)
}
