package main

import (
	"fmt"
)

// Mozart plans moves
type Mozart struct {
	x, y int
}

// NewMozart handles movement
func NewMozart(x, y int) Mozart {
	o := Mozart{x, y}
	return o
}

func (o Mozart) Move(*Player) {

}

func validateIntent(intents map[Location][]*Player) {

}

// MoveAll calculates the anticipated move by a player
func (o Mozart) MoveAll(ps *Players) {
	var intents = make(map[Location][]*Player, len(ps.Players))

	for _, p := range ps.Players {
		loc := p.Move(ps)
		fmt.Printf("%v moves to %v", *p, loc)
		if _, ok := intents[loc]; !ok {
			fresh := make([]*Player, 0, len(ps.Players))
			intents[loc] = fresh
		}
		intents[loc] = append(intents[loc], p)
	}

	for k, v := range intents {
		if len(v) > 1 {
			for _, p := range v {
				HandleBump(k, p)
			}
		}

		Complete()

		if len(v) == 1 {
			v[0].Location = k
		}
	}

	fmt.Printf("Intents are %v\n", intents)

	// TODO if len > 1 clash

}
