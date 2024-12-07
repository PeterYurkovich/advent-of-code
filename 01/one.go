package one

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/peteryurkovich/advent-of-code/helpers"
)

func One() {
	fmt.Println("\n\n\n\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	a, err := os.ReadFile("./a.txt")
	helpers.AssertError(err)
	aStringArray := strings.Split(string(a), "\n")

	aListOne := make([]int, len(aStringArray))
	aListTwo := make([]int, len(aStringArray))

	for i, aString := range aStringArray {
		values := strings.Fields(aString)
		helpers.Assert(len(values) == 2, fmt.Sprintf("non 2 value: %s", aString))
		firstValue, err := strconv.Atoi(values[0])
		helpers.AssertError(err)
		secondValue, err := strconv.Atoi(values[1])
		helpers.AssertError(err)
		aListOne[i] = firstValue
		aListTwo[i] = secondValue
	}
	slices.Sort(aListOne)
	slices.Sort(aListTwo)

	aListLength := make([]int, len(aStringArray))
	aDifference := 0
	for j := 0; j < len(aStringArray); j++ {
		val := int(math.Abs(float64(aListOne[j] - aListTwo[j])))
		aListLength[j] = val
		aDifference += val
	}

	fmt.Printf("Total Difference: %d\n", aDifference)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	b, err := os.ReadFile("./a.txt")
	helpers.AssertError(err)
	bStringArray := strings.Split(string(b), "\n")

	bListOne := make([]int, len(bStringArray))
	bMapTwo := make(map[int]int)

	for i, bString := range bStringArray {
		values := strings.Fields(bString)
		helpers.Assert(len(values) == 2, fmt.Sprintf("non 2 value: %s", bString))
		firstValue, err := strconv.Atoi(values[0])
		helpers.AssertError(err)
		secondValue, err := strconv.Atoi(values[1])
		helpers.AssertError(err)
		bListOne[i] = firstValue
		bMapValue, ok := bMapTwo[secondValue]
		if !ok {
			bMapTwo[secondValue] = 1
		} else {
			bMapTwo[secondValue] = bMapValue + 1
		}
	}
	slices.Sort(bListOne)

	bListLength := make([]int, len(bStringArray))
	bSimilarity := 0
	for j, bItem := range bListOne {
		bMapValue, ok := bMapTwo[bItem]
		if !ok {
			bMapValue = 0
		}
		val := bItem * bMapValue
		bListLength[j] = val
		bSimilarity += val
	}

	fmt.Printf("Total Similarity: %d\n", bSimilarity)
	return
}
