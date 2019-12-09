package main

import (
	"reflect"
	"testing"
)

func TestCoordinate(t *testing.T) {
	t.Run("calculateManhattenDistance {0,0}", func(t *testing.T) {
		center := Coordinate{0, 0}
		other := Coordinate{1, 2}
		assertEqual(other.calculateManhattenDistance(center), 3, t)

		other = Coordinate{5, 2}
		assertEqual(other.calculateManhattenDistance(center), 7, t)
	})

	t.Run("calculateManhattenDistance {4,2}", func(t *testing.T) {
		center := Coordinate{4, 2}
		other := Coordinate{1, 2}
		assertEqual(other.calculateManhattenDistance(center), 3, t)

		other = Coordinate{5, 2}
		assertEqual(other.calculateManhattenDistance(center), 1, t)
	})

	t.Run("calculateManhattenDistance {-4,2}", func(t *testing.T) {
		center := Coordinate{0, 0}
		other := Coordinate{-5, 2}
		assertEqual(other.calculateManhattenDistance(center), 7, t)
	})

	t.Run("travel 3 blocks", func(t *testing.T) {
		center := Coordinate{0, 0}
		other := Coordinate{0, 3}
		got := center.travel(other)
		want := Coordinates{
			Coordinate{0, 1},
			Coordinate{0, 2},
			Coordinate{0, 3},
		}
		assertDeepEqual(got, want, t)
	})

	t.Run("travel negative 3 blocks horizontally", func(t *testing.T) {
		center := Coordinate{0, 0}
		destination := Coordinate{0, -3}
		got := center.travel(destination)
		want := Coordinates{
			Coordinate{0, -1},
			Coordinate{0, -2},
			Coordinate{0, -3},
		}
		assertDeepEqual(got, want, t)
	})
	t.Run("travel negative 3 blocks vertically", func(t *testing.T) {
		center := Coordinate{0, 0}
		destination := Coordinate{-3, 0}
		got := center.travel(destination)
		want := Coordinates{
			Coordinate{-1, 0},
			Coordinate{-2, 0},
			Coordinate{-3, 0},
		}
		assertDeepEqual(got, want, t)
	})
	t.Run("travel from non-center", func(t *testing.T) {
		curPos := Coordinate{367, 998}
		destination := Coordinate{0, 2}
		got := curPos.travel(destination)
		want := Coordinates{
			Coordinate{367, 999},
			Coordinate{367, 1000},
		}
		assertDeepEqual(got, want, t)
	})
}

func TestFindIntersections(t *testing.T) {
	coordinateWireMap := CoordinateWireMap{
		Coordinate{0, 1}: []int{0},
		Coordinate{0, 2}: []int{0},
		Coordinate{6, 7}: []int{0, 1},
		Coordinate{4, 3}: []int{1, 2},
	}
	got := findIntersections(coordinateWireMap)
	want := Coordinates{
		Coordinate{6, 7},
		Coordinate{4, 3},
	}
	assertDeepEqual(got, want, t)
}

func TestFindClosestIntersectionPoint(t *testing.T) {
	coords := Coordinates{
		Coordinate{6, 7},
		Coordinate{4, 3},
	}
	got := findDistanceToClosestIntersection(coords)
	want := Coordinate{4, 3}
	assertDeepEqual(got, want, t)
}

func TestInstruction(t *testing.T) {
	t.Run("convertToCoordinate", func(t *testing.T) {
		testCases := []struct {
			instruction string
			want        Coordinate
		}{
			{"L75", Coordinate{0, -75}},
			{"R75", Coordinate{0, 75}},
			{"D25", Coordinate{-25, 0}},
			{"U15", Coordinate{15, 0}},
		}

		for _, tc := range testCases {
			inst := Instruction{tc.instruction}
			assertDeepEqual(inst.convertToCoordinate(), tc.want, t)
		}
	})
}

func assertDeepEqual(got, want interface{}, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
func assertEqual(got, want int, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Returned: [%v], Expected: [%v]", got, want)
	}
}
