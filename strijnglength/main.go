package main

import "fmt"

func main() {
	cities := []string{"Mumbai", "Hyderabad", "Calicut", "Chennai"}
	targetlength := 9
	count := 0
	for _, value := range cities {
		if len(value) >= targetlength {
			count = count + 1
		}
	}

	fmt.Println(count)
}
