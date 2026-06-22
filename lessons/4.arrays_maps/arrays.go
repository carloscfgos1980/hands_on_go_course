package main

import "fmt"

func UniqueElements() {
	intSlice := []int{1, 5, 5, 5, 5, 7, 8, 6, 6, 6}
	fmt.Println(intSlice)
	uniqueIntSlice := unique(intSlice)
	fmt.Println(uniqueIntSlice)
}
func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	uniqueElements := []int{}
	// Iterate through the slice and add unique elements to the map
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueElements = append(uniqueElements, entry)
		}
	}
	return uniqueElements
}
