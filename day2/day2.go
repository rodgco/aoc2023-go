package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"regexp"
)

type Cube struct {
	red, green, blue int
}

var cubesMax = map[string]int{"red": 12, "green": 13, "blue": 14}

func part1(s *bufio.Scanner) (total int) {
	for s.Scan() {
		str := s.Text()
		parts := strings.Split(str, ":")
		game := parts[0]
		id, _ := strconv.Atoi(game[5:])
		rounds := strings.Split(parts[1], ";")
		// fmt.Println("Game: ", id)
		Rounds:
		for _, round := range rounds {
			cubes := map[string]int{"red": 0, "green": 0, "blue": 0}
			outcomes := regexp.MustCompile(`(\d*)\s(green|blue|red)`).FindAllStringSubmatch(round, -1)
			for _, outcome := range outcomes {
				result := strings.Split(outcome[0], " ")
				count, _ := strconv.Atoi(result[0])
				color := result[1]
				cubes[color] += count
			}
			for color, count := range cubesMax {
				if cubes[color] > count {
					total += id
					break Rounds
				}
			}
		}
	}
	// 5050 = Gauss sum of all numbers between 1-100
	return 5050 - total
}

func part2(s *bufio.Scanner) (total int) {
	for s.Scan() {
		str := s.Text()
		parts := strings.Split(str, ":")
		game := parts[0]
		id, _ := strconv.Atoi(game[5:])
		rounds := strings.Split(parts[1], ";")
		cubes := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, round := range rounds {
			outcomes := regexp.MustCompile(`(\d*)\s(green|blue|red)`).FindAllStringSubmatch(round, -1)
			for _, outcome := range outcomes {
				result := strings.Split(outcome[0], " ")
				count, _ := strconv.Atoi(result[0])
				color := result[1]
				if cubes[color] < count {
					cubes[color] = count
				}
			}
		}
		power := 1
		for color, _ := range cubes {
			power *= cubes[color]
		}
		total += power
	}
	return total
}

func main() {
	f, err := os.Open("day2/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
		}(f)

	s := bufio.NewScanner(f)
	fmt.Println(part1(s))

	_, err = f.Seek(0, 0)
	if err != nil {
		return
	}
	s = bufio.NewScanner(f)
	fmt.Println(part2(s))

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
