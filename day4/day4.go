package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"golang.org/x/exp/slices"
)

func calculateMatches(winning [10]int, str string) (points int) {
  matches := 0
  for i := 0; i < 25; i++ {
    n, _ := strconv.Atoi(strings.TrimSpace(str[42+i*3:44+i*3]))
    if slices.Contains(winning[:], n) {
      matches += 1
    }
  }
  return matches
}

func part1(lines []string) (total int) {
  for _, str := range lines {
    if len(str) < 1 {
      continue
    }
    winning := [10]int{}
    for i := 0; i < 10; i++ {
        winning[i], _ = strconv.Atoi(strings.TrimSpace(str[10+i*3:12+i*3]))
    }
    matches := calculateMatches(winning, str)
    points := int(math.Pow(2, float64(matches-1)))
    total += points
  }
  return total
}

func part2(lines []string) (total int) {
  cardsCount := make(map[int]int, len(lines))
  for index, str := range lines {
    if len(str) < 1 {
      continue
    }
    winning := [10]int{}
    for i := 0; i < 10; i++ {
        winning[i], _ = strconv.Atoi(strings.TrimSpace(str[10+i*3:12+i*3]))
    }
    // The original card grants 1 copy of itself
    cardsCount[index] += 1
    matches := calculateMatches(winning, str)
    for i := index+1; i < index+1+matches; i++ {
      // The original card plus it's copies grant copies of matching cards
      cardsCount[i] += cardsCount[index]
    }
    total += cardsCount[index]
  }
  return total
}

func main() {
  f, err := os.ReadFile("day4/input.txt")

  if err != nil {
		log.Fatal(err)
	}

  lines := strings.Split(string(f), "\n") 
	fmt.Println(part1(lines))

	fmt.Println(part2(lines))
}
