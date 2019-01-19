package main

// Mozart plans moves
type Mozart struct {
	x, y int
}

// NewMozart handles movement
func NewMozart(x, y int) Mozart {
	o := Mozart{x, y}
	return o
}

// MoveAll calculates the anticipated move by a player
func (o Mozart) MoveAll(ps *Players) {
	players := append(make([]*Player, 0, len(ps.Locations)), ps.GetPlayers()...)
	var intents = make(map[Location][]*Player, len(players))

	for _, p := range players {
		loc := p.Move(ps)
		if _, ok := intents[loc]; !ok {
			fresh := make([]*Player, 0, len(players))
			intents[loc] = fresh
		}
		intents[loc] = append(intents[loc], p)
	}

	for k, v := range intents {
		if len(v) == 1 {
			ps.Locations[k] = v
		}
	}

	// TODO if len > 1 clash

}
