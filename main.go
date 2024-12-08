package main

import (
	"os"

	dayOne "github.com/peteryurkovich/advent-of-code/01"
	dayTwo "github.com/peteryurkovich/advent-of-code/02"
	dayThree "github.com/peteryurkovich/advent-of-code/03"
	dayFour "github.com/peteryurkovich/advent-of-code/04"
	dayFive "github.com/peteryurkovich/advent-of-code/05"
	daySix "github.com/peteryurkovich/advent-of-code/06"
)

func main() {
	day := os.Args[1]
	if day == "one" {
		dayOne.One()
	} else if day == "two" {
		dayTwo.Two()
	} else if day == "three" {
		dayThree.Three()
	} else if day == "four" {
		dayFour.Four()
	} else if day == "five" {
		dayFive.Five()
	} else if day == "six" {
		daySix.Six()
	}
}
