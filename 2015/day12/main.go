package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func readInputFile() (string, error) {
	b, err := os.ReadFile("./input.txt") // just pass the file name
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func parseChildren(input interface{}) int {
	var numbersCnt int
	switch _input := input.(type) {
	case []interface{}:
		fmt.Println("array")
		for _, elem := range _input {
			numbersCnt += parseChildren(elem)
		}
		return numbersCnt
	case map[string]interface{}:
		fmt.Println("object")
		for _, v := range _input {
			numbersCnt += parseChildren(v)
			if v == "red" {
				return 0
			}
		}
		return numbersCnt
	case string:
		fmt.Println("string")
		numbersCnt = 0
		return numbersCnt
	case float64:
		fmt.Println("int")
		numbersCnt = int(_input)
		return numbersCnt
	default:
		fmt.Println("other", reflect.TypeOf(_input))
	}
	return numbersCnt
}

func main() {
	content, err := readInputFile()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	var data interface{}
	json.Unmarshal([]byte(content), &data)
	result := parseChildren(data)
	fmt.Println(result)
}
