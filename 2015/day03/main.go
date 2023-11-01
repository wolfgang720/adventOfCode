package main

import (
	"fmt"
	"os"
)

func readInputFile() (string, error) {
	b, err := os.ReadFile("./input.txt") // just pass the file name
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type Position = struct {
	x int32
	y int32
}

func main() {
	content, err := readInputFile()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	var (
		posSanta1, posSanta2 Position
	)
	houses := map[Position]uint16{{x: 0, y: 0}: 1}
	for i, direction := range content {
		var pos *Position
		if i%2 == 0 {
			pos = &posSanta1
		} else {
			pos = &posSanta2
		}

		if direction == '^' {
			pos.y += 1
		} else if direction == '>' {
			pos.x += 1
		} else if direction == '<' {
			pos.x -= 1
		} else if direction == 'v' {
			pos.y -= 1
		} else {
			fmt.Println("could not read: ", direction)
		}
		houses[*pos] += 1
	}
	fmt.Println("visited houses", len(houses))
}
