package four

import (
	"fmt"
	"strings"

	"github.com/peteryurkovich/advent-of-code/helpers"
)

func Four() {
	fmt.Println("\n\n\n\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	stringFile := helpers.GetFileString("./04/a.txt")
	lines := strings.Split(stringFile, "\n")

	matrix := helpers.TwoDMatrix{}

	for _, line := range lines {
		stringArray := []string{}
		for _, character := range line {
			str := string(character)
			stringArray = append(stringArray, str)
		}
		matrix.Data = append(matrix.Data, stringArray)
	}
	yLen := len(matrix.Data)
	helpers.Assert(yLen != 0, "No Data")
	xLen := len(matrix.Data[0])

	var fourCount = 0
	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {
			for i := helpers.Negative; i <= helpers.Positive; i++ {
				for j := helpers.Negative; j <= helpers.Positive; j++ {
					four, err := matrix.GetFour(helpers.Coordinate{X: x, Y: y}, i, j)
					if err != nil {
						continue
					}
					if four == "XMAS" {
						fourCount++
					}
				}
			}
		}
	}

	fmt.Printf("Number of XMAS: %d\n", fourCount)
	var xCount = 0
	for x := 0; x < xLen; x++ {
		for y := 0; y < yLen; y++ {
			five, err := matrix.GetX(helpers.Coordinate{X: x, Y: y})
			if err != nil {
				continue
			}
			if string(five[2]) == "A" {
				fmt.Println(five, x, y)
			}
			if five == "MMASS" || five == "MSAMS" || five == "SSAMM" || five == "SMASM" {
				xCount++
			}
		}
	}

	fmt.Printf("Number of X-MAS: %d\n", xCount)
}
