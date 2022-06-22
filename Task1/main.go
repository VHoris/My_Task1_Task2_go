package main

import "fmt"

func main() {

	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int
	var isNeed bool

	result = arr[:1]
	for _, v := range arr {
		isNeed = false
		for _, k := range result {
			if v == k {
				isNeed = true
			}
		}
		if !isNeed {
			result = append(result, v)
		}
	}
	fmt.Printf("%v\n", result)
}
