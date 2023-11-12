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

type outwire = string
type Gate = struct {
	gateStrs    []string
	output      uint16
	outputValid bool
}
type Circuit = map[outwire]Gate

func main() {
	content, err := readInputFile()
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	lines := strings.Split(content, "\n")

	// wire up circuit
	var circuit = Circuit{}
	for _, line := range lines {
		_x := strings.Split(line, " -> ")[0:2]
		circuit[_x[1]] = Gate{gateStrs: strings.Split(_x[0], " ")}
	}

	for key, value := range circuit {
		fmt.Println(key, value)
	}

	outputOfA := calcGateOutput(&circuit, circuit["a"])
	fmt.Println("Output of wir a:", outputOfA)
}

type BitOperation = func(in uint16) (out uint16)

func calcGateOutput(circuit *Circuit, gate Gate) uint16 {

	var ins []string
	var op string

	gateStrs := gate.gateStrs

	if len(gateStrs) == 0 {
		fmt.Println("len 0:", gateStrs)
		return 0
	}

	if len(gateStrs) == 1 {
		ins = gateStrs
		op = "EQ"
	} else if gateStrs[0] == "NOT" {
		ins = []string{gateStrs[1]}
		op = "NOT"
	} else if gateStrs[1] == "AND" || gateStrs[1] == "OR" {
		ins = []string{gateStrs[0], gateStrs[2]}
		op = gateStrs[1]
	} else if gateStrs[1] == "LSHIFT" || gateStrs[1] == "RSHIFT" {
		ins = []string{gateStrs[0], gateStrs[2]}
		op = gateStrs[1]
	}

	var inputs []uint16
	for _, input := range ins {
		num, err := strconv.ParseUint(input, 10, 16)
		if err == nil {
			inputs = append(inputs, uint16(num))
			continue
		}
		fmt.Println("fetching for ", input, "with:", (*circuit)[input])
		gateOnInput := (*circuit)[input]
		if gateOnInput.outputValid {
			inputs = append(inputs, gateOnInput.output)
			continue
		}
		gateOutput := calcGateOutput(circuit, gateOnInput)
		inputs = append(inputs, gateOutput)
		(*circuit)[input] = Gate{gateStrs: gateOnInput.gateStrs, output: gateOutput, outputValid: true}
	}

	switch op {
	case "EQ":
		return inputs[0]
	case "NOT":
		return ^inputs[0]
	case "AND":
		return inputs[0] & inputs[1]
	case "OR":
		return inputs[0] | inputs[1]
	case "LSHIFT":
		return inputs[0] << inputs[1]
	case "RSHIFT":
		return inputs[0] >> inputs[1]
	default:
		{
			fmt.Println("ERROR:", gateStrs)
			return 0
		}
	}

}
