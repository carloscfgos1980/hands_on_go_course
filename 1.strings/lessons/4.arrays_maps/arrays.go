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
	// Find the index of "Draper" in the slice using a map
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
	// Copy the original slice to the new slice before sorting
	copy(sortedList, str)
	// Sort the new slice in place
	sort.Strings(sortedList)
	fmt.Println(sortedList)
	// Find the index of "Sandy" in the sorted slice using binary search
	index := sort.SearchStrings(sortedList, "Sandy")
	fmt.Println(index)

}

func RevertArray() {
	str := []string{"Sandy", "Provo", "St. George", "Salt Lake City", "Draper", "South Jordan", "Murray"}
	fmt.Println(str)
	// Revert the slice in place
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	fmt.Println(str)
	numbers := []int{1, 5, 3, 6, 2, 10, 8}
	// Sort the slice in ascending order
	tobeSorted := sort.IntSlice(numbers)
	sort.Sort(tobeSorted)
	fmt.Println(tobeSorted)
	// Sort the slice in descending order
	sort.Sort(sort.Reverse(tobeSorted))
	fmt.Println(tobeSorted)
}

func IterateArray() {
	str := []string{"Sandy", "Provo", "St. George", "Salt Lake City", "Draper", "South Jordan", "Murray"}
	fmt.Println(str)
	// Iterate through the slice using a for loop
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
	// Iterate through the slice using a range loop
	for i, v := range str {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
	// Iterate through the slice in reverse order
	str = []string{"Sandy", "Provo", "St. George", "Salt Lake City", "Draper", "South Jordan", "Murray"}
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Println(str[i])
	}
}

func MapToArray() {
	type NameAge struct {
		Name string
		Age  int
	}
	var nameAgeSlice []NameAge
	nameAges := map[string]int{
		"Michael": 30,
		"John":    25,
		"Jessica": 26,
		"Ali":     18,
	}
	for key, value := range nameAges {
		nameAgeSlice = append(nameAgeSlice, NameAge{key, value})
	}
	fmt.Println(nameAgeSlice)
}

func MergeArrays() {
	// Merge two slices into one
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	mergedSlice := append(slice1, slice2...)
	fmt.Println(mergedSlice)
}

func MergeMaps() {
	map1 := map[string]int{
		"Michael": 10,
		"Jessica": 20,
		"Tarik":   33,
		"Jon":     22,
	}
	fmt.Println(map1)
	map2 := map[string]int{
		"Lord":  11,
		"Of":    22,
		"The":   36,
		"Rings": 23,
	}
	for key, value := range map2 {
		map1[key] = value
	}
	fmt.Println(map1)
}

func CheckValueInMap() {
	map1 := map[string]int{
		"Michael": 10,
		"Jessica": 20,
		"Tarik":   33,
		"Jon":     22,
	}
	fmt.Println(map1)
	if value, ok := map1["Michael"]; ok {
		fmt.Println("Michael is in the map with value:", value)
	} else {
		fmt.Println("Michael is not in the map")
	}
	if value, ok := map1["Ali"]; ok {
		fmt.Println("Ali is in the map with value:", value)
	} else {
		fmt.Println("Ali is not in the map")
	}
}
