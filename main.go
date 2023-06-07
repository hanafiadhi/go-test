package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var numFamilies int
	fmt.Print("Input the number of families: ")
	fmt.Scanln(&numFamilies)

	fmt.Print("Input the number of members in the family (separated by a space): ")
	var input string
	fmt.Scanln(&input)

	members := strings.Split(input, " ")

	if len(members) != numFamilies {
		fmt.Println("Input must be equal to the count of families.")
		return
	}

	totalPassengers := 0
	for _, member := range members {
		var numMembers int
		fmt.Sscanf(member, "%d", &numMembers)
		totalPassengers += numMembers
	}

	minBuses := int(math.Ceil(float64(totalPassengers) / 4))
	fmt.Println("Minimum buses required:", minBuses)
}
