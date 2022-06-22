package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	input := "1 9 3 4 -5"
	var result string
	var inputSliceInt []int

	inputWithoutSpaces := strings.Split(input, " ")

	for _, z := range inputWithoutSpaces {
		if str, err := strconv.Atoi(z); err == nil {
			inputSliceInt = append(inputSliceInt, str)
		}
	}

	minInt := inputSliceInt[0]
	maxInt := inputSliceInt[0]
	for _, z := range inputSliceInt {
		if minInt > z {
			minInt = z
		}
		if maxInt < z {
			maxInt = z
		}
	}

	result = strconv.Itoa(maxInt) + " " + strconv.Itoa(minInt)
	fmt.Printf("%v\n", result)

}
