package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pk "github.com/itsmandrew/Pokedex-Cli/internal"
	m "github.com/itsmandrew/Pokedex-Cli/models"
)

const LOCATION_URL = "https://pokeapi.co/api/v2/location-area/"

func GetLocationAreas(url string, cache *pk.Cache) (m.Config, error) {
	var config m.Config

	cacheData, cacheHit := cache.Get(url)

	var data []byte

	if cacheHit {
		data = cacheData
		fmt.Println("Cache hit")

	} else {

		fmt.Println("Cache miss")

		resp, err := http.Get(url)

		if err != nil {
			return m.Config{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		cache.Add(url, data)

		if err != nil {
			return m.Config{}, err
		}
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return m.Config{}, err
	}

	return config, nil
}

func GetAreaPokemon(location string) (*m.LocationArea, error) {

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", location)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for non-200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var area m.LocationArea
	if err := json.NewDecoder(resp.Body).Decode(&area); err != nil {
		return nil, err
	}
	return &area, nil
}
