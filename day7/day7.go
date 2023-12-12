package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
  high_card int = iota
  one_pair
  two_pairs
  three_of_a_kind
  full_house
  four_of_a_kind
  five_of_a_kind
)

type Bid struct {
  card string
  value int
}

var cards map[string]int = map[string]int{
  "*": 0,
  "2": 2,
  "3": 3,
  "4": 4,
  "5": 5,
  "6": 6,
  "7": 7,
  "8": 8,
  "9": 9,
  "T": 10,
  "J": 11,
  "Q": 12,
  "K": 13,
  "A": 14,
}

func getType(hand string, jokery bool) (typ3 int) {
  count := make(map[string]int)
  for _, c := range hand {
    count[string(c)]++
  }
  jokers := 0
  if (jokery) {
    jokers = count["*"]
    delete(count, "*")
  }
  calculate: switch len(count) {
  case 0, 1:
    typ3 = five_of_a_kind
  case 2:
    for _, v := range count {
      if v+jokers == 4 {
        typ3 = four_of_a_kind
        break calculate
      }
    }
    typ3 = full_house
  case 3:
    for _, v := range count {
      if v+jokers == 3 {
        typ3 = three_of_a_kind
        break calculate
      }
    }
    typ3 = two_pairs
  case 4:
    typ3 = one_pair
  case 5:
    typ3 = high_card
}
  return typ3
}

func compareCards(a, b string, jokery bool) (int) {
  typ3a := getType(a, jokery)
  typ3b := getType(b, jokery)
  if typ3a > typ3b {
    return -1
  } else if typ3a < typ3b {
    return 1
  } else {
    for i := 0; i < len(a); i++ {
      if cards[string(a[i])] > cards[string(b[i])] {
        return -1
      } else if cards[string(a[i])] < cards[string(b[i])] {
        return 1
      }
    }
  }
  return 0
}

func calcHand(lines []string, jokery bool) (total int) {
  bids := make([]Bid, len(lines)-1)
  for i, line := range lines[:len(lines)-1] {
    values := strings.Split(line, " ")[0:2]
    bids[i].card = values[0]
    if (jokery) {
      bids[i].card = strings.Replace(bids[i].card, "J", "*", -1)
    }
    bids[i].value, _ = strconv.Atoi(values[1])
  }

  sort.SliceStable(bids, func(i, j int) bool {
    return compareCards(bids[i].card, bids[j].card, true) == 1
  })

  for i, bid := range bids {
    total += (bid.value * (i + 1))
  }
  return total
}

func main() {
  f, err := os.ReadFile("day7/input.txt")

  if err != nil {
		log.Fatal(err)
	}

  lines := strings.Split(string(f), "\n") 
	fmt.Println(calcHand(lines, false))
	fmt.Println(calcHand(lines, true))
}
