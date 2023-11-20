package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
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

	var totalCntMem uint64 = 0
	var totalCntChars uint64 = 0

	for _, line := range lines {
		lineCntMem := uint64(utf8.RuneCountInString(line))
		totalCntMem += lineCntMem

		strParts := strings.Split(line, "\\")
		backslashCnt := 0
		for _, strPart := range strParts {
			if backslashCnt > 0 {
				// double \
				totalCntChars += uint64(utf8.RuneCountInString(strPart)) + 1
				backslashCnt = 0
				continue
			}
			if len(strPart) == 0 {
				// first part of double \
				backslashCnt += 1
				continue
			}
			switch strPart[0] {
			case '"':
				totalCntChars += uint64(utf8.RuneCountInString(strPart))
			case 'x':
				totalCntChars += uint64(utf8.RuneCountInString(strPart)) - 2
			default:
				fmt.Println("error", strPart)
			}
		}
	}

	// do not count enclosing double quotes
	totalCntChars -= uint64(len(lines) * 2)

	fmt.Println("total:", totalCntMem, totalCntChars, "diff:", totalCntMem-totalCntChars)

}
