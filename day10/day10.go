package main

import (
  "fmt"
  "strings"
  "log"
  "os"
)

func part1(lines []string) (total int) {
  for _, str := range lines {
    total += len(str)
  }
  return total
}

func part2(lines []string) (total int) {
  for _, str := range lines {
    total += len(str)
  }
  return total
}

/**
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
*/

type Pos struct {
  x int
  y int
  next string
}

func findStartExits(maze [][]string, start Pos) (Pos) {
  // Using an array so it's easier to identify the exits
  response := make([]Pos, 2)
  response[0] = Pos{x: start.x, y: start.y}
  response[1] = Pos{x: start.x, y: start.y}

  // Count will be used to track the first and second exit
  count := 0
  if start.x > 0 && strings.Contains("F|7", maze[start.x-1][start.y]) {
    response[count].next = "N"
    count++
  }
  if start.y < len(maze[0]) && strings.Contains("-7J", maze[start.x][start.y+1]) {
    response[count].next = "E"
    count++
  }
  if start.x < len(maze) && strings.Contains("L|J", maze[start.x+1][start.y]) {
    response[count].next = "S"
    count++
  }
  if start.y > 0 && strings.Contains("-FL", maze[start.x][start.y-1]) {
    response[count].next = "W"
  }

  // Let's replace the start with the correct character
  switch response[0].next+response[1].next {
  case "NE", "EN":
    maze[start.x][start.y] = "L"
  case "NS", "SN":
    maze[start.x][start.y] = "|"
  case "NW", "WN":
    maze[start.x][start.y] = "J"
  case "ES", "SE":
    maze[start.x][start.y] = "F"
  case "EW", "WE":
    maze[start.x][start.y] = "-"
  case "SW", "WS":
    maze[start.x][start.y] = "7"
  }

  // Take the firs step
  walk(maze, &response[0])

  return response[0]
}

func boxize(str string) string {
  switch str {
  case "-":
    return "\u2500"
  case "|":
    return "\u2502"
  case "L":
    return "\u2514"
  case "J":
    return "\u2518"
  case "7":
    return "\u2510"
  case "F":
    return "\u250C"
  }
  return str
}

func walk(maze [][]string, pos *Pos) {

  if pos.next == "N" {
    pos.x--
  } else if pos.next == "E" {
    pos.y++
  } else if pos.next == "S" {
    pos.x++
  } else if pos.next == "W" {
    pos.y--
  }


  newPos := maze[pos.x][pos.y]

  if newPos == "L" {
    if pos.next == "S" {
      pos.next = "E"
    } else {
      pos.next = "N"
    }
  } else if newPos == "J" {
    if pos.next == "S" {
      pos.next = "W"
    } else {
      pos.next = "N"
    }
  } else if newPos == "7" {
    if pos.next == "N" {
      pos.next = "W"
    } else {
      pos.next = "S"
    }
  } else if newPos == "F" {
    if pos.next == "N" {
      pos.next = "E"
    } else {
      pos.next = "S"
    }
  }
}

func main() {
  f, err := os.ReadFile("day10/input.txt")

  if err != nil {
		log.Fatal(err)
	}

  lines := strings.Split(string(f), "\n") 

  maze  := make([][]string, len(lines)-1)
  mask  := make([][]bool, len(lines)-1)

  var start Pos

  for x, line := range lines[:len(lines)-1] {
    maze[x] = strings.Split(line, "")
    mask[x] = make([]bool, len(maze[x]))
    for y, char := range maze[x] {
      mask[x][y] = false
      if char == "S" {
        start = Pos{x: x, y: y}
      }
    }
  }

  // Let's take our firs step
  mask[start.x][start.y] = true
  pos1 := findStartExits(maze, start)

  // Part 1
  count1 := 1
  for !(pos1.x == start.x && pos1.y == start.y) {
    mask[pos1.x][pos1.y] = true
    walk(maze, &pos1)
    count1++
  }

  // Part 2
  // Ray tracing
  count2 := 0
  for x, line := range maze {
    for y, _ := range line {
      if !mask[x][y] {
        in := false
        last := ""
        for z := y+1; z < len(line); z++ {
          if mask[x][z] {
            if maze[x][z] == "|" {
              in = !in
            } else if strings.Contains("FL", maze[x][z]) { // Handle U turns and SZ like shapes
              in = !in
              last = maze[x][z]
            } else if maze[x][z] == "J" && last == "L" {
              in = !in
              last = ""
            } else if maze[x][z] == "7" && last == "F" {
              in = !in
              last = ""
            }
          }
        }
        if in {
          count2++
        }
      }
    }
  } 

	fmt.Println(count1/2)
	fmt.Println(count2)
}
