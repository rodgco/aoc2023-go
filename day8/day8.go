package main

import (
  "fmt"
  "strings"
  "log"
  "os"
)

type Element struct {
  left string
  right string
}

type Node struct {
  id string
  Z bool
}

func part1(inst []string, mapa map[string]Element) (total int) {
  id := "AAA"
  for i :=0 ; id != "ZZZ"; i = (i + 1) % len(inst) {
    total++
    instruction := inst[i]
    element := mapa[id]

    if instruction == "R" {
      id = element.right
    } else {
      id = element.left
    }
  }

  return total
}

func check(nodes []Node) (bool) {
  for _, node := range nodes {
    if !node.Z { return false }
  }
  return true
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func part2(instructions []string, elements map[string]Element) (total int) {
  nodes := make([]Node, 0)
  for key, _ := range elements {
    if key[2:3] == "A" {
      nodes = append(nodes, Node{key, false})
    }
  }

  counter := 0
  total = 1

  for i :=0 ; !check(nodes); i = (i + 1) % len(instructions) {
    instruction := instructions[i]

    counter++

    for i, _ := range nodes {
      if nodes[i].Z { continue }
      element := elements[nodes[i].id]
      if instruction == "R" {
        nodes[i].id = element.right
      } else {
        nodes[i].id = element.left
      }

      // That's the new node
      if nodes[i].id[2:3] == "Z" {
        nodes[i].Z = true
        total = LCM(total, counter)
      }
    }
  }

  return total
}

func main() {
  f, err := os.ReadFile("day8/input.txt")

  if err != nil {
		log.Fatal(err)
	}

  lines := strings.Split(string(f), "\n") 

  instructions := strings.Split(lines[0], "")
  maps := lines[2:len(lines)-1]

  elements := make(map[string]Element)
  for _, maq := range maps {
    id := maq[0:3]
    left := maq[7:10]
    right := maq[12:15]

    elements[id] = Element{left, right}
  }
  
	fmt.Println(part1(instructions, elements))
	fmt.Println(part2(instructions, elements))
}
