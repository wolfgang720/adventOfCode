package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rectangle = struct {
	x1, y1, x2, y2 uint16
}

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

	grid := [1000][1000]bool{}
	var onCnt uint64

	for lnum, line := range lines {
		action, rectangle := parseInstruction(line)
		fmt.Println(action, " | ", rectangle)

		f := func(in bool) bool {
			fmt.Println(lnum, "Error reading action", action)
			return in
		}
		switch action {
		case "turn on":
			f = func(in bool) bool { return true }
		case "turn off":
			f = func(in bool) bool { return false }
		case "toggle":
			f = func(in bool) bool { return !in }
		}

		var x1, y1, x2, y2 uint16
		if rectangle.y1 > rectangle.y2 {
			y2, y1 = rectangle.y1, rectangle.y2
		} else {
			y1, y2 = rectangle.y1, rectangle.y2
		}
		if rectangle.x1 > rectangle.x2 {
			x2, x1 = rectangle.x1, rectangle.x2
		} else {
			x1, x2 = rectangle.x1, rectangle.x2
		}

		for row := y1; row <= y2; row++ {
			for col := x1; col <= x2; col++ {
				in := grid[row][col]
				out := f(in)
				grid[row][col] = out

				if !in && out {
					onCnt += 1
				}
				if in && !out {
					onCnt -= 1
				}
			}
		}

	}
	fmt.Println("on:", onCnt)

}

func parseCoordinates(coords string) (uint16, uint16) {
	_coords := strings.Split(coords, ",")
	x, _ := strconv.ParseUint(_coords[0], 10, 16)
	y, _ := strconv.ParseUint(_coords[1], 10, 16)
	return uint16(x), uint16(y)
}

func parseInstruction(instr string) (action string, rectangle Rectangle) {
	parts := strings.Split(instr, " ")
	x2, y2 := parseCoordinates(parts[len(parts)-1])
	x1, y1 := parseCoordinates(parts[len(parts)-3])
	action = strings.Join(parts[0:len(parts)-3], " ")
	return action, Rectangle{x1, y1, x2, y2}
}
