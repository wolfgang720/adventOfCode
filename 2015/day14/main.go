package main

import (
	"fmt"
	"math"
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

type Reindeer struct {
	speed, durationSec, restSec float64
	name                        string
}

func main() {
	content, err := readInputFile()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	lines := strings.Split(content, "\n")

	const raceDurS = 2503
	// reindeer := []Reindeer{}
	maxDist := 0.0
	for _, line := range lines {
		_p := strings.Split(line, " ")
		_speed, _ := strconv.ParseFloat(_p[3], 64)
		_durS, _ := strconv.ParseFloat(_p[6], 64)
		_restS, _ := strconv.ParseFloat(_p[13], 64)
		_name := _p[0]
		// newReindeer := Reindeer{speed: _speed, durationSec: _durS, restSec: _restS, name: _name}

		sprintD := _speed * _durS
		cycleLen := _durS + _restS

		dist := math.Floor(raceDurS/cycleLen) * sprintD
		_r := math.Mod(raceDurS, cycleLen)
		if _r >= _durS {
			dist += sprintD
		} else {
			dist += _r * _speed
		}
		fmt.Println(_name, dist)
		if dist > maxDist {
			maxDist = dist
		}
		// reindeer = append(reindeer, newReindeer)
	}
	fmt.Println(maxDist)

}
