package conundrum_1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Max number
var red = 12
var green = 13
var blue = 14

func Max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func Parse(line string) int {
	// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	r_game := regexp.MustCompile(`Game (\d*):`)
	// Extract Game number and the rest, just messing
	number_end := r_game.FindStringSubmatch(line)
	game_number, _ := strconv.Atoi(number_end[1])
	// fmt.Println(game_number)
	// Extract the rest
	game_start := r_game.FindStringIndex(line)
	games := line[game_start[1]:]
	// Split string
	hands := strings.Split(games, ";")
	// Regex match each hands
	r_hand := regexp.MustCompile(`(\d*) (green|blue|red)`)
	game_red := 0
	game_blue := 0
	game_green := 0
	for _, hand := range hands {
		all_matches := r_hand.FindAllStringSubmatch(hand, -1)
		for _, match := range all_matches {
			number, _ := strconv.Atoi(match[1])
			color := match[2]
			switch color {
			case "red":
				game_red = Max(game_red, number)
			case "blue":
				game_blue = Max(game_blue, number)
			case "green":
				game_green = Max(game_green, number)
			}
		}
	}
	if game_red > red {
		return 0
	}
	if game_blue > blue {
		return 0
	}
	if game_green > green {
		return 0
	}
	fmt.Println(line)
	fmt.Println("r", game_red, "b", game_blue, "g", game_green)
	return game_number
}
