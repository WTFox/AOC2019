package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// CENTER represents the center (duh)
var CENTER Coordinate = Coordinate{0, 0}

// CoordinateWireMap keeps track of which wires are over what coordindates
type CoordinateWireMap map[Coordinate][]int

// Coordinates is a list of Coordinate
type Coordinates []Coordinate

// Coordinate represents a location on a grid
type Coordinate struct {
	x,
	y int
}

func (c Coordinate) travel(dest Coordinate) Coordinates {
	// omg so ugly..
	var newPositions = make(Coordinates, 0)

	if dest.x == 0 && dest.y > 0 {
		// going right
		for i := c.y + 1; i <= c.y+dest.y; i++ {
			newPositions = append(newPositions, Coordinate{c.x, i})
		}
	} else if dest.x > 0 && dest.y == 0 {
		// going up
		for i := c.x + 1; i <= c.x+dest.x; i++ {
			newPositions = append(newPositions, Coordinate{i, c.y})
		}
	} else if dest.x < 0 && dest.y == 0 {
		// going down
		for i := c.x + -1; i >= c.x+dest.x; i-- {
			newPositions = append(newPositions, Coordinate{i, c.y})
		}
	} else if dest.x == 0 && dest.y < 0 {
		// going left
		for i := c.y + -1; i >= c.y+dest.y; i-- {
			newPositions = append(newPositions, Coordinate{c.x, i})
		}
	}

	return newPositions
}

func (c Coordinate) calculateManhattenDistance(pos Coordinate) int {
	// | a - c | + | b - d |
	return int(math.Abs(float64(c.x-pos.x)) + math.Abs(float64(c.y-pos.y)))
}

// Instruction manages conversions between instructions and coordinates
type Instruction struct {
	value string
}

func (i Instruction) convertToCoordinate() Coordinate {
	direction := string(i.value[0])
	distance, _ := strconv.Atoi(i.value[1:])

	var result Coordinate
	switch direction {
	case "U":
		result = Coordinate{distance, 0}
	case "D":
		result = Coordinate{distance * -1, 0}
	case "L":
		result = Coordinate{0, distance * -1}
	case "R":
		result = Coordinate{0, distance}
	default:
		result = Coordinate{0, 0}
	}
	return result
}

func findIntersections(coordinateWireMap CoordinateWireMap) Coordinates {
	var intersectedCoordinates []Coordinate
	for k, v := range coordinateWireMap {
		if (k.x != 0 && k.y != 0) && len(v) > 1 {
			intersectedCoordinates = append(intersectedCoordinates, k)
		}
	}
	return intersectedCoordinates
}

func findDistanceToClosestIntersection(intersectedCoordinates Coordinates) int {
	closestDistance := math.Inf(0)

	center := Coordinate{0, 0}

	for _, coord := range intersectedCoordinates {
		distance := float64(coord.calculateManhattenDistance(center))
		if distance < closestDistance {
			closestDistance = distance
		}
	}
	return int(closestDistance)
}

func buildCoordinateWireMap(instructions []string) CoordinateWireMap {
	coordinateWireMap := make(CoordinateWireMap)
	for wireIndex, wire := range instructions {
		wire = strings.TrimSpace(wire)
		if wire == "" {
			continue
		}

		position := Coordinate{0, 0}
		coordinateWireMap[position] = append(coordinateWireMap[position], wireIndex)
		for _, input := range strings.Split(wire, ",") {
			stepsToDestination := position.travel(Instruction{input}.convertToCoordinate())
			for _, nextPosition := range stepsToDestination {
				if !contains(coordinateWireMap[nextPosition], wireIndex) {
					coordinateWireMap[nextPosition] = append(coordinateWireMap[nextPosition], wireIndex)
				}
				position = nextPosition
			}
		}
	}
	return coordinateWireMap
}

func dayThreePartOne(testInputs []string) {
	wireMapping := buildCoordinateWireMap(testInputs)
	intersectedCoordinates := findIntersections(wireMapping)
	lowestDistance := findDistanceToClosestIntersection(intersectedCoordinates)
	fmt.Println(lowestDistance)
}

func countStepsToIntersection(wire string, intersectedCoordinate Coordinate) int {
	position := CENTER
	steps := 0
	for _, input := range strings.Split(wire, ",") {
		stepsToDestination := position.travel(Instruction{input}.convertToCoordinate())
		for _, nextPosition := range stepsToDestination {
			if position.x == intersectedCoordinate.x && position.y == intersectedCoordinate.y {
				return steps
			}
			position = nextPosition
			steps++
		}
	}
	return steps
}

func dayThreePartTwo(testInputs []string) {
	wireMapping := buildCoordinateWireMap(testInputs)
	intersectedCoordinates := findIntersections(wireMapping)

	lowestNumber := math.Inf(0)
	for _, intersectedCoordinate := range intersectedCoordinates {
		wireOneStepCount := countStepsToIntersection(testInputs[0], intersectedCoordinate)
		wireTwoStepCount := countStepsToIntersection(testInputs[1], intersectedCoordinate)
		result := float64(wireOneStepCount + wireTwoStepCount)

		if result < lowestNumber {
			lowestNumber = result
		}
	}
	fmt.Println(lowestNumber)

}

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	testInputs := strings.Split(string(inputBytes), "\n")

	dayThreePartOne(testInputs)
	dayThreePartTwo(testInputs)
}
