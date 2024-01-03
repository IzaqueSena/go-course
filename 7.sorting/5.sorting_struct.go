package main

import (
	"fmt"
	"sort"
)

type Player struct {
	score int
	name  string
}

func sortPlayers(s []*Player) {
	sort.Slice(s, func(i, j int) bool {
		if s[i].score != s[j].score {
			return s[i].score < s[j].score
		} else {
			return s[i].name > s[j].name
		}
	})
}

func main() {
	players := []*Player{
		{10, "abc"},
		{1, "dbc"},
		{4, "zbc"},
		{6, "abc"},
		{4, "abc"},
	}
	sortPlayers(players)
	for i := 0; i < len(players); i++ {
		fmt.Println(*players[i])
	}
}
