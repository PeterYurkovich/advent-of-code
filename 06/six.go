package six

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/peteryurkovich/advent-of-code/helpers"
)

func Six() {
	fmt.Println("\n\n\n\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	stringFile := helpers.GetFileString("./06/a.txt")
	lines := strings.Split(stringFile, "\n")

	matrix := helpers.TwoDMatrix{}
	var guardLocation = helpers.Coordinate{X: 0, Y: 0}

	for i, line := range lines {
		stringArray := []string{}
		for j, character := range line {
			str := string(character)
			stringArray = append(stringArray, str)
			if str == "^" {
				guardLocation = helpers.Coordinate{X: j, Y: i}
			}
		}
		matrix.Data = append(matrix.Data, stringArray)
	}

	visitedCount, visitedLocations, _ := solveMap(guardLocation, matrix, false)

	fmt.Printf("Total Locations visited: %d\n", visitedCount)

	var loopCount = 0
	for visitedCoordinate := range visitedLocations {
		oldValue, err := matrix.At(visitedCoordinate)
		helpers.AssertError(err)
		if oldValue == "#" {
			continue
		}
		matrix.Data[visitedCoordinate.Y][visitedCoordinate.X] = "#"

		_, _, looped := solveMap(guardLocation, matrix, false)
		if looped {
			loopCount++
		}

		// cleanup the matrix to be how it was before
		matrix.Data[visitedCoordinate.Y][visitedCoordinate.X] = oldValue

	}

	fmt.Printf("Total Loops Caused: %d", loopCount)
}

type Direction int

var (
	North Direction = 0
	East  Direction = 1
	South Direction = 2
	West  Direction = 3
)

type Guard struct {
	location  helpers.Coordinate
	direction Direction
}

// new location, new direction, have we looped, error
func (guard Guard) nextLocation(matrix helpers.TwoDMatrix, log bool) (helpers.Coordinate, Direction, bool, error) {
	for i := 0; i < 4; i++ {
		nextDirection := (int(guard.direction) + i) % 4
		nextStep := potentialStepForward(guard.location, Direction(nextDirection))
		if log {
			fmt.Printf("nextLocation: %v, nextDirection: %d\n", nextStep, nextDirection)
		}
		newLocation, err := matrix.At(nextStep)
		if err != nil {
			// we cannot retrieve the data, the meaning here being we have left the map
			// return the same data as current as a way to indicate that we are done
			return guard.location, guard.direction, false, nil
		}
		if newLocation == "#" {
			continue
		}
		return nextStep, Direction(nextDirection), true, nil
	}
	return helpers.Coordinate{}, North, false, errors.New("yeah idk")
}

func potentialStepForward(coordinate helpers.Coordinate, direction Direction) helpers.Coordinate {
	if direction == North {
		return helpers.Coordinate{X: coordinate.X, Y: coordinate.Y - 1}
	} else if direction == South {
		return helpers.Coordinate{X: coordinate.X, Y: coordinate.Y + 1}
	} else if direction == West {
		return helpers.Coordinate{X: coordinate.X - 1, Y: coordinate.Y}
	} else if direction == East {
		return helpers.Coordinate{X: coordinate.X + 1, Y: coordinate.Y}
	}
	helpers.Assert(false, "Invalid Direction")
	return helpers.Coordinate{}
}

// return visited count & if the game ended in a loop
func solveMap(initalGuardLocation helpers.Coordinate, matrix helpers.TwoDMatrix, log bool) (int, map[helpers.Coordinate][]Direction, bool) {
	guard := Guard{location: initalGuardLocation, direction: North}

	var visitedCount = 1
	visitedLocations := map[helpers.Coordinate][]Direction{initalGuardLocation: {North}}

	for true {
		if log {
			fmt.Println(matrix)
			fmt.Printf("currentLocation: %v, currentDirection: %d\n", guard.location, guard.direction)
		}
		newCoordinate, newDirection, looped, err := guard.nextLocation(matrix, log)
		helpers.AssertError(err)

		location, ok := visitedLocations[newCoordinate]
		if !ok {
			visitedCount++
			visitedLocations[newCoordinate] = []Direction{newDirection}
			guard.location = newCoordinate
			guard.direction = newDirection
			continue
		}
		if slices.Contains(location, newDirection) {
			// we have left the map, OR are looping
			return visitedCount, visitedLocations, looped
		}
		visitedLocations[newCoordinate] = append(location, newDirection)
		guard.location = newCoordinate
		guard.direction = newDirection
	}
	helpers.Assert(false, "Shouldn't be possible to get here")
	return visitedCount, visitedLocations, false
}
