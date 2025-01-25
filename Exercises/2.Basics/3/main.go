package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(stringStat("ФиНИшшш"))
}

func stringStat(word string) string {
	var result []string
	l := make(map[rune]int)
	sl := make(map[rune]int)

	for _, c := range strings.ToLower(word) {
		l[c]++
	}
	for _, c := range strings.ToLower(word) {
		if c == ' ' {
			continue
		}
		sl[c]++
		if sl[c] == 1 {
			result = append(result, string(c)+" - "+strconv.Itoa(l[c]))
		}
	}

	return strings.Join(result, "\n")
}
