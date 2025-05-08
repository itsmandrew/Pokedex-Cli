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

	err := fetchAndUnmarshal(url, cache, &config)

	return config, err
}

func GetAreaPokemon(location string, cache *pk.Cache) (m.LocationArea, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", location)
	var area m.LocationArea

	err := fetchAndUnmarshal(url, cache, &area)

	return area, err
}

func fetchAndUnmarshal(url string, cache *pk.Cache, target any) error {

	cacheData, cacheHit := cache.Get(url)
	var data []byte
	var err error

	if cacheHit {
		fmt.Println("LOG --- Cache hit")
		data = cacheData
	} else {
		fmt.Println("LOG --- Cache miss")

		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status: %s", resp.Status)
		}

		data, err = io.ReadAll(resp.Body)

		if err != nil {
			return err
		}

		cache.Add(url, data)
	}

	if err = json.Unmarshal(data, &target); err != nil {
		return err
	}

	return nil
}
