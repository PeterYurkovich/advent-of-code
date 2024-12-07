package five

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/peteryurkovich/advent-of-code/helpers"
)

func Five() {
	fmt.Println("\n\n\n\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	stringFile := helpers.GetFileString("./05/a.txt")
	segments := strings.Split(stringFile, "\n\n")
	helpers.Assert(len(segments) == 2, "Move or less segments then expected")

	rules := parseRules(segments[0])
	updates := parseUpdates(segments[1])

	invalidUpdates := [][]int{}
	var sumMiddle = 0
	for _, update := range updates {
		var updateValid = true
		for i, page := range update {
			if !updateValid {
				continue
			}
			pageRules, ok := rules[page]
			if !ok {
				continue
			}
			for _, pageRule := range pageRules {
				if !pageRule.valid(update, i) {
					updateValid = false
				}
			}
		}
		if updateValid {
			sumMiddle += update[(len(update)-1)/2]
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	fmt.Printf("Sum Middles: %d\n", sumMiddle)

	var invalidSumMiddle = 0
	for _, invalidUpdate := range invalidUpdates {
		afters := map[int][]Rule{}
		for _, page := range invalidUpdate {
			pageRules, ok := rules[page]
			if !ok {
				continue
			}
			for _, pageRule := range pageRules {
				// The rule is relevant for the update and the after value is the current page
				// this means that the other value must be before it
				if pageRule.after == page && slices.Contains(invalidUpdate, pageRule.before) {
					afterList, ok := afters[page]
					if !ok {
						afters[page] = []Rule{pageRule}
					} else {
						afters[page] = append(afterList, pageRule)
					}
				}
			}
		}

		validOrder := []int{}
		// chain variable
		for true {
			for _, page := range invalidUpdate {
				if slices.Contains(validOrder, page) {
					continue
				}
				pageAfters, ok := afters[page]
				if !ok || len(pageAfters) == 0 {
					validOrder = append(validOrder, page)
					continue
				}

				neededRules := []Rule{}
				for _, pageAfter := range pageAfters {
					if !slices.Contains(validOrder, pageAfter.before) {
						neededRules = append(neededRules, pageAfter)
					}
				}
				afters[page] = neededRules
			}
			if len(validOrder) == len(invalidUpdate) {
				break
			}
		}
		invalidSumMiddle += validOrder[(len(validOrder)-1)/2]
	}

	fmt.Printf("Sum Invalid Middles: %d", invalidSumMiddle)
}

func parseRules(ruleSegment string) map[int][]Rule {
	rules := map[int][]Rule{}

	ruleLines := strings.Split(ruleSegment, "\n")
	for _, ruleLine := range ruleLines {
		ruleSegments := strings.Split(ruleLine, "|")
		helpers.Assert(len(ruleSegments) == 2, "More or less rule fragments than expected")

		before, err := strconv.Atoi(ruleSegments[0])
		helpers.AssertError(err)
		after, err := strconv.Atoi(ruleSegments[1])
		helpers.AssertError(err)

		rule := Rule{before, after}

		beforeList, ok := rules[before]
		if !ok {
			rules[before] = []Rule{rule}
		} else {
			rules[before] = append(beforeList, rule)
		}

		afterList, ok := rules[after]
		if !ok {
			rules[after] = []Rule{rule}
		} else {
			rules[after] = append(afterList, rule)
		}
	}

	return rules
}

func parseUpdates(updateSegment string) [][]int {
	updates := [][]int{}

	updateLines := strings.Split(updateSegment, "\n")
	for _, updateLine := range updateLines {
		pages := strings.Split(updateLine, ",")

		update := []int{}
		for _, page := range pages {
			intPage, err := strconv.Atoi(page)
			helpers.AssertError(err)
			update = append(update, intPage)
		}
		updates = append(updates, update)
	}

	return updates
}

type Rule struct {
	before int
	after  int
}

func (rule Rule) valid(update []int, index int) bool {
	if rule.after == update[index] {
		for j := index + 1; j < len(update); j++ {
			if update[j] == rule.before {
				return false
			}
		}
	} else if rule.before == update[index] {
		for k := index - 1; k >= 0; k-- {
			if update[k] == rule.after {
				return false
			}
		}
	}
	return true
}
