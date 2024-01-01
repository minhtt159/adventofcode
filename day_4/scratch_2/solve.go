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
	scratches := []int{}
	for i := 0; i < len(schematic); i++ {
		scratches = append(scratches, 0)
	}

	for index, line := range schematic {
		isWin := false
		matches := 0

		reg := regexp.MustCompile(`Card.*\d: (?P<win_number>.*) \| (?P<card_number>.*)`)
		groups := reg.FindStringSubmatch(line)
		win_number := groups[reg.SubexpIndex("win_number")]
		card_number := groups[reg.SubexpIndex("card_number")]
		array_card_number := AizuArray(card_number)
		// fmt.Println(win_number, "\t", array_card_number)
		for _, number := range regexp.MustCompile(`(\d*)`).FindAllStringSubmatch(win_number, -1) {
			if len(number[0]) == 0 {
				continue
			}
			// fmt.Println(number)
			this_number, err := strconv.Atoi(number[1])
			if err != nil {
				panic(err)
			}
			// fmt.Println(this_number)
			if slices.Contains(array_card_number, this_number) {
				matches += 1
				isWin = true
			}
		}
		if !isWin {
			scratches[index] += 1
			continue
		}
		scratches[index] += 1
		multiplier := scratches[index]
		// fmt.Printf("win: %d, multiplier: %d ", index, multiplier)
		for i := 1; i <= matches; i++ {
			if index+i > len(schematic) {
				break
			}
			scratches[index+i] += multiplier
		}
		// fmt.Println(scratches)
		// break
	}
	fmt.Println(scratches)
	final_score := 0
	for _, cards := range scratches {
		final_score += cards
	}
	return final_score
}
