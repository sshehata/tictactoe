package game

import "errors"

// Tile tictactoe tile for one player
type Tile int

const (
	// Undefined place holder tile
	Undefined Tile = iota
	// OTile tile for first player with symbol O
	OTile
	// XTile tile for second player with symbol X
	XTile
)

// NewTile convert int to tile
func NewTile(t int) (Tile, error) {
	if Tile(t) < Undefined || Tile(t) > XTile {
		return Undefined, errors.New("%v is not a valid tile")
	}
	return Tile(t), nil
}

func (t Tile) String() string {
	switch t {
	case OTile:
		return "O"
	case XTile:
		return "X"
	default:
		return "-"
	}
}
