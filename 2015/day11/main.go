package main

import (
	"fmt"
	"os"
	"slices"
)

func readInputFile() (string, error) {
	b, err := os.ReadFile("./input.txt") // just pass the file name
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func calcNextPassword(password []rune) []rune {
	// count up new password

	newpassword := []rune{}
	carry := true
	for i := range password {
		r := password[len(password)-1-i]
		if !carry {
			newpassword = append(newpassword, r)
			continue
		}
		if r == 'z' {
			carry = true
			newpassword = append(newpassword, 'a')
		} else {
			carry = false
			newpassword = append(newpassword, r+1)
		}
	}
	slices.Reverse(newpassword)
	return newpassword
}

func validatePassword(password []rune) bool {
	if slices.Contains(password, 'i') || slices.Contains(password, 'o') || slices.Contains(password, 'l') {
		return false
	}
	pairsIdx := []uint{}
	straitLen := 1
	maxStraitLen := 1
	prev := password[0]
	for i, r := range password[1:] {
		if r == prev {
			var lastPairsIdx uint
			if len(pairsIdx) > 0 {
				lastPairsIdx = pairsIdx[len(pairsIdx)-1]
			}
			if uint(i+1)-lastPairsIdx > 1 {
				pairsIdx = append(pairsIdx, uint(i+1))
			}
		}
		if r == prev+1 {
			straitLen += 1
			if straitLen > maxStraitLen {
				maxStraitLen = straitLen
			}
		} else {
			straitLen = 1
		}
		prev = r
	}
	if len(pairsIdx) < 2 || maxStraitLen < 3 {
		return false
	}
	return true
}

func main() {
	password := []rune("hepxcrrq")
	pwdCnt := 0
	for i := 0; pwdCnt < 2; i++ {
		password = calcNextPassword(password)
		if validatePassword(password) {
			fmt.Println("newpassword:", string(password), "index:", i)
			pwdCnt += 1
		}
	}

}
