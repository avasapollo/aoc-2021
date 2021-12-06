package main

import (
	"github.com/avasapollo/aoc-2021/tools"
	"log"
	"strconv"
	"strings"
)

type Action struct {
	Direction string
	Value int
}

func main() {
	s,err := tools.ToStringSlice("day-2/part-1/part-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	actions,err := toActions(s)
	if err != nil {
		log.Fatal(err)
	}
	h := 0
	d := 0
	for _,action:= range actions {
		switch action.Direction {
		case "forward":
			h += action.Value
		case "up":
			d -= action.Value
		case "down":
			d += action.Value
		}
	}

	log.Printf("h: %d d: %d t: %d",h,d,h*d)
}


func toActions(input []string) ([]*Action,error) {
	var actions []*Action
	for _,v := range input {
		l := strings.Split(v," ")
		i,err :=  strconv.Atoi(l[1])
		if err != nil {
			return nil, err
		}
		actions = append(actions,&Action{
			Direction: l[0],
			Value:   i ,
		})
	}
	return actions,nil
}