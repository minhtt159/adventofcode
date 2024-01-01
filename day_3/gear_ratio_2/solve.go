package gear_2

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
)

func isNumber(x byte) bool {
	return unicode.IsDigit(rune(x))
}

func isGear(x byte) bool {
	if x != []byte("*")[0] {
		return false
	}
	return true
}

func isValid(i int, j int, schematic []string, visited []string) bool {
	// Check if this index is in the board and also isNumber
	// Out of map
	if i < 0 || j < 0 || i >= len(schematic) || j >= len(schematic[0]) {
		return false
	}
	// Already visited
	if visited[i][j] == []byte("1")[0] {
		return false
	}
	if !isNumber(schematic[i][j]) {
		return false
	}
	return true
}

// Direction
var row = []int{1, 1, 1, 0, 0, -1, -1, -1}
var col = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func replaceAtIndex2(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func BFS(i int, j int, schematic []string, visited []string) int {
	/*
	* So schematic[i][j] is a gear, we BFS from here
	* to see if we have 2 numbers
	*   if yes: return multiplication
	*   if no: return 0
	 */
	if schematic[i][j] != []byte("*")[0] {
		// Something is wrong here
		return 0
	}
	// fmt.Println(i, j, "is a gear")
	// Mark gear as visited
	visited[i] = replaceAtIndex2(visited[i], "1", j)

	numbers_adjacent := []int{}
	// BFS
	for k := 0; k < len(row); k++ {
		new_x := i + row[k]
		new_y := j + col[k]
		// If new position is not visited and isNumber
		if !isValid(new_x, new_y, schematic, visited) {
			continue
		}
		// Move to left most number
		for {
			if !isValid(new_x, new_y, schematic, visited) {
				new_y = new_y + 1
				break
			}
			new_y = new_y - 1
		}
		// fmt.Println(new_x, new_y)
		visited[new_x] = replaceAtIndex2(visited[new_x], "1", new_y)
		// Trying to parse the adjacent numbers
		number_buffer := string(schematic[new_x][new_y])
		// Move to the right until the end of number_buffer
		for {
			new_y = new_y + 1
			if !isValid(new_x, new_y, schematic, visited) {
				break
			}
			visited[new_x] = replaceAtIndex2(visited[new_x], "1", new_y)
			number_buffer = number_buffer + string(schematic[new_x][new_y])
		}
		// fmt.Println(number_buffer)
		if len(number_buffer) > 0 {
			num, err := strconv.Atoi(number_buffer)
			if err != nil {
				panic(err)
			}
			numbers_adjacent = append(numbers_adjacent, num)
		}
	}
	fmt.Println(numbers_adjacent)
	if len(numbers_adjacent) == 2 {
		return numbers_adjacent[0] * numbers_adjacent[1]
	}
	return 0
}

func Solve(schematic []string) int {
	/*
	* Bro, since the gear is only *
	* How about scan for *-gear and BFS from there?
	 */

	// fmt.Println(schematic)
	//
	sum := 0

	visited := []string{}
	for _, line := range schematic {
		empty_string := string(bytes.Repeat([]byte("0"), len(line)))
		visited = append(visited, empty_string)
		// fmt.Println(empty_string, reflect.TypeOf(empty_string))
	}

	// Loop through the schematic
	for x, line := range schematic {
		for y := 0; y < len(line); y++ {
			if isGear(schematic[x][y]) {
				sum = sum + BFS(x, y, schematic, visited)
			}
		}
	}
	return sum
}
