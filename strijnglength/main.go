package main

import "fmt"

func citiesGreaterThan(city []string, targetvalue int) int {
	count := 0
	for _, value := range city {
		if len(value) >= targetvalue {
			count = count + 1
		}
	}
	return count
}

func main() {
	cities := []string{"Mumbai", "Hyderabad", "Calicut", "Chennai"}
	targetlength := 9
	fmt.Println(citiesGreaterThan(cities, targetlength))
}
