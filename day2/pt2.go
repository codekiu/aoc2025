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
		fmt.Println("Ranges ", val)
		ranges := strings.Split(val, "-")
		min, _ := strconv.Atoi(ranges[0])
		max, _ := strconv.Atoi(strings.TrimSpace(ranges[1]))
		for i := min; i <= max; i++ {
			lenOfNumber := len(strconv.Itoa(i))
			var numberOfItemsToSelect []int

			for i := lenOfNumber; i > 0; i-- {
				if lenOfNumber%i == 0 {
					numberOfItemsToSelect = append(numberOfItemsToSelect, i)
				}
			}

			found := totalFound(strconv.Itoa(i), numberOfItemsToSelect)

			total += found
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

func totalFound(s string, sizes []int) int {
	numbers := map[string]struct{}{}
	seen := false
	foundCases := 0
	fmt.Println("Processing: ", s)

	for _, size := range sizes {
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
				} else {
					seen = false
					break
				}
			}
		}

		if seen {
			found, _ := strconv.Atoi(s)
			fmt.Println("Found: ", s)
			foundCases += found
			seen = false
			break
		}
	}
	return foundCases
}

func contains(s map[string]struct{}, char string) bool {
	if _, ok := s[char]; ok {
		return true
	}
	return false
}
