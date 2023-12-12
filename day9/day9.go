package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func extrapolate(values []int) (int, int) {
  sum := 0
  for _, value := range values {
    sum += value
  }
  if sum == 0 { return 0, 0 }

  next := make([]int, len(values)-1)
  for i := range values[:len(values)-1] {
    value := values[i+1] - values[i]
    next[i] = value
  }

  before, after := extrapolate(next)
  return values[0] - before, values[len(values)-1] + after
}

func main() {
  f, err := os.ReadFile("day9/input.txt")

  if err != nil {
		log.Fatal(err)
	}

  lines := strings.Split(string(f), "\n") 

  report := make([][]int, 0)

  // TODO: refactor this
  for i, line := range lines[:len(lines)-1] {
    sValues := strings.Split(line, " ")
    report = append(report, make([]int, 0))
    for _, sValue := range sValues {
      value, _ := strconv.Atoi(sValue)
      report[i] = append(report[i], value)
    }
  }

  sumBefore, sumAfter := 0, 0
  for _, values := range report {
    before, after := extrapolate(values)
    sumBefore += before
    sumAfter += after
  }

	fmt.Println(sumAfter)
	fmt.Println(sumBefore)
}
