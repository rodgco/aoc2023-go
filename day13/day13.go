package main

import (
	"fmt"
	"strings"
	"log"
	"os"
	//	"github.com/rodgco/aoc2023-go/util"
)

func mirrored(line, line2 string) bool {
	return line == line2
}

func rotate(block []string) []string {
	rotated := make([]string, len(block[0]))

	for i := len(block)-1; i >= 0; i-- {
		for j, char := range block[i] {
			rotated[j] += string(char)
		}
	}
	return rotated
}

func findHorizontalMirror(lines []string) int {
	main: for i, line := range lines[:len(lines)-1] {
		if mirrored(line, lines[i+1]) {
			p := i + 1 // Count of lines above
			j := i - 1 // Index counting backwards
			k := i + 2 // Index counting forward
			for j >= 0 && k < len(lines)-1 {
				if !mirrored(lines[j], lines[k]) {
					continue main
				}
				j--
				k++
			}
			fmt.Println("=> mirrored", p-1)
			return p-1
		}
	}
	return 0
}

func part1(blocks []string) (total int) {
	for i, block := range blocks {
		fmt.Println("\ntotal", total, "block", i)
		lines := strings.Split(block, "\n")
		horizontal := findHorizontalMirror(lines)
		total += 100 *  horizontal
		if horizontal > 0 {
			continue
		}
		fmt.Println("=> rotating")
		rotated := rotate(lines)
		vertical := findHorizontalMirror(rotated)
		total += vertical
		if vertical > 0 {
			continue
		}
	}

	return total
}

func part2(block []string) (total int) {
  for _, str := range block {
    total += len(str)
  }
  return total
}

func main() {
  f, err := os.ReadFile("day13/sample.txt")

  if err != nil {
		log.Fatal(err)
	}
  blocks := strings.Split(strings.Trim(string(f), "\n"), "\n\n") 
	fmt.Println(part1(blocks))
	fmt.Println(part2(blocks))
}
