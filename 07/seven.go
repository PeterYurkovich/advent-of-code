package seven

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/peteryurkovich/advent-of-code/helpers"
)

func Seven() {
	fmt.Println("\n\n\n\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	stringFile := helpers.GetFileString("./07/a.txt")
	lines := strings.Split(stringFile, "\n")
	problems := parseProblems(lines)

	solvedProblem := []Problem{}
	var solvedSum = 0
	for _, problem := range problems {
		// fmt.Printf("------new problem: %v\n", problem)
		res := solvableProblem(problem)
		if res {
			solvedProblem = append(solvedProblem, problem)
			solvedSum += problem.result
		}
		// fmt.Println(res)
	}
	fmt.Printf("Solved Problem Sum: %d", solvedSum)
}

type Operator string

var (
	Add      Operator = "+"
	Multiply Operator = "*"
)

type Problem struct {
	result    int
	values    []int
	operators []Operator
}

func parseProblems(lines []string) []Problem {
	problems := []Problem{}
	for _, line := range lines {
		segments := strings.Split(line, ": ")
		helpers.Assert(len(segments) == 2, "Unexpected number of segments")

		result, err := strconv.Atoi(segments[0])
		helpers.AssertError(err)

		stringValues := strings.Split(segments[1], " ")
		values := []int{}
		for _, stringValue := range stringValues {
			value, err := strconv.Atoi(stringValue)
			helpers.AssertError(err)
			values = append(values, value)
		}

		problems = append(problems, Problem{result: result, values: values, operators: []Operator{}})
	}
	return problems
}

func solvableProblem(problem Problem) bool {
	// fmt.Printf("problem: %v\n", problem)
	if len(problem.values) < 2 {
		return false
	}
	sum := problem.values[0] + problem.values[1]
	product := problem.values[0] * problem.values[1]
	// fmt.Printf("product: %d\n", product)
	if problem.result == sum {
		return true
	}
	if problem.result == product {
		return true
	}
	if sum > problem.result && product > problem.result {
		return false
	}
	if sum < problem.result {
		sumvalues := []int{sum}
		sumvalues = append(sumvalues, problem.values[2:]...)
		res := solvableProblem(Problem{result: problem.result, values: sumvalues, operators: []Operator{}})
		if res {
			return true
		}
	}
	if product < problem.result {
		productvalues := []int{product}
		productvalues = append(productvalues, problem.values[2:]...)
		return solvableProblem(Problem{result: problem.result, values: productvalues, operators: []Operator{}})
	}
	return false
}
