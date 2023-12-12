package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"github.com/rodgco/aoc2023-go/util"
)

type Block struct {
  title string
  changes []Shift
}

type Shift struct {
  start int
  end int
  shift int
}

type Item struct {
  name string
  start int
  end int
}

func createSeedLine(str string) (items []Item) {
  seeds := strings.Split(str, " ")[1:]
  items = make([]Item, len(seeds))

  for i, seed := range seeds {
    initial, _ := strconv.Atoi(seed)
    length := 1
    items[i] = Item{"seed", initial, initial+length}
  }
  return items
}

func createSeedRange(str string) (items []Item) {
  tuples := regexp.MustCompile(`(\d+\s{1}\d+)`).FindAllStringSubmatch(str, -1)
  items = make([]Item, len(tuples))

  for i, tuple := range tuples {
    values := strings.Split(tuple[0], " ")
    initial, _ := strconv.Atoi(values[0])
    length, _ := strconv.Atoi(values[1])
    items[i] = Item{"seed", initial, initial+length}
  }
  return items
}

func sortAndConcatenateSeeds(seeds []Item) ([]Item) {
  // Let's sort the seeds, so adjacent seeds are next to each other
  sort.Slice(seeds, func(i, j int) bool { 
    return seeds[i].start < seeds[j].start
  })
  for i := int(0); i < len(seeds)-1; i++ {
    if seeds[i].end == seeds[i+1].start {
      // First extend the current seed
      seeds[i].end = seeds[i+1].end
      // Then remove the next seed
      seeds = append(seeds[:i+1], seeds[i+2:]...)
      // Let's do this one again
      i -= 1
    }
  }
  return seeds
}

func getBlock(input string) (Block) {
  lines := strings.Split(input, "\n")
  block := Block{strings.Split(lines[0], " ")[0], make([]Shift, 0)}

  for _, line := range lines[1:] {
    parts := strings.Split(line, " ")
    if len(parts) != 3 {
      continue
    }
    dest, _ := strconv.Atoi(parts[0])
    src, _ := strconv.Atoi(parts[1])
    size, _ := strconv.Atoi(parts[2])
    // Range -> start, end, shift
    block.changes = append(block.changes, Shift{src, src+size, dest-src})
  }
  return block
}

func shiftItem(to string, item Item, changes []Shift) (newItem []Item) {
  // Track items that are not shifted
  shifted := false
  for _, change := range changes {
    // [item.start                                item.end)
    //          [change.start    change.end)
    // [BEFORE )[INTER                     )[AFTER        )
    inter := Item{to, util.Max(item.start, change.start), util.Min(item.end, change.end)}
    hasInter := inter.end > inter.start

    // fmt.Println("\tChange:", change)
    // fmt.Println("\t\tInter:", hasInter, inter)

    if hasInter {
      shifted = true

      newItem = append(newItem, Item{to, inter.start+change.shift, inter.end+change.shift})

      before := Item{to, item.start, util.Min(item.end, change.start)}
      after := Item{to, util.Max(item.start, change.end), item.end}

      hasBefore := before.start < change.start
      hasAfter := after.end > change.end

      // fmt.Println("\t\tBefore:", hasBefore, before)
      // fmt.Println("\t\tAfter:", hasAfter, after)

      if hasBefore {
        newItem = append(newItem, Item{to, before.start, before.end})
      }
      if hasAfter {
        newItem = append(newItem, Item{to, after.start, after.end})
      }
      break
    }
  }
  if !shifted {
    // fmt.Println("\tNot reranged")
    item.name = to
    newItem = append(newItem, item)
  }
  return
}

func processItems(items []Item, blocks []Block) (total int) {
  for _, block := range blocks {

    to := strings.Split(block.title, "-")[2]

    newItems := make([]Item, 0)
    for len(items) > 0 {
      // Pop the last item
      item := items[len(items)-1]
      items = items[:len(items)-1]

      // fmt.Println("Item:", item)
      newItem := shiftItem(to, item, block.changes)
      newItems = append(newItems, newItem...)
    }
    items = newItems
  }
  minLocation := math.MaxInt
  for _, seed := range items {
    minLocation = util.Min(seed.start, minLocation)
  }
  return minLocation
}

func main() {
  f, err := os.ReadFile("day5/input.txt")

  if err != nil {
		log.Fatal(err)
	}

  input := strings.Split(string(f), "\n\n")

  itemsPart1 := createSeedLine(input[0])
  itemsPart2 := createSeedRange(input[0])

  blocks := make([]Block, len(input)-1)
  for i, line := range input[1:] {
    blocks[i] = getBlock(line)
  }

  fmt.Println("Part1:", processItems(itemsPart1, blocks))
  fmt.Println("Part2:", processItems(itemsPart2, blocks))
}
