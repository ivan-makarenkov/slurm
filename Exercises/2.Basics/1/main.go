package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(doubleDetector(nums))
}

func doubleDetector(nums []int) bool {
	var values = make(map[int]int)
	for _, num := range nums {
		values[num]++
		if values[num] > 1 {
			return true
		}
	}

	return false
}
