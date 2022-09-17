package main

import (
	"fmt"
	"strconv"
)

func numString(array []string, input int) []string {
	//array[0]---"Golang"
	//array[1]--"123"
	for i, value := range array {
		convertedValue, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		addNum := convertedValue + input
		array[i] = strconv.Itoa(addNum)

	}
	return array
}

func main() {
	//values := []string{"Golang", "123", "Data"}
	values := []string{"Golang", "123", "Data", "456"}
	k := 4
	fmt.Println("Result", numString(values, k))

}
