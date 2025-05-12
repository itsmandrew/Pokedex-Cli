package models

type Config struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type Command struct {
	Name        string
	Description string
	Callback    func(args []string) error
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocationArea mirrors the top‐level JSON for /location-area/{id or name}/.
type LocationArea struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

// PokemonEncounter is each element in “pokemon_encounters”.
type PokemonEncounter struct {
	// The nested “pokemon” resource
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []EncounterVersionDetail `json:"version_details"`
}

// EncounterVersionDetail tells you in which game versions this Pokémon appears
// and how rare it is.
type EncounterVersionDetail struct {
	Version NamedAPIResource `json:"version"`
	// The “rarity” field is an integer percentage (0–100).
	Rarity int `json:"rarity"`
}

type Pokemon struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	Type     NamedAPIResource `json:"stat"`
	BaseStat int              `json:"base_stat"`
}

type PokemonType struct {
	Type NamedAPIResource `json:"type"`
}
