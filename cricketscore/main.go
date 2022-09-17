package main

import "fmt"

func sum(array []int) int {
	total := 0

	if len(array) > 11 {
		return -1
	}

	for _, v := range array {
		total = total + v
	}
	return total
}

func main() {
	arr := []int{11, 13, 101, 14, 33, 141, 1, 1, 1, 1, 1, 1}
	//arr := []int{11, 13, 101, 14, 33, 141}
	fmt.Print("Result :", sum(arr))
}
