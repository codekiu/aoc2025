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
		fmt.Println("Ranges: ", ranges)
		min, _ := strconv.Atoi(ranges[0])
		max, _ := strconv.Atoi(ranges[1])
		for i := min; i <= max; i++ {
			fmt.Println("Processing: ", i)
			lenOfNumber := len(strconv.Itoa(i))
			if lenOfNumber%2 != 0 {
				continue
			}

			numberOfItemsToSelect := lenOfNumber / 2
			numberSplitted := chunkString(strconv.Itoa(i), numberOfItemsToSelect)
			seen := false
			numbers := map[string]struct{}{}

			fmt.Println("NumbersSplitted:", numberSplitted, " splitting by every ", numberOfItemsToSelect)

			for _, x := range numberSplitted {
				found := contains(numbers, x)

				if found {
					seen = true
				}

				numbers[x] = struct{}{}
			}

			if seen {
				total += i
			}

		}
		break
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

func chunkString(s string, size int) []string {
	fmt.Println("Chunking - ", s)
	var chunks []string
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

func contains(s map[string]struct{}, char string) bool {
	if _, ok := s[char]; ok {
		return true
	}
	return false
}
