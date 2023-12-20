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

type Ingredient struct {
	capacity, durabiltiy, flavour, texture, calories int
}

func main() {
	content, err := readInputFile()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	var ings []Ingredient
	for _, line := range strings.Split(content, "\n") {
		_p := strings.Split(line, " ")
		capa, _ := strconv.Atoi(strings.TrimSuffix(_p[2], ","))
		dura, _ := strconv.Atoi(strings.TrimSuffix(_p[4], ","))
		flav, _ := strconv.Atoi(strings.TrimSuffix(_p[6], ","))
		text, _ := strconv.Atoi(strings.TrimSuffix(_p[8], ","))
		cals, _ := strconv.Atoi(strings.TrimSuffix(_p[10], ","))
		ings = append(ings, Ingredient{capacity: capa, durabiltiy: dura, flavour: flav, texture: text, calories: cals})
	}
	fmt.Println(ings)

	maxScore := uint64(0)
	for i1 := 0; i1 <= 100; i1++ {
		for i2 := 0; i2 <= 100-i1; i2++ {
			for i3 := 0; i3 <= 100-i1-i2; i3++ {
				for i4 := 0; i4 <= 100-i1-i2-i3; i4++ {
					capa := max(0, i1*ings[0].capacity+i2*ings[1].capacity+i3*ings[2].capacity+i4*ings[3].capacity)
					dura := max(0, i1*ings[0].durabiltiy+i2*ings[1].durabiltiy+i3*ings[2].durabiltiy+i4*ings[3].durabiltiy)
					flav := max(0, i1*ings[0].flavour+i2*ings[1].flavour+i3*ings[2].flavour+i4*ings[3].flavour)
					text := max(0, i1*ings[0].texture+i2*ings[1].texture+i3*ings[2].texture+i4*ings[3].texture)
					score := uint64(capa * dura * flav * text)
					if score > maxScore {
						maxScore = score
					}
				}
			}
		}
	}

	fmt.Println("max Score:", maxScore)
}
