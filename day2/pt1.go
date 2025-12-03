package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	values := load_file("input")
	total := 0
	fmt.Println("Go!")
	for _, val := range values {
		ranges := strings.Split(val, "-")
		min, _ := strconv.Atoi(ranges[0])
		max, _ := strconv.Atoi(ranges[1])
		for i := min; i <= max; i++ {
			lenOfNumber := len(strconv.Itoa(i))
			if lenOfNumber%2 != 0 {
				continue
			}

			numberOfItemsToSelect := lenOfNumber / 2
			found := wasFound(strconv.Itoa(i), numberOfItemsToSelect)

			if found {
				total += i
			}

		}
	}

	fmt.Println("total = ", total)

}

func load_file(name string) []string {
	file, err := os.ReadFile(name)
	check(err)

	dat := strings.Split(string(file), ",")
	var values []string

	for _, i := range dat {
		values = append(values, i)
	}

	return values
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func wasFound(s string, size int) bool {
	numbers := map[string]struct{}{}
	seen := false

	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}

		if i == 0 {
			numbers[s[i:end]] = struct{}{}
		} else {
			if contains(numbers, s[i:end]) {
				seen = true
			}
		}
	}
	return seen
}

func contains(s map[string]struct{}, char string) bool {
	if _, ok := s[char]; ok {
		return true
	}
	return false
}
