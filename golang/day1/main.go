package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dial := 50
	data := load_file("input")
	total := 0

	for _, val := range data {
		direction := string(val[0])
		rotations, _ := strconv.Atoi(val[1:])
		fmt.Println("Now processing", direction, rotations)
		timesZero, newDial := recursiveFunc(dial, direction, rotations, 0, 99)
		total += timesZero
		dial = newDial
		fmt.Println("Current Dial === ", dial)
		fmt.Println("Current total ===", total)
		// if idx == 1 {
		// 	break
		// }
	}

	fmt.Println("Password is ", total)
}

func load_file(name string) []string {
	dat, err := os.ReadFile(name)
	check(err)

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	return lines
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func recursiveFunc(currentDial int, direction string, rotations int, min int, max int) (int, int) {
	fmt.Println("RecursiveFunc: ", currentDial, direction, rotations, min, max)
	counter := 0
	newDial := 0

	if direction == "L" {
		newDial = currentDial - rotations
	} else {
		newDial = currentDial + rotations
	}

	if newDial > max {
		newCounter, dial := recursiveFunc(newDial, "L", 100, min, max)
		counter += newCounter + 1
		newDial = dial
	} else if newDial < min {
		newCounter, dial := recursiveFunc(newDial, "R", 100, min, max)
		counter += newCounter + 1
		newDial = dial
	}
	// else if newDial == 0 { required for pt1
	//
	// 	counter =+ 1
	// }

	return counter, newDial
}
