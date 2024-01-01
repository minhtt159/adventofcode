package gear_1

import (
	"bytes"
	"container/list"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

type Pair struct {
	x, y interface{}
}

func isNumber(x byte) bool {
	return unicode.IsDigit(rune(x))
}

func isGear(x byte) bool {
	if x == []byte(".")[0] {
		return false
	}
	return !isNumber(x)
}

func isValid(x int, y int, visited []string) bool {
	if x < 0 || y < 0 || x >= len(visited) || y >= len(visited[0]) {
		return false
	}
	if visited[x][y] != []byte("0")[0] {
		return false
	}
	return true
}

func BFS(i int, j int, schematic []string, visited []string) {
	if visited[i][j] != []byte("0")[0] {
		// Already visited, skip
		return
	}
	is_part := false
	// And mark as visited
	if isGear(schematic[i][j]) {
		// a gear
		visited[i] = replaceAtIndex2(visited[i], "2", j)
		is_part = true
	} else if schematic[i][j] == []byte(".")[0] {
		// a dot
		visited[i] = replaceAtIndex2(visited[i], "1", j)
		return
	}
	// Else, it is a number, and do BFS to see if it connected to any gear

	// Init the queue
	queue := list.New()
	mark_queue := list.New()

	// Put it into a queue
	queue.PushBack(Pair{i, j})
	mark_queue.PushBack(Pair{i, j})

	// Direction
	row := []int{1, 1, 1, 0, 0, -1, -1, -1}
	col := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for {
		if queue.Len() == 0 {
			break
		}
		node := queue.Front()
		x := node.Value.(Pair).x.(int)
		y := node.Value.(Pair).y.(int)
		// fmt.Println(x, y)
		queue.Remove(node)
		// traverse
		for k := 0; k < len(row); k++ {
			new_x := x + row[k]
			new_y := y + col[k]
			// If new node isValid, and not visited
			if isValid(new_x, new_y, visited) {
				// Mark visited
				if isGear(schematic[new_x][new_y]) {
					// It is a gear, so visit and mark the whole chain is_part
					visited[new_x] = replaceAtIndex2(visited[new_x], "2", new_y)
					is_part = true
				} else if isNumber(schematic[new_x][new_y]) {
					if is_part {
						visited[new_x] = replaceAtIndex2(visited[new_x], "2", new_y)
					} else {
						// It is not yet is_part, so put it in a queue and fill later
						visited[new_x] = replaceAtIndex2(visited[new_x], "1", new_y)
						mark_queue.PushBack(Pair{new_x, new_y})
					}
				} else {
					// It is a dot, so visited and skip
					visited[new_x] = replaceAtIndex2(visited[new_x], "1", new_y)
					continue
				}
				// Push new node to queue
				queue.PushBack(Pair{new_x, new_y})
			}
		}
	}
	if is_part {
		for {
			if mark_queue.Len() == 0 {
				break
			}
			node := mark_queue.Front()
			x := node.Value.(Pair).x.(int)
			y := node.Value.(Pair).y.(int)
			mark_queue.Remove(node)
			visited[x] = replaceAtIndex2(visited[x], "2", y)
		}
	}
	return
}

func replaceAtIndex2(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func Solve(schematic []string) int {
	// fmt.Println(schematic)
	//
	sum := 0
	// Create a visit table
	visited := []string{}
	for _, line := range schematic {
		empty_string := string(bytes.Repeat([]byte("0"), len(line)))
		visited = append(visited, empty_string)
		// fmt.Println(empty_string, reflect.TypeOf(empty_string))
	}
	// Loop through the schematic
	for x, line := range schematic {
		for y := 0; y < len(line); y++ {
			BFS(x, y, schematic, visited)
		}
	}
	// Debug
	if false {
		for x, line := range schematic {
			fmt.Println(line)
			fmt.Println(visited[x])
		}
	}
	// Regex
	r_gear := regexp.MustCompile(`(2*)`)
	r_number := regexp.MustCompile(`(\d*)`)
	// Get all number is_part
	for x, line := range schematic {
		gear_part := r_gear.FindAllStringIndex(visited[x], -1)
		// fmt.Println(line)
		// fmt.Println(visited[x])
		// fmt.Println(gear_part)

		for _, part := range gear_part {
			origin_schematic := line[part[0]:part[1]]
			if len(origin_schematic) == 0 {
				continue
			}
			number_part := r_number.FindAllString(origin_schematic, -1)
			fmt.Println(number_part)
			for _, number := range number_part {
				real_number, err := strconv.Atoi(number)
				if err != nil {
					continue
				}
				// fmt.Println(number)
				sum += real_number
			}
		}
	}
	return sum
}
