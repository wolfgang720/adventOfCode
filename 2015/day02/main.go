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

	dimensions := make([][3]uint64, len(lines))
	var (
		paperArea uint64
		ribbonLen uint64
	)

	for i, line := range lines {
		lwh := strings.Split(line, "x")
		if len(lwh) < 3 {
			continue
		}
		l, _ := strconv.ParseUint(lwh[0], 10, 64)
		w, _ := strconv.ParseUint(lwh[1], 10, 64)
		h, _ := strconv.ParseUint(lwh[2], 10, 64)

		dimensions[i] = [3]uint64{l, w, h}
		smallest := min(l*w, w*h, l*h)

		paperArea += 2*(l*w+w*h+l*h) + smallest

		volume := l * w * h
		smallesPerimeter := 2 * min(l+w, w+h, l+h)
		ribbonLen += volume + smallesPerimeter
	}

	fmt.Println("Paper to order: ", paperArea)
	fmt.Println("Ribbon to order: ", ribbonLen)
}
