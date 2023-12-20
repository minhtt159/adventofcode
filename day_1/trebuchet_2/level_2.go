package trebuchet_2

import (
	"fmt"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Regex Magicr
var r = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)

func Match(substr string) ([]string, error) {
	// Unpack the number out of this smaller string
	interesting_string := r.FindStringSubmatch(substr)
	// fmt.Println(substr, interesting_string)
	if interesting_string == nil {
		return []string{}, fmt.Errorf("Nothing here")
	}
	result := []string{}
	for i := 1; i < len(interesting_string); i++ {
		word := interesting_string[i]
		switch word {
		case "one":
			word = "1"
		case "two":
			word = "2"
		case "three":
			word = "3"
		case "four":
			word = "4"
		case "five":
			word = "5"
		case "six":
			word = "6"
		case "seven":
			word = "7"
		case "eight":
			word = "8"
		case "nine":
			word = "9"
		case "":
			continue
		}
		result = append(result, word)
		// fmt.Println(word)
	}
	return result, nil
}

func Parse(old_line string) int {
	// Init empty array
	var arr []string
	line := fmt.Sprintf("%5s", old_line)
	// Loop through all character in string
	for index := 0; index <= len(line); index++ {
		last_char := index + 5
		if last_char > len(line) {
			last_char = len(line)
		}
		sub_str := line[index:last_char]
		numbers, err := Match(sub_str)
		if err == nil {
			// Number is good, proceed
			arr = append(arr, numbers...)
			// fmt.Println(number)
		}
	}
	// fmt.Println(arr, line)

	// Constructing the final number
	res := fmt.Sprintf("%s%s", arr[0], arr[len(arr)-1])
	num, _ := strconv.Atoi(res)
	return num
}
