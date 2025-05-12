package utils

import "testing"

func TestCalculateCatchRate(t *testing.T) {

	testCases := []struct {
		baseExp     int
		minExpected float64
		maxExpected float64
	}{
		// Should return max catch rate (lowest exp)
		{0, 95.0, 95.0},
		{-10, 95.0, 95.0},

		// Variable base experiences
		{50, 75.0, 85.0},
		{250, 45.0, 55.0},
		{500, 20.0, 30.0},

		// Should give lowest experience
		{600, 5.0, 5.0},
		{1000, 5.0, 5.0},
	}

	for _, tc := range testCases {
		result := CalculateCatchRate(tc.baseExp)

		if result < tc.minExpected || result > tc.maxExpected {
			t.Errorf("CalculateCatchRate(%d) = %.1f, expected between %.1f and %.1f",
				tc.baseExp, result, tc.minExpected, tc.maxExpected)
		}
	}
}
