package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cmd "github.com/itsmandrew/Pokedex-Cli/commands"
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

		fields := strings.Fields(scanner.Text())

		if len(fields) == 0 {
			fmt.Println("Please provide an argument")
			continue
		}

		cmdName, args := fields[0], fields[1:]

		val, ok := cmd.Table[cmdName]

		if !ok {
			fmt.Println("Unknown command:", cmdName)
			continue
		}

		if err := val.Callback(args); err != nil {
			fmt.Println("Issue with callback:", err)
			os.Exit(1)
		}
	}
}
