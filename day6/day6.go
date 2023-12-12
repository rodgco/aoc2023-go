package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func bhaskara(a, b, c float64) (x1, x2 float64) {
  delta := math.Pow(b, 2) - 4 * a * c
  x1 = (-b + math.Sqrt(delta)) / (2 * a)
  x2 = (-b - math.Sqrt(delta)) / (2 * a)
  return x1, x2
}

type Races struct{
  time int
  distance int
}

func calcRaces(races []Races) (total int) {
  total = 1
  for _, race := range races {
    x1, x2 := bhaskara(-1, float64(race.time), float64(-1*race.distance))
    beat := int(math.Ceil(x2) - math.Trunc(x1) - 1)
    total *= beat
  }
  return total
}

func part1(lines []string) (total int) {
  re := regexp.MustCompile(`(\d+)`)
  times := re.FindAllString(lines[0], -1)
  distances := re.FindAllString(lines[1], -1)

  races := make([]Races, len(times))
  for i, _ := range times {
    races[i].time, _ = strconv.Atoi(times[i])
    races[i].distance, _ = strconv.Atoi(distances[i])
  }

  return calcRaces(races)
}

func part2(lines []string) (total int) {
  re := regexp.MustCompile(`(\d+)`)
  times := re.FindAllString(lines[0], -1)
  distances := re.FindAllString(lines[1], -1)

  var timeString string 
  var distanceString string
  for i, _ := range times {
    timeString += times[i]
    distanceString += distances[i]
  }

  time, _ := strconv.Atoi(timeString)
  distance, _ := strconv.Atoi(distanceString)

  races := make([]Races, 1)
  races[0] = Races{ time, distance }

  return calcRaces(races)
}

func main() {
  f, err := os.ReadFile("day6/input.txt")
  if err != nil {
		log.Fatal(err)
	}
  lines := strings.Split(string(f), "\n") 

	fmt.Println("Part 1", part1(lines))
	fmt.Println("Part 2", part2(lines))
}
