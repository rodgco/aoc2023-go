package main

import (
  "fmt"
  "bufio"
  "log"
  "os"
  "regexp"
  "strconv"
)

func part1(s *bufio.Scanner) (total int) {
  var firstDigit, lastDigit int
	for s.Scan() {
		str := s.Text()

    reg := regexp.MustCompile(`^.*?(\d{1}).*(\d{1}).*$|^.*(\d{1}).*$`)
    matches := reg.FindStringSubmatch(str)

    if matches[1] == "" {
      firstDigit, _ = strconv.Atoi(matches[3])
      lastDigit, _ = strconv.Atoi(matches[3])
    } else {
      firstDigit, _ = strconv.Atoi(matches[1])
      lastDigit, _ = strconv.Atoi(matches[2])
    }

    calibration := firstDigit*10 + lastDigit

    total += calibration
  }
  return total
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
  return digit
}
// [0-9]|one|two|three|four|five|six|seven|eight|nine
func part2(s *bufio.Scanner) (total int) {
  var firstDigit, lastDigit int
	for s.Scan() {
		str := s.Text()

    reg := regexp.MustCompile(`^.*?([0-9]|one|two|three|four|five|six|seven|eight|nine).*([0-9]|one|two|three|four|five|six|seven|eight|nine).*$|^.*([0-9]|one|two|three|four|five|six|seven|eight|nine).*$`)
    matches := reg.FindStringSubmatch(str)

    if matches[1] == "" {
      firstDigit = stringToDigit(matches[3])
      lastDigit = stringToDigit(matches[3])
    } else {
      firstDigit = stringToDigit(matches[1])
      lastDigit = stringToDigit(matches[2])
    }

    calibration := firstDigit*10 + lastDigit

    total += calibration
  }
  return total
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
