package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

	type Aunt map[string]int

	aunts := map[uint]Aunt{}
	for i, line := range lines {
		_p := strings.Split(line, " ")
		comp1 := strings.TrimSuffix(_p[2], ":")
		val1, _ := strconv.Atoi(strings.TrimSuffix(_p[3], ","))
		comp2 := strings.TrimSuffix(_p[4], ":")
		val2, _ := strconv.Atoi(strings.TrimSuffix(_p[5], ","))
		comp3 := strings.TrimSuffix(_p[6], ":")
		val3, _ := strconv.Atoi(strings.TrimSuffix(_p[7], ","))
		newAunt := Aunt{comp1: val1, comp2: val2, comp3: val3}
		aunts[uint(i)+1] = newAunt
	}

	compounds := map[string]int{"children": 3, "cats": 7, "samoyeds": 2, "pomeranians": 3, "akitas": 0, "vizslas": 0, "goldfish": 5, "trees": 3, "cars": 2, "perfumes": 1}

auntsloop:
	for idx, aunt := range aunts {

		for auntcomp, val := range aunt {
			if val != compounds[auntcomp] {
				continue auntsloop
			}
		}
		fmt.Println(idx, aunt)
	}

}
