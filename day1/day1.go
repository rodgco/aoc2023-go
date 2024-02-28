package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
	"regexp"
)

func calibrate(str string) (calibration int) {
	firstIndex := strings.IndexAny(str, "123456789")
	firstDigit, _ := strconv.Atoi(string(str[firstIndex]))

	lastIndex := strings.LastIndexAny(str, "123456789")
	lastDigit, _ := strconv.Atoi(string(str[lastIndex]))

	calibration = firstDigit*10 + lastDigit
	return
}

func part1(s *bufio.Scanner) (total int) {
	for s.Scan() {
		str := s.Text()
		calibration := calibrate(str)
		total += calibration
	}
	return
}

func stringToDigit(str string) (digit int) {
	switch str {
	case "one", "1":
		digit = 1
	case "two", "2":
		digit = 2
	case "three", "3":
		digit = 3
	case "four", "4":
		digit = 4
	case "five", "5":
		digit = 5
	case "six", "6":
		digit = 6
	case "seven", "7":
		digit = 7
	case "eight", "8":
		digit = 8
	case "nine", "9":
		digit = 9
	}
	return
}

func part2(s *bufio.Scanner) (total int) {
	var reg *regexp.Regexp
	var matches []string
	var firstDigit, lastDigit int

	for s.Scan() {
		str := s.Text()

		reg = regexp.MustCompile(`^.*?([0-9]|one|two|three|four|five|six|seven|eight|nine)`)
		matches = reg.FindStringSubmatch(str)
		firstDigit = stringToDigit(matches[1])

		reg = regexp.MustCompile(`^.*([0-9]|one|two|three|four|five|six|seven|eight|nine).*?$`)
		matches = reg.FindStringSubmatch(str)
		lastDigit = stringToDigit(matches[1])

		calibration := firstDigit*10 + lastDigit

		total += calibration
	}
	return
}

func main() {
	f, err := os.Open("day1/input.txt")

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
