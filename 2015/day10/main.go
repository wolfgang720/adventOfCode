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

func main() {

	input := []rune("3113322113")

	type NumsRecord struct {
		num rune
		cnt uint16
	}

	for its := range [40]int{} {
		// count number of subsequently equal runes
		records := []NumsRecord{}
		currentR := input[0]
		currentLen := uint16(1)
		for _, r := range input[1:] {
			if r == currentR {
				currentLen += 1
			} else {
				records = append(records, NumsRecord{num: currentR, cnt: currentLen})
				currentR = r
				currentLen = 1
			}
		}
		records = append(records, NumsRecord{num: currentR, cnt: currentLen})
		// build new num
		var newNum []rune
		for _, rec := range records {
			newNum = append(newNum, rune(rec.cnt+'0'), rec.num)
		}
		input = newNum
		fmt.Println(its, len(newNum))
	}
}
