package three

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/peteryurkovich/advent-of-code/helpers"
)

func Three() {
	fmt.Println("\n\n\n\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	byteFile, err := os.ReadFile("./a.txt")
	helpers.AssertError(err)
	file := string(byteFile)

	fileLen := len(file)
	var sum = 0
	var do = true
	sums := []int{}
	for i := 0; i < fileLen-4; i++ {
		fourSlice := file[i : i+4]
		if fourSlice == "do()" {
			do = true
			continue
		}
		if fileLen > i+7 && file[i:i+7] == "don't()" {
			do = false
			continue
		}
		if !do {
			continue
		}
		if fourSlice != "mul(" {
			continue
		}
		valueOne, locationOne, err := getNumberFromPos(file, i+4, ",")
		if err != nil {
			continue
		}
		valueTwo, _, err := getNumberFromPos(file, i+5+locationOne, ")")
		if err != nil {
			continue
		}
		product := valueOne * valueTwo
		sums = append(sums, product)
		sum += product
	}

	fmt.Printf("Total Sum: %d\n", sum)
}

// Returns value, location of the separator, and an error if the number doesn't match
func getNumberFromPos(fullString string, location int, endValue string) (int, int, error) {
	var value = 0
	for i := 0; i < 4; i++ {
		character := string(fullString[location+i])
		if value > 0 && character == endValue {
			return value, i, nil
		}

		singleVal, err := strconv.Atoi(character)
		if err != nil {
			return 0, 0, err
		}
		value = value*10 + singleVal
	}
	return 0, 0, errors.New("Failed to end the string with the right value")
}
