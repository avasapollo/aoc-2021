package main

import (
	"github.com/avasapollo/aoc-2021/tools"
	"log"
)

func main() {
	input,err := tools.ToIntSlice("day-1/part-2/part-2.txt")
	if err != nil {
		log.Fatal(err)
	}

	var preResult []int
	for i,v := range input[:len(input) - 2] {
		sum := v + input[i+1] + input[i+2]
		preResult = append(preResult,sum)
	}

	count := 0
	for i,v := range preResult {
		if i+1 == len(preResult) {
			continue
		}
		if v < preResult[i+1] {
			count ++
		}
	}
	log.Printf("the counter is %d ",count)
}
