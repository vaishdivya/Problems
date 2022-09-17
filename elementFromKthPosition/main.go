package main

import "fmt"

func main() {
	cityList := []string{"Mumbai", "Delhi", "Australia", "Nigeria", "USA", "London", "Canada"}
	input := 2
	city := cityList[input-1:]
	fmt.Println(city)
}
