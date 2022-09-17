package main

import "fmt"

func sum(array []int) int {
	count := 0
	total := 0

	for _, v := range array {

		count = count + 1
		if count <= 11 {
			total = total + v
		} else {
			total = -1
			//return -1
		}
	}
	return total
}

func main() {
	arr := []int{11, 13, 101, 14, 33, 141, 1, 1, 1, 1, 1, 1}
	//arr := []int{11, 13, 101, 14, 33, 141}
	fmt.Print("Result :", sum(arr))
}
