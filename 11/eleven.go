package eleven

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/peteryurkovich/advent-of-code/helpers"
)

func Eleven() {
	fmt.Println("\n\n\n\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	stringFile := helpers.GetFileString("./11/a.txt")
	inputs := strings.Split(stringFile, " ")

	var stoneCounts = map[int64]int64{}
	for _, input := range inputs {
		stone, err := strconv.Atoi(input)
		helpers.AssertError(err)
		stoneCounts[int64(stone)] += 1
	}

	for range 75 {
		stoneCounts = blink(stoneCounts)
		fmt.Println(stoneCounts)
	}

	var counts = int64(0)
	for _, count := range stoneCounts {
		counts += count
	}
	fmt.Println(counts)
	// part two doesn't work cause of some overflow or something idk, wrote the exact same thing in JS and it worked
}

func blink(stoneCounts map[int64]int64) map[int64]int64 {
	newStoneCounts := map[int64]int64{}
	for stone, count := range stoneCounts {
		if stone == 0 {
			newStoneCounts[1] += count
		}
		stoneString := strconv.Itoa(int(stone))
		stoneLength := len(stoneString)
		if stoneLength%2 == 1 {
			newStoneCounts[stone*2024] = count
		} else {
			newStoneOne, err := strconv.Atoi(stoneString[:stoneLength/2])
			helpers.AssertError(err)
			newStoneCounts[int64(newStoneOne)] += count

			newStoneTwo, err := strconv.Atoi(stoneString[stoneLength/2:])
			helpers.AssertError(err)
			newStoneCounts[int64(newStoneTwo)] += count
		}
	}

	return newStoneCounts
}
