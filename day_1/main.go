package main

import (
	"bufio"
	"fmt"
	"os"

	// custom
	// solver "trebuchet/trebuchet_1"
	solver "trebuchet/trebuchet_2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Open input file
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	// Init
	sum := 0
	scanner := bufio.NewScanner(file)
	// Loop though lines
	for scanner.Scan() {
		line := scanner.Text()
		sum = sum + solver.Parse(line)
	}
	fmt.Println(sum)
}
