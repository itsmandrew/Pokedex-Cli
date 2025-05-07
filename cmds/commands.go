package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/itsmandrew/Pokedex-Cli/api"
	pk "github.com/itsmandrew/Pokedex-Cli/internal"
	"github.com/itsmandrew/Pokedex-Cli/models"
)

var Table map[string]cliCommand
var Cache *pk.Cache

func init() {

	var cfg models.Config
	cfg.Next = api.LOCATION_URL
	cfg.Previous = ""

	Table = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    func() error { return CommandExit(&cfg) },
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    func() error { return CommandHelp(&cfg) },
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas",
			Callback:    func() error { return CommandMap(&cfg) },
		},
		"mapb": {
			Name:        "map",
			Description: "Displays the previous 20 location areas",
			Callback:    func() error { return CommandMapb(&cfg) },
		},
	}

	Cache = pk.NewCache(3 * time.Minute)
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func CommandExit(config *models.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(config *models.Config) error {

	fmt.Println("Usage:")
	fmt.Println()
	for k, com := range Table {
		fmt.Printf("%s: %v\n", k, com.Description)
	}

	return nil
}

func CommandMap(config *models.Config) error {

	cfg, err := api.GetLocationAreas(config.Next, Cache)
	if err != nil {
		return err
	}

	config.Next = cfg.Next
	config.Previous = cfg.Previous

	for _, area := range cfg.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func CommandMapb(config *models.Config) error {

	if config.Previous == "" {
		fmt.Println("no previous URL")
		return nil
	}

	cfg, err := api.GetLocationAreas(config.Previous, Cache)

	if err != nil {
		return err
	}

	config.Next = cfg.Next
	config.Previous = cfg.Previous

	for _, area := range cfg.Results {
		fmt.Println(area.Name)
	}

	return nil
}
