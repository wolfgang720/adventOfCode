package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
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
	fmt.Println(content)

	// var hInput []byte
	// for _, c := range "abcdef609043" {
	// 	// v, _ := strconv.ParseUint(c, 16, 8)
	// 	hInput = append(hInput)
	// }
	for i := 0; true; i++ {
		h := md5.Sum([]byte(content + strconv.Itoa(i)))
		if h[0] == 0 && h[1] == 0 && h[2] == 0 {
			fmt.Println(hex.EncodeToString(h[:]), i)
			break
		}
	}

}
