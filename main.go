package main

import (
	"fmt"
	"strings"
)

func main() {

	value := "Charmander Bulbasaur PIKACHU"

	arr := cleanInput(value)

	fmt.Printf("result: %v\n", cleanInput(value))

	for _, val := range arr {
		fmt.Printf("%v\n", val)
	}
}

func cleanInput(text string) []string {

	lower := strings.ToLower(text)

	return strings.Fields(lower)
}
