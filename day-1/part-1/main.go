package main

import (
	"github.com/avasapollo/aoc-2021/tools"
	"log"
)

func main() {
	input,err := tools.ToIntSlice("day-1/part-1/part-1.txt")
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for i,v := range input {
		if i+1 == len(input) {
			continue
		}
		if v < input[i+1] {
			count ++
		}
	}
	log.Printf("the counter is %d ",count)
}

