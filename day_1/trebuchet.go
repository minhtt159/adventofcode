package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	// Match regex
	two_number := `[^\d]*(\d).*(\d)[^\d]*`
	r2 := regexp.MustCompile(two_number)
	one_number := "^[^\\d]*(\\d)[^\\d]*$"
	r1 := regexp.MustCompile(one_number)

	// Init
	sum := 0
	scanner := bufio.NewScanner(file)
	// Loop though lines
	for scanner.Scan() {
		line := scanner.Text()
		// Match
		line_number := ""
		if r2.MatchString(line) == true {
			res := r2.FindStringSubmatch(line)
			line_number = fmt.Sprintf("%s%s", res[1], res[2])
		} else if r1.MatchString(line) == true {
			res := r1.FindStringSubmatch(line)
			line_number = fmt.Sprintf("%s%s", res[1], res[1])
		} else {
			fmt.Println("Number not found")
			line_number = ""
		}
		// Convert to number
		number, err := strconv.Atoi(line_number)
		check(err)
		sum = sum + number
	}
	fmt.Println(sum)
}
