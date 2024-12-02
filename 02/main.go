package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func assert(passed bool, reason string) {
	if !passed {
		panic(reason)
	}
}

func assertError(err error) {
	if err != nil {
		assert(false, err.Error())
	}
}

func Abs(potentiallyNegative int) int {
	if potentiallyNegative < 0 {
		return -1 * potentiallyNegative
	}
	return potentiallyNegative
}

func remove(s []int, i int) []int {
	newArray := make([]int, len(s))
	copy(newArray, s)
	return append(newArray[:i], newArray[i+1:]...)
}

func main() {
	fmt.Println("\n\n\n\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	byteFile, err := os.ReadFile("./a.txt")
	assertError(err)
	lines := strings.Split(string(byteFile), "\n")

	reports := make([][]int, len(lines))

	for i, line := range lines {
		splitLine := strings.Fields(line)
		levels := make([]int, len(splitLine))
		for j, level := range splitLine {
			intLevel, err := strconv.Atoi(level)
			assertError(err)
			levels[j] = intLevel
		}
		reports[i] = levels
	}

	var safeCount = 0
	safeReports := [][]int{}
	for _, report := range reports {
		if len(report) < 2 {
			continue
		}
		if testReport(report) {
			safeCount++
			safeReports = append(safeReports, report)
		}
	}

	fmt.Printf("Safe Report Count: %d\n", safeCount)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	var safeDampenedCount = 0
	unsafeReports := [][]int{}

	for _, report := range reports {
		lenReport := len(report)
		if lenReport < 2 {
			continue
		}
		var safe = false
		for k := 0; k < lenReport; k++ {
			removedReport := remove(report, k)
			if testReport(removedReport) {
				safeDampenedCount++
				safe = true
				break
			}
		}
		if !safe {
			unsafeReports = append(unsafeReports, report)
		}
	}

	fmt.Printf("Safe Dampened Report Count: %d\n", safeDampenedCount)
}

func testReport(report []int) bool {
	var direction = 1
	if report[0] > report[1] {
		direction = -1
	}
	for l := 1; l < len(report); l++ {
		difference := (report[l] - report[l-1]) * direction
		if difference > 3 || difference < 1 {
			return false
		}
	}

	return true
}
