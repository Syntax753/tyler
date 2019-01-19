package main

// import "fmt"

var (
	worldX, worldY int
)

// Location for movement
type Location struct {
	x, y int
}

// GetCoords for coordinates of a location
func (l *Location) GetCoords() (x, y int) {
	return l.x, l.y
}

// NewLocation constructor
func NewLocation(x, y int) Location {
	loc := Location{x, y}
	return loc
}

// MovePlayer the player
func MovePlayer(p Mover, hor, vert int) Location {
	loc := Location{p.GetLocation().x + hor, p.GetLocation().y + vert}
	return loc
}

// Mover handles all movement of players
type Mover interface {
	Move(*Players) Location
	GetType() string
	GetLocation() Location
}

// Move moves a single player
func (p *Player) Move(ps *Players) Location {
	worldX = ps.X
	worldY = ps.Y

	// fmt.Printf("&&&&&&&&&&&&&&%v", p)

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
