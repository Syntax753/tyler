package main

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
	loc := Location{p.GetLocation().x + hor, (p.GetLocation().y + vert) * worldX}
	return loc
}

// Mover handles all movement of players
type Mover interface {
	Move(*Players) Location
	GetType() string
	GetLocation() Location
}
