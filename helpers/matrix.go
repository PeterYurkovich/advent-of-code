package helpers

import (
	"errors"
)

type TwoDMatrix struct {
	// First reference is y, second is x
	/**

	1, 2, 3
	4, 5, 6
	7, 8, 9

	6 is located at [1, 2]
	*/
	Data [][]string
}

type Coordinate struct {
	X int
	Y int
}

type Direction int

const (
	Negative Direction = -1
	Neither  Direction = 0
	Positive Direction = 1
)

func (mat TwoDMatrix) At(coordinate Coordinate) (string, error) {
	if coordinate.Y >= len(mat.Data) || coordinate.Y < 0 {
		return "", errors.New("Out of Bounds, Y")
	}
	row := mat.Data[coordinate.Y]
	if coordinate.X >= len(row) || coordinate.X < 0 {
		return "", errors.New("Out of Bounds, X")
	}
	return row[coordinate.X], nil
}

func (mat TwoDMatrix) GetFour(coordinate Coordinate, x Direction, y Direction) (string, error) {
	if x == 0 && y == 0 {
		return "", errors.New("Invalid direction")
	}
	var four string
	for i := 0; i < 4; i++ {
		singleLetter, err := mat.At(Coordinate{coordinate.X + i*int(x), coordinate.Y + i*int(y)})
		if err != nil {
			return "", err
		}
		four += singleLetter
	}
	return four, nil
}

/**
 * For a given X:
 *      a . b
 *        c
 *      d . e
 *
 * The return will be abcde
 */
func (mat TwoDMatrix) GetX(coordinate Coordinate) (string, error) {
	var five string
	firstLetter, err := mat.At(Coordinate{coordinate.X - 1, coordinate.Y - 1})
	if err != nil {
		return "", errors.New("Invalid coordinate")
	}
	five += firstLetter
	secondLetter, err := mat.At(Coordinate{coordinate.X + 1, coordinate.Y - 1})
	if err != nil {
		return "", errors.New("Invalid coordinate")
	}
	five += secondLetter
	thirdLetter, err := mat.At(Coordinate{coordinate.X, coordinate.Y})
	if err != nil {
		return "", errors.New("Invalid coordinate")
	}
	five += thirdLetter
	fourthLetter, err := mat.At(Coordinate{coordinate.X - 1, coordinate.Y + 1})
	if err != nil {
		return "", errors.New("Invalid coordinate")
	}
	five += fourthLetter
	fifthLetter, err := mat.At(Coordinate{coordinate.X + 1, coordinate.Y + 1})
	if err != nil {
		return "", errors.New("Invalid coordinate")
	}
	five += fifthLetter

	return five, nil
}
