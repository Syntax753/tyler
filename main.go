package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	o Mozart
)

const (
	empty = "."
	wall  = "#"
)

// Players stores all players and grid size
type Players struct {
	Players []*Player
	LastKey string
	X       int
	Y       int
}

// GetLocations returns all players
func (ps *Players) GetLocations() map[Location]*Player {
	locs := make(map[Location]*Player)
	for _, p := range ps.Players {
		locs[p.Location] = p
	}

	// fmt.Printf("Locations are %v\n", locs)
	return locs
}

func newPlayers(w world) *Players {
	var ps Players
	ps.Players = make([]*Player, len(w.Ents))
	ps.X = w.X
	ps.Y = w.Y

	if w.Iswalled {
		for i := 0; i < w.X; i++ {
			ent := Player{
				Location: NewLocation(i, 0),
				Type:     "wall",
				Symbol:   wall,
			}

			ps.Players = append(ps.Players, &ent)

			ent2 := Player{
				Location: NewLocation(i, w.Y-1),
				Type:     "wall",
				Symbol:   wall,
			}

			ps.Players = append(ps.Players, &ent2)

		}
	}

	for i, ent := range w.Ents {
		origin := ent
		origin.SetLocation(NewLocation(ent.X, ent.Y))
		ps.Players[i] = &origin
	}

	return &ps
}

var (
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	jsonFile, err := os.Open("world.json")
	if err != nil {
		panic(err)
	}
	fmt.Println("Reading world")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var w world
	json.Unmarshal(byteValue, &w)
	ps := newPlayers(w)
	o = NewMozart(w.X, w.Y)

	game(ps)
}

func game(ps *Players) {
	for {
		draw(ps)

		r, err := reader.ReadString('\n')
		r = strings.Replace(r, "\n", "", -1)
		if err != nil {
			fmt.Println(err)
			break
		}

		if r == "p" {
			break
		}

		ps.LastKey = r
		o.MoveAll(ps)
	}
}

func draw(ps *Players) {
	locs := ps.GetLocations()

	var ch string
	for y := 0; y < ps.Y; y++ {
		for x := 0; x < ps.X; x++ {
			loc := NewLocation(x, y)
			ent, ok := locs[loc]
			if ok {
				ch = ent.Symbol
			} else {
				ch = empty
			}

			fmt.Printf(ch)
		}
		fmt.Println()
	}

}

// SetLocation finalises the position of the player
func (p *Player) SetLocation(loc Location) {
	p.Location = loc
}

// GetType returns the type (should be movement pref)
func (p *Player) GetType() string {
	return p.Type
}

// GetLocation gives coords of the player
func (p *Player) GetLocation() Location {
	return p.Location
}

type world struct {
	X        int      `json:"x"`
	Y        int      `json:"y"`
	Ents     []Player `json:"ents"`
	Iswalled bool     `json:"iswalled"`
}

// Player represents an object in the world
type Player struct {
	Ref      string `json:"ref"`
	Symbol   string `json:"symbol"`
	Type     string `json:"type"`
	Location Location
	X        int
	Y        int
}
