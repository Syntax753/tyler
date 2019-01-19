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

// Players stores all players and grid size
type Players struct {
	Locations map[Location]*Player
	LastKey   string
	X         int
	Y         int
}

// GetPlayers returns all players
func (p *Players) GetPlayers() []*Player {
	ps := make([]*Player, 0, len(p.Locations))
	for _, v := range p.Locations {
		ps = append(ps, v)
	}

	return ps
}

func newPlayers(w world) *Players {
	cnt := len(w.Ents)

	var p Players
	p.Locations = make(map[Location]*Player, cnt)
	p.X = w.X
	p.Y = w.Y

	for _, ent := range w.Ents {
		origin := ent
		p.Locations[NewLocation(ent.X, ent.Y*w.X)] = &origin
	}

	return &p
}

const (
	empty = "."
)

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
	p := newPlayers(w)
	o = NewMozart(w.X, w.Y)

	game(p)

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
	var ch string
	for y := 0; y < ps.Y; y++ {
		for x := 0; x < ps.X; x++ {
			ent, ok := ps.Locations[NewLocation(x, y)]
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

type world struct {
	X    int      `json:"x"`
	Y    int      `json:"y"`
	Ents []Player `json:"ents"`
}

// Player represents an object in the world
type Player struct {
	Ref    string `json:"ref"`
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

// SetLocation finalises the position of the player
func (p *Player) SetLocation(loc Location) {
	p.X, p.Y = loc.GetCoords()
}

// GetType returns the type (should be movement pref)
func (p *Player) GetType() string {
	return p.Type
}

// Move moves a single player
func (p *Player) Move(ps *Players) Location {
	worldX = ps.X
	worldY = ps.Y

	fmt.Printf("&&&&&&&&&&&&&&%v", p)

	var d = p.GetLocation()

	switch p.GetType() {
	case "player":
		switch ps.LastKey {
		case "a":
			d = MovePlayer(p, -1, 0)
		case "d":
			d = MovePlayer(p, 1, 0)
		case "w":
			d = MovePlayer(p, 0, -1)
		case "s":
			d = MovePlayer(p, 0, 1)
		}
	}

	return d

}

// GetLocation gives coords of the player
func (p *Player) GetLocation() Location {
	return Location{p.X, p.Y}
}
