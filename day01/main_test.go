package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	fuelTests := []struct {
		mass,
		fuel float64
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, tt := range fuelTests {
		got := calculateFuelRequirement(tt.mass)
		if got != tt.fuel {
			t.Errorf("got %g want %g", got, tt.fuel)
		}
	}
}
