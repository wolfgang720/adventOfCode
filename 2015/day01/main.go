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
	ups := strings.Count(content, "(")
	downs := strings.Count(content, ")")

	fmt.Println("Floor: ", ups-downs)

	var floor int
	var up rune = '('
	// var down rune = ')'
	for i, c := range content {
		if c == up {
			floor += 1
		} else {
			floor -= 1
		}
		if floor < 0 {
			fmt.Println("basement: ", i+1)
			break
		}
	}
}
