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
		fmt.Println("LOG --- Cache hit")
		data = cacheData

	} else {

		fmt.Println("LOG --- Cache miss")

		resp, err := http.Get(url)

		if err != nil {
			return m.Config{}, err
		}

		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)

		if err != nil {
			return m.Config{}, err
		}
		cache.Add(url, data)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return m.Config{}, err
	}

	return config, nil
}

func GetAreaPokemon(location string, cache *pk.Cache) (m.LocationArea, error) {

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", location)

	cacheData, cacheHit := cache.Get(url)
	var area m.LocationArea

	var data []byte

	if cacheHit {
		fmt.Println("LOG --- Cache hit")
		data = cacheData

	} else {
		fmt.Println("LOG --- Cache miss")

		resp, err := http.Get(url)

		if err != nil {
			return m.LocationArea{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return m.LocationArea{}, fmt.Errorf("unexpected status: %s", resp.Status)
		}

		data, err = io.ReadAll(resp.Body)

		if err != nil {
			return m.LocationArea{}, err
		}

		cache.Add(url, data)
	}

	if err := json.Unmarshal(data, &area); err != nil {
		return m.LocationArea{}, err
	}

	return area, nil
}
