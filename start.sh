#!/bin/bash

export $(xargs <.env)

CURRENT=$(date +%-d)
DAY=${1:-$CURRENT}

read -d '' TEMPLATE << EOF
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

func main() {
  f, err := os.ReadFile("day${DAY}/input.txt")

  if err != nil {
		log.Fatal(err)
	}

  lines := strings.Split(string(f), "\\n") 
	fmt.Println(part1(lines))

	fmt.Println(part2(lines))
}
EOF

if [ ! -d "day$DAY" ]; then
    mkdir "day$DAY"
fi

cd "day$DAY"

echo "$TEMPLATE" > day$DAY.go

aocdl -day $DAY -year 2023 -session-cookie $SESSION_COOKIE -output input.txt
