package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/itsmandrew/Pokedex-Cli/api"
	pk "github.com/itsmandrew/Pokedex-Cli/internal"
	"github.com/itsmandrew/Pokedex-Cli/models"
)

var Table map[string]models.Command
var Cache *pk.Cache

func init() {

	var cfg models.Config
	cfg.Next = api.LOCATION_URL
	cfg.Previous = ""

	Table = map[string]models.Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    func(args []string) error { return CommandExit(&cfg, args) },
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    func(args []string) error { return CommandHelp(&cfg, args) },
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of 20 location areas",
			Callback:    func(args []string) error { return CommandMap(&cfg, args) },
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 location areas",
			Callback:    func(args []string) error { return CommandMapb(&cfg, args) },
		},
		"explore": {
			Name:        "explore",
			Description: "Lists all the possible pokemon encounters in a location area",
			Callback:    func(args []string) error { return CommandExplore(&cfg, args) },
		},
		"catch": {
			Name:        "catch",
			Description: "Give the user a chance to catch a Pokemon",
			Callback:    func(args []string) error { return CommandCatch(&cfg, args) },
		},
	}

	Cache = pk.NewCache(3 * time.Minute)
}

func CommandExit(config *models.Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(config *models.Config, args []string) error {

	fmt.Println("Usage:")
	fmt.Println()
	for k, com := range Table {
		fmt.Printf("%s: %v\n", k, com.Description)
	}

	return nil
}

func CommandMap(config *models.Config, args []string) error {

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

func CommandMapb(config *models.Config, args []string) error {

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

func CommandExplore(config *models.Config, args []string) error {

	if len(args) == 0 {
		fmt.Println("No location provided for explore command")
		return nil
	}

	locationName := args[0]

	area, err := api.GetAreaPokemon(locationName, Cache)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Println("Found Pokemon:")
	for _, pk := range area.PokemonEncounters {
		fmt.Printf("- %s\n", pk.Pokemon.Name)
	}
	return nil
}

func CommandCatch(config *models.Config, args []string) error {

	if len(args) == 0 {
		fmt.Println("Specify pokemon to catch")
		return nil
	}

	pokeName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)
	wildPokemon, err := api.GetPokemon(pokeName, Cache)

	if err != nil {
		return err
	}

	baseExperience := wildPokemon.BaseExperience
	fmt.Println(wildPokemon.BaseExperience)

	caught := pk.SimulateCatch(baseExperience)

	switch caught {
	case true:
		fmt.Printf("%s was caught!\n", pokeName)
		// TODO add to Pokedex
	case false:
		fmt.Printf("%s escaped!\n", pokeName)

	}

	return nil
}
