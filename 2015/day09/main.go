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

func main() {
	content, err := readInputFile()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	lines := strings.Split(content, "\n")

	type CityPair struct {
		src, dest string
	}

	var ways = map[CityPair]int{}
	var cities = []string{}
	for _, line := range lines {
		_parts := strings.Split(line, " ")
		src := _parts[0]
		dest := _parts[2]
		dist, _ := strconv.Atoi(_parts[4])

		ways[CityPair{src, dest}] = dist
		ways[CityPair{dest, src}] = dist

		if !slices.Contains(cities, src) {
			cities = append(cities, src)
		}
		if !slices.Contains(cities, dest) {
			cities = append(cities, dest)
		}
	}

	citiesCnt := len(cities)
	var maxDist uint64
	for _, path := range combin.Permutations(citiesCnt, citiesCnt) {
		_src := cities[path[0]]
		var pathDist uint64
		for _destIdx := 1; _destIdx < citiesCnt; _destIdx++ {
			_dest := cities[path[_destIdx]]
			pathDist += uint64(ways[CityPair{_src, _dest}])
			_src = _dest
		}
		if pathDist > maxDist {
			maxDist = pathDist
		}
	}
	fmt.Println("max path:", maxDist)

}
