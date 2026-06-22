package main

import (
	"fmt"
	"sort"
)

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

func FindUniqueElements() {
	// Find the index of "Sandy" in the slice
	str := []string{"Sandy", "Provo", "St. George", "Salt Lake City", "Draper", "South Jordan", "Murray"}
	for i, v := range str {
		if v == "Sandy" {
			fmt.Println(i)
		}
	}
	// Find the index of "Sandy" in the slice using a map
	indexMap := make(map[string]int)
	for i, v := range str {
		indexMap[v] = i
	}
	fmt.Println(indexMap["Draper"])
	// Find the index of "Sandy" in the slice using a map and check if it exists
	if index, ok := indexMap["Sandy"]; ok {
		fmt.Println(index)
	} else {
		fmt.Println("Sandy not found")
	}
	if index, ok := indexMap["Carlos"]; ok {
		fmt.Println(index)
	} else {
		fmt.Println("Carlos not found")
	}
}

func SortSlice() {
	str := []string{"Sandy", "Provo", "St. George", "Salt Lake City", "Draper", "South Jordan", "Murray"}
	fmt.Println(str)
	sortedList := make([]string, len(str))
	copy(sortedList, str)
	sort.Strings(sortedList)
	fmt.Println(sortedList)
	index := sort.SearchStrings(sortedList, "Sandy")
	fmt.Println(index)

}
