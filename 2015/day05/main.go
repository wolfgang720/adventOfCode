package main

import (
	"fmt"
	"os"
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

	var naughtyCnt, goodCnt uint16

	for _, line := range lines {
		if line == "" {
			continue
		}
		var (
			repAfterTwo bool
			hasDouble   bool
		)
		substrs := make(map[string]int)

		substrs[line[0:2]] = 1
		for i, c := range line {
			prevIdx := i - 2
			if prevIdx >= 0 && c == rune(line[prevIdx]) {
				repAfterTwo = true
			}
			if i > 0 {
				substr := line[i-1 : i+1]
				lastSubstrIdx := substrs[substr]
				if lastSubstrIdx < (i - 1) {
					if lastSubstrIdx != 0 {
						hasDouble = true
						fmt.Println(substr, lastSubstrIdx, i-1)
					} else {
						substrs[substr] = i
					}
				}
			}
		}
		if !repAfterTwo || !hasDouble {
			naughtyCnt += 1
			// fmt.Println(repAfterTwo, hasDouble)
		} else {
			goodCnt += 1
		}

	}

	fmt.Println("Good strings: ", goodCnt)
	fmt.Println("Naughty strings: ", naughtyCnt)
}
