package main

import (
	"fmt"
)

type bump struct {
	location Location
	players  []*Player
}

var bumps = make(map[Location][]*Player)

// HandleBump records what happens when events occur
func HandleBump(loc Location, target *Player) {

	if _, ok := bumps[loc]; !ok {
		bumps[loc] = make([]*Player, 1)
	}

	bumps[loc] = append(bumps[loc], target)
}

// Complete means all bumps have been processed
func Complete() {
	for k, v := range bumps {
		b := bump{
			k,
			v,
		}

		ai(b)
	}

	// Reset
	bumps = make(map[Location][]*Player)
}

func ai(b bump) {

	fmt.Printf("The following players are bumping into another player\n%+q\n", b.players)
	fmt.Printf("The square affected is %v o\n", b.location)

}
