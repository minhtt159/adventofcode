package scratch1

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func AizuArray(A string) []int {
	strs := strings.Split(A, " ")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}
	return ary
}

func Solve(schematic []string) int {
	final_score := 0

	for _, line := range schematic {
		this_score := 0
		reg := regexp.MustCompile(`Card.*\d: (?P<win_number>.*) \| (?P<card_number>.*)`)
		groups := reg.FindStringSubmatch(line)
		win_number := groups[reg.SubexpIndex("win_number")]
		card_number := groups[reg.SubexpIndex("card_number")]
		array_card_number := AizuArray(card_number)
		fmt.Println(win_number, "\t", array_card_number)
		for _, number := range regexp.MustCompile(`(\d*)`).FindAllStringSubmatch(win_number, -1) {
			if len(number[0]) == 0 {
				continue
			}
			fmt.Println(number)
			this_number, err := strconv.Atoi(number[1])
			if err != nil {
				panic(err)
			}
			// fmt.Println(this_number)
			if slices.Contains(array_card_number, this_number) {
				if this_score == 0 {
					this_score = 1
				} else {
					this_score = this_score * 2
				}
			} else {
				// fmt.Println("number not found")
			}
		}
		final_score += this_score
		// break
	}
	return final_score
}
