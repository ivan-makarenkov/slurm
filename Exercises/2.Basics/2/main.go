package main

import "fmt"

func main() {
	fmt.Println(isSorted([]string{"a", "ab", "b", "bc", "c", "cd"}))
}

func isSorted(ww []string) bool {
	if ww == nil {
		return false
	}
	var lw string
	for _, w := range ww {
		if lw > w {
			return false
		}
		lw = w
	}
	return true
}
