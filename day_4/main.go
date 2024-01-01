package main

import (
	"bufio"
	"fmt"
	"os"
	// custom
	solver "scratch/scratch_2"
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
	schematic := []string{}
	scanner := bufio.NewScanner(file)
	// Loop though lines
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, line)
	}
	// End
	fmt.Println(solver.Solve(schematic))
}
