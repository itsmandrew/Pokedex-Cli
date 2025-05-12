package utils

import (
	"math"
	"math/rand/v2"
)

func CalculateCatchRate(baseExperience int) float64 {
	const (
		maxBaseExp     = 600 // Approximate max base experience in the Pokémon games
		minCatchRate   = 5   // Minimum catch rate percentage (for highest base experience)
		maxCatchRate   = 95  // Maximum catch rate percentage (for lowest base experience)
		curveSteepness = 1.5 // Higher values make the curve steeper
	)

	// If baseExperience is negative or zero, return the maximum catch rate
	if baseExperience <= 0 {
		return maxCatchRate
	}

	if baseExperience >= maxBaseExp {
		return minCatchRate
	}

	// Normalize the base experience relative to the maximum
	normalizedExp := float64(baseExperience) / float64(maxBaseExp)

	// Calculate catch rate with an inverse exponential curve
	// This gives a smooth decrease from maxCatchRate to minCatchRate
	catchRate := maxCatchRate * math.Exp(-curveSteepness*normalizedExp)

	// Ensure the catch rate doesn't go below the minimum
	if catchRate < minCatchRate {
		catchRate = minCatchRate
	}

	// Round to one decimal place
	return math.Round(catchRate*10) / 10
}

func SimulateCatch(baseExperience int) bool {
	catchRate := CalculateCatchRate(baseExperience)
	randomValue := 100 * rand.Float64()

	// If the random value is less than the catch rate, the catch is successful
	return randomValue < catchRate
}
