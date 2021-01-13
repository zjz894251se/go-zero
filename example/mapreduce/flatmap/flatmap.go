package main

import (
	"fmt"

	"github.com/zjz894251se/go-zero/core/mr"
)

var (
	persons = []string{"john", "mary", "alice", "bob"}
	friends = map[string][]string{
		"john":  {"harry", "hermione", "ron"},
		"mary":  {"sam", "frodo"},
		"alice": {},
		"bob":   {"jamie", "tyrion", "cersei"},
	}
)

func main() {
	var allFriends []string
	for v := range mr.Map(func(source chan<- interface{}) {
		for _, each := range persons {
			source <- each
		}
	}, func(item interface{}, writer mr.Writer) {
		writer.Write(friends[item.(string)])
	}, mr.WithWorkers(100)) {
		allFriends = append(allFriends, v.([]string)...)
	}
	fmt.Println(allFriends)
}
