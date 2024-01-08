package part_1

import (
	"fmt"
	"regexp"
	// "regexp"
	// "slices"
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

type Pair struct {
	upper  int
	lower  int
	weight int
	dest   int
}

func reset(upper int, size int) []Pair {
	mapper := []Pair{}
	for i := 0; i < size; i++ {
		mapp := Pair{upper, 0, 0, 0}
		mapper = append(mapper, mapp)
	}
	return mapper
}

func Solve(schematic []string) int {
	seeds_line := schematic[0]
	// fmt.Println(seeds_line[7:])
	seeds := AizuArray(seeds_line[7:])
	fmt.Println(seeds)
	// Init the game
	max_seed := 0
	for _, seed := range seeds {
		if seed > max_seed {
			max_seed = seed
		}
	}
	mapper := reset(max_seed, len(seeds))
	// Start from line 3
	index := 3
	for {
		if index >= len(schematic) {
			// end of file
			for i := 0; i < len(seeds); i++ {
				// fmt.Println(mapper[i])
				seeds[i] = seeds[i] - mapper[i].lower + mapper[i].dest
			}
			break
		}
		this_line := schematic[index]
		if this_line == "" {
			// end of map, translate mapper
			for i := 0; i < len(seeds); i++ {
				// fmt.Println(mapper[i])
				seeds[i] = seeds[i] - mapper[i].lower + mapper[i].dest
			}
			fmt.Println("new map", mapper)
			mapper = reset(max_seed, len(seeds))
			index += 2
			continue
		}
		// Parse int
		reg := regexp.MustCompile(`(?P<dest>\d*) (?P<src>\d*) (?P<len>\d*)`)
		groups := reg.FindStringSubmatch(this_line)
		dest, err := strconv.Atoi(groups[reg.SubexpIndex("dest")])
		if err != nil {
			panic(err)
		}
		src, err := strconv.Atoi(groups[reg.SubexpIndex("src")])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(groups[reg.SubexpIndex("len")])
		if err != nil {
			panic(err)
		}
		// Prepare the range
		lower_src := src
		upper_src := src + length
		fmt.Println(lower_src, upper_src)
		// Update mapper
		for i := 0; i < len(seeds); i++ {
			if lower_src <= seeds[i] && upper_src > seeds[i] {
				// range found
				mapper[i] = Pair{upper_src, lower_src, length, dest}
			}
		}
		index += 1
	}
	for i := 0; i < len(seeds); i++ {
		fmt.Println(seeds[i])
		if max_seed > seeds[i] {
			max_seed = seeds[i]
		}
	}
	return max_seed
}
