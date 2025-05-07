package models

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Config struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}
