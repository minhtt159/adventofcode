package part_1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
  "container/list"
)

// AizuArray convert a string to an array of int
func AizuArray(A string) []int {
	strs := strings.Split(A, " ")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}
	return ary
}

// SeedRange is a range of seed
type SeedRange struct {
	lower int // start of the seed range
	upper int // end of the seed range
}

type MapperSet struct {
	dest   int // destination of the mapper
	src    int // source of the mapper
	length int // length of the mapper
}

// Translate seed range to next layer
func translate(seed SeedRange, mapper MapperSet) SeedRange {
  if seed.lower < mapper.src || seed.upper > mapper.src + mapper.length {
    fmt.Println("seed is not within the mapper", seed, mapper)
    panic(1)
  }
  new_lower := mapper.dest + (seed.lower - mapper.src)
  new_upper := mapper.dest + (seed.upper - mapper.src)
  result := SeedRange{new_lower, new_upper}
  fmt.Println("Translate: ", seed, "to", result)
  return result 
}

func Solve(schematic []string) int {
	seeds_line := schematic[0]
	// fmt.Println(seeds_line[7:])
	game_line := AizuArray(seeds_line[7:])
  // Golang doesn't have a set, so init the game using container/list
  seeds := list.New()
	for i := 0; i < len(game_line); i += 2 {
		// seeds = append(seeds, SeedRange{game_line[i], game_line[i] + game_line[i+1]})
    fmt.Println(game_line[i], game_line[i] + game_line[i+1])
    seeds.PushBack(SeedRange{game_line[i], game_line[i] + game_line[i+1]})
	}
	// Start from line 3
	index := 3
  mapper := []MapperSet{}
	for {
    for {
      if index >= len(schematic) {
        // end of file
        break
      }
      this_line := schematic[index]
      if this_line == "" {
        // end of map, translate mapper
        break
      }
      // Insert the mapper
      reg := regexp.MustCompile(`(?P<dest>\d*)( )*(?P<src>\d*)( )*(?P<len>\d*)`)
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
      mapper = append(mapper, MapperSet{dest, src, length})
      index += 1
    }
    // At this point, we have a full mapper for translation
    fmt.Println("This page mapper: ", mapper)
    next_layer := list.New()
    for {
      if seeds.Len() == 0 {
        // So ideally we should have a next layer
        break
      }
      seed_element := seeds.Front()
      this_seed := seed_element.Value.(SeedRange)
      seeds.Remove(seed_element)
      translated_seed := SeedRange{0, 0}
      // Loop through the mapper to translate this seed
      for i := 0; i < len(mapper); i++ {
        lower_src := mapper[i].src
        upper_src := mapper[i].src + mapper[i].length
        // fmt.Println("Seed: ", this_seed, "Mapper: ", mapper[i])
        if (this_seed.lower < lower_src) && (this_seed.upper >= lower_src) {
          // Which mean that this seed range start before this mapper and overlap with it
          // So we need to split this seed range into two parts
          // First part is the part before the mapper
          buffer_range := SeedRange{this_seed.lower, lower_src - 1}
          fmt.Println("Split before: ", buffer_range)
          seeds.PushBack(buffer_range)
          // Second part is the part after the lower src
          this_seed.lower = lower_src
        }
        if (this_seed.lower <= upper_src) && (this_seed.upper > upper_src) {
          // Which mean that this seed range overlap with this mapper and ends after it
          // So we need to split this seed range into two parts
          // Second part is the part after the mapper
          buffer_range := SeedRange{upper_src + 1, this_seed.upper}
          fmt.Println("Split after: ", buffer_range)
          seeds.PushBack(buffer_range)
          // First part is the part before the upper src
          this_seed.upper = upper_src
        }
        if this_seed.lower >= lower_src && this_seed.upper <= upper_src {
          // Which mean this seed range is within this mapper
          translated_seed = translate(this_seed, mapper[i])
          // fmt.Println("Translated seed: ", this_seed, translated_seed)
          break
        } else {
          // Which mean this seed range is not within this mapper
          // So we just put it back to the seeds
          // fmt.Println("Seed not within mapper: ", this_seed, mapper[i])
        }
      }
      if translated_seed.lower == 0 && translated_seed.upper == 0 {
        // Which mean this seed range has not been translated
        next_layer.PushBack(this_seed)
      } else { 
        next_layer.PushBack(translated_seed)
      }
    }
    // At this point, we have translated all seeds range, take the next layer and move on  
    mapper = []MapperSet{}
    seeds = next_layer
    index += 2
    if index >= len(schematic) {
      break
    }
    // break
	}
  min_seed := 1000000000000000000
  fmt.Println("Seeds final")
  for e:= seeds.Front(); e != nil; e = e.Next() {
    this_seed := e.Value.(SeedRange)
    fmt.Println(this_seed)
    if this_seed.lower < min_seed {
      min_seed = this_seed.lower
    }
  }
	return min_seed 
}
