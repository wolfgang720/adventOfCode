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

type Reindeer struct {
	speed, durationSec, restSec uint64
	name                        string
	distance                    uint64
	points                      uint64
}

func main() {
	content, err := readInputFile()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	lines := strings.Split(content, "\n")

	const raceDurS = uint64(2503)
	reindeer := []Reindeer{}

	for _, line := range lines {
		_p := strings.Split(line, " ")
		_speed, _ := strconv.ParseUint(_p[3], 10, 64)
		_durS, _ := strconv.ParseUint(_p[6], 10, 64)
		_restS, _ := strconv.ParseUint(_p[13], 10, 64)
		_name := _p[0]
		reindeer = append(reindeer, Reindeer{speed: _speed, durationSec: _durS, restSec: _restS, name: _name})
	}

	// reindeer = append(reindeer, Reindeer{name: "Dancer", speed: 16, durationSec: 11, restSec: 162})
	// reindeer = append(reindeer, Reindeer{name: "Comet", speed: 14, durationSec: 10, restSec: 127})

	for sec := uint64(0); sec < raceDurS; sec++ {
		leadingDist := uint64(0)
		for ridx, r := range reindeer {
			cycleSec := r.durationSec + r.restSec
			newDistance := r.distance
			if sec%cycleSec < r.durationSec {
				// currently running
				// fmt.Println(sec, cycleSec, sec%cycleSec, r.speed)
				newDistance = r.distance + r.speed
			}
			reindeer[ridx].distance = newDistance
			if newDistance >= leadingDist {
				leadingDist = newDistance
			}
		}
		for ridx, r := range reindeer {
			if r.distance == leadingDist {
				reindeer[ridx].points = r.points + 1
			}
		}
	}

	var mostPoints = uint64(0)
	for _, r := range reindeer[1:] {
		fmt.Println(r, r.points, mostPoints)
		if r.points > mostPoints {
			mostPoints = r.points
		}
	}

	fmt.Println(reindeer)
	fmt.Println(mostPoints)
}
