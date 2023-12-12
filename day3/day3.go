package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
  "github.com/rodgco/aoc2023-go/util"
)

func part1(lines []string) (total int) {
  for index, str := range lines {
    // Find the "coordinates" of all numbers in the string
    numbers := regexp.MustCompile(`\d+`).FindAllStringSubmatchIndex(str, -1)
    for _, number := range numbers {
      var colStart, colEnd, rowStart, rowEnd int
      colStart = util.Max(0, number[0]-1);
      colEnd = util.Min(len(str), number[1]+1);
      rowStart = util.Max(0, index-1);
      rowEnd = util.Min(len(lines)-2, index+1);
      // Build a string of all the characters around the current number
      var text string
      for i := rowStart; i <= rowEnd; i++ {
        text += lines[i][colStart:colEnd]
      }
      // Find all the symbols in the string
      symbol := regexp.MustCompile(`[^\d\.]`).FindAllString(text, -1)
      if len(symbol) > 0 {
        // If there are symbols, we need to get the number and do the math
        value, _ := strconv.Atoi(str[number[0]:number[1]])
        total += value
      }
    }
  }
  return total
}

func part2(lines []string) (total int32) {
  // Find the "coordinates" of all numbers in the string and build a map
  numbers := make([][][]int, len(lines))
  for index, str := range lines {
    inLine := regexp.MustCompile(`\d+`).FindAllStringSubmatchIndex(str, -1)
    numbers[index] = make([][]int, len(inLine))
    for i, number := range inLine {
      numbers[index][i] = []int{number[0], number[1]-1}
    }
  }
  for index, str := range lines {
    // Find the "coordinates" of all gears in the string
    gears := regexp.MustCompile(`\*`).FindAllStringSubmatchIndex(str, -1)
    for _, gear := range gears {
      // Find the first and last rows to look at
      var rowStart, rowEnd int
      rowStart = util.Max(0, index-1);
      rowEnd = util.Min(len(lines)-2, index+1);
      ratio := 0
      add := false
      for i := rowStart; i <= rowEnd; i++ {
        for _, number := range numbers[i] {
          // If the number is "touching" the gear, we need to get the number and do the math
          if (number[0] >= gear[0]-1 && number[0] <= gear[0]+1) || (number[1] >= gear[0]-1 && number[1] <= gear[0]+1)  {
            value, _ := strconv.Atoi(lines[i][number[0]:number[1]+1])
            if ratio == 0 {
              ratio = value
            } else {
              add = true
              ratio *= value
            }
          }
        }
      }
      // Only add the ratio if we have a number to add
      if add {
        total += int32(ratio)
      }
    }
  }
  return total
}

func main() {
  f, err := os.ReadFile("day3/input.txt")

  if err != nil {
		log.Fatal(err)
	}

  lines := strings.Split(string(f), "\n") 
	fmt.Println(part1(lines))

	fmt.Println(part2(lines))
}
