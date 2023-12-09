package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	combin "gonum.org/v1/gonum/stat/combin"
)

func readInputFile() (string, error) {
	b, err := os.ReadFile("./input.txt") // just pass the file name
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type Pair struct {
	who  string
	with string
}

func newSortedPair(p1 string, p2 string) Pair {
	if strings.Compare(p1, p2) == -1 {
		return Pair{p2, p1}
	} else {
		return Pair{p1, p2}
	}
}

func main() {
	content, err := readInputFile()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	lines := strings.Split(content, "\n")

	happinesses := make(map[Pair]int)
	var people []string
	for _, line := range lines {
		elems := strings.Split(line, " ")

		happinum, _ := strconv.Atoi(elems[3])
		if elems[2] == "lose" {
			happinum = -happinum
		}
		pair := newSortedPair(elems[0], strings.TrimSuffix(elems[10], "."))

		if h, ok := happinesses[pair]; ok {
			happinesses[pair] = h + happinum
		} else {
			happinesses[pair] = happinum
		}
		if !slices.Contains(people, pair.who) {
			people = append(people, pair.who)
		}
		if !slices.Contains(people, pair.with) {
			people = append(people, pair.with)
		}
	}

	maxHappiness := 0
	for _, arrangement := range combin.Permutations(len(people), len(people)) {

		arrHappiness := 0
		prevPIdx := arrangement[0]
		for _, pIdx := range arrangement[1:] {
			pair := newSortedPair(people[prevPIdx], people[pIdx])
			arrHappiness += happinesses[pair]
			prevPIdx = pIdx
		}
		lastpair := newSortedPair(people[arrangement[len(arrangement)-1]], people[arrangement[0]])
		arrHappiness += happinesses[lastpair]

		if arrHappiness > maxHappiness {
			maxHappiness = arrHappiness
		}
	}

	fmt.Println(maxHappiness)
}
