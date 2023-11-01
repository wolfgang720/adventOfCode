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
	badStrs := [...]string{"ab", "cd", "pq", "xy"}
	vowels := "aeiou"

	for _, line := range lines {
		if line == "" {
			continue
		}
		var (
			prev      rune
			vowelsCnt uint8
			doubles   bool
			hasBadStr bool
		)

		for _, c := range line {
			if strings.ContainsRune(vowels, c) {
				vowelsCnt += 1
			}
			if c == prev {
				doubles = true
			}
			prev = c
		}
		for _, badStr := range badStrs {
			if strings.Contains(line, badStr) {
				hasBadStr = true
				break
			}
		}
		if hasBadStr || !doubles || vowelsCnt < 3 {
			naughtyCnt += 1
			fmt.Println(hasBadStr, doubles, vowelsCnt)
		} else {
			goodCnt += 1
		}

	}

	fmt.Println("Good strings: ", goodCnt)
	fmt.Println("Naughty strings: ", naughtyCnt)
}
