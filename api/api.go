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
