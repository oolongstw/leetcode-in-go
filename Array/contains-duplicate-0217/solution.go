package main

import "fmt"

func containsDuplicate(nums []int) bool {
	record := make(map[int]bool, len(nums))
	for _, n := range nums {
		if _, found := record[n]; found {
			return true
		}
		record[n] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
