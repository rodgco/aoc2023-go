package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Galaxy struct {
  id int
  x int
  y int
}

func expandRows(universe [][]string) [][]string {
  expandedRow := make([]string, len(universe[0]))
  for i := range expandedRow {
    expandedRow[i] = "X"
  }
  expandedUniverse := make([][]string, 0)
  for _, row := range universe {
    galaxies := 0
    for _, data := range row {
      if data == "#" {
        galaxies++
      }
    }
    if galaxies == 0 {
      expandedUniverse = append(expandedUniverse, expandedRow)
    } else {
      expandedUniverse = append(expandedUniverse, row)
    }
  }
  return expandedUniverse
}

func rotateUniverse(universe [][]string) [][]string {
  newUniverse := make([][]string, len(universe[0]))
  for i := range newUniverse {
    newUniverse[i] = make([]string, len(universe))
  }

  for i, row := range universe {
    for j, col := range row {
      newUniverse[j][i] = col
    }
  }
  return newUniverse
}

func printUniverse(universe [][]string) {
  for _, row := range universe {
    fmt.Println(row)
  }
}

func getGalaxies(universe [][]string, factor int) (galaxies []Galaxy) {
  idCounter := 0
  rowExpander := 0
  for i, row := range universe {
    if row[0] == "X" {
      rowExpander += factor - 1
    }
    colExpander := 0
    for j, col := range row {
      if col == "X" {
        colExpander += factor - 1
      }
      if col == "#" {
        idCounter++
        galaxies = append(galaxies, Galaxy{idCounter, rowExpander+i, colExpander+j})
      }
    }
  }
  return galaxies
}

func distance(galaxy1 Galaxy, galaxy2 Galaxy) int {
  return int(math.Abs(float64((galaxy1.x - galaxy2.x))) + math.Abs(float64((galaxy1.y - galaxy2.y))))
}

func main() {
  f, err := os.ReadFile("day11/input.txt")

  if err != nil {
		log.Fatal(err)
	}

  lines := strings.Split(string(f), "\n") 
  
  universe := make([][]string, len(lines)-1)
  for i, line := range lines[:len(lines)-1] {
    universe[i] = strings.Split(line, "")
  }

  // TODO: refactor this
  expandedUniverse := expandRows(universe)
  expandedUniverse = rotateUniverse(expandedUniverse)
  expandedUniverse = expandRows(expandedUniverse)
  expandedUniverse = rotateUniverse(expandedUniverse)
  expandedUniverse = rotateUniverse(expandedUniverse)
  expandedUniverse = rotateUniverse(expandedUniverse)

  galaxies := getGalaxies(expandedUniverse, 2)

  total := 0
  for i, galaxy1 := range galaxies[:len(galaxies)-1] {
    for _, galaxy2 := range galaxies[i+1:] {
      total += distance(galaxy1, galaxy2)
    }
  }
  fmt.Println("Part 1:", total)

  galaxiesM := getGalaxies(expandedUniverse, 1000000)

  totalM := 0
  for i, galaxy1 := range galaxiesM[:len(galaxiesM)-1] {
    for _, galaxy2 := range galaxiesM[i+1:] {
      totalM += distance(galaxy1, galaxy2)
    }
  }
	fmt.Println("Part 2:", totalM)
}
