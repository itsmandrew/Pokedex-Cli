package api

import (
	"encoding/json"
	"io"
	"net/http"

	m "github.com/itsmandrew/Pokedex-Cli/models"
)

const LOCATION_URL = "https://pokeapi.co/api/v2/location-area/"

func GetLocationAreas(url string) (m.Config, error) {
	var config m.Config

	resp, err := http.Get(url)

	if err != nil {
		return m.Config{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return m.Config{}, err
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return m.Config{}, err
	}

	return config, nil
}
