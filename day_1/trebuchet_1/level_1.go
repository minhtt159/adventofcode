package trebuchet_1

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

func Parse(line string) int {
	// Match regex
	two_number := `[^\d]*(\d).*(\d)[^\d]*`
	r2 := regexp.MustCompile(two_number)
	one_number := `[^\d]*(\d)[^\d]*`
	r1 := regexp.MustCompile(one_number)

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

	return number
}
