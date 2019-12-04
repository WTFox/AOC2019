package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	fuelTests := []struct {
		mass,
		fuel float64
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tt := range fuelTests {
		got := calculateFuel(tt.mass)
		if got != tt.fuel {
			t.Errorf("got %g want %g", got, tt.fuel)
		}
	}
}
