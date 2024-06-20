package main

import (
	"fmt"
	"strings"
	"log"
	"os"
	"regexp"
	"reflect"
)

func buildRegex(str []string) (*regexp.Regexp, *regexp.Regexp) {
	pre := "^[?.]*"
	dam := "([?#]{%s})"
	mid := "[?.]+"
	end := "[?.]*"


	regexGrd := pre
	regexLaz := pre + "?"
	for i, d := range str {
		regexGrd += fmt.Sprintf(dam, d)
		regexLaz += fmt.Sprintf(dam, d)
		if i < len(str) - 1 {
			regexGrd += mid
			regexLaz += mid + "?"
		}
	}
	regexGrd += end + "$"
	regexLaz += end + "?$"

	return regexp.MustCompile(regexGrd), regexp.MustCompile(regexLaz)
} 

func recursive(springs string, conditions []string, regex []*regexp.Regexp, r int) (total int) {
	indexesL := regex[0].FindAllStringSubmatchIndex(springs, -1)
	indexesG := regex[1].FindAllStringSubmatchIndex(springs, -1)

	// If lazy and greedy match it means theres is only one possibility
	if len(indexesL) == 0 {
		return 0
	}

	if reflect.DeepEqual(indexesL[0], indexesG[0]) {
		return 1
	}

	var lazy, greedy []int
	var size int

	for i := 2; i < len(indexesL[0]); i += 2 {
		lazy = indexesL[0][i:i+2]
		greedy = indexesG[0][i:i+2]
		size = lazy[1] - lazy[0] // Lazy and Greedy should have the same size
		if !reflect.DeepEqual(lazy, greedy) {
			break
		}
	}

	for i := lazy[0]; i <= greedy[0]; i++ {
		// If there are unknown before the first spring match, fill with dots
		// Be greedy
		newSprings := strings.ReplaceAll(springs[:i], "?", ".")
		// Replace the first spring match with faulty springs
		newSprings += strings.Repeat("#", size)
		// Maintain the rest of map
		newSprings += springs[i+size:]
		// fmt.Println("==>", newSprings)
		total += recursive(newSprings, conditions, regex, r+1)
	}

	return total
}

func part1(lines []string) (total int) {
	for _, str := range lines {
		series := strings.Split(str, " ")
		springs := series[0]
		conditions := strings.Split(series[1], ",")

		regexG, regexL := buildRegex(conditions)
		regex := []*regexp.Regexp{regexL, regexG}

		value := recursive(springs, conditions, regex, 1)
		if value == 0 {
			continue	
		}
		total += value
	}
	return total
}

func part2(lines []string) (total int) {
	for _, str := range lines {
		total += len(str)
	}
	return total
}

func main() {
	f, err := os.ReadFile("day12/sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(f), "\n")

	fmt.Println("Part 1:", part1(lines[:len(lines)-1]))
//	fmt.Println(part2(lines))
}
