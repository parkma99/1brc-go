package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type station struct {
	name string
	max  float64
	sum  float64
	min  float64
	num  int64
}

func main() {
	maps := make(map[string]station)
	file, err := os.Open("./measurements.txt")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		index := strings.Index(text, ";")
		if index == -1 {
			fmt.Println(text)
			os.Exit(1)
		}
		city := text[:index]
		temp, _ := strconv.ParseFloat(text[index+1:], 32)

		if _, ok := maps[city]; ok {
			cur_max := maps[city].max
			if temp > cur_max {
				cur_max = temp
			}
			cur_min := maps[city].min
			if cur_min > temp {
				temp = cur_min
			}
			maps[city] = station{name: city, max: cur_max, sum: maps[city].sum + 1, min: cur_min, num: maps[city].num + 1}
		} else {
			maps[city] = station{name: city, max: temp, sum: temp, min: temp, num: 1}
		}
	}
	var result []string
	for _, station := range maps {
		input := fmt.Sprintf("%s=%.1f/%.1f/%.1f", station.name, station.max, station.min, station.sum/float64(station.num))
		result = append(result, input)
	}
	sort.Strings(result)
	output := strings.Join(result, ", ")
	fmt.Println(output)
}
