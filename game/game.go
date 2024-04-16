package game

import (
	"errors"
	"fmt"
)

const BOARD_W = 10
const BOARD_H = 10
const NUM_OF_BOATS = 5
    
type Boat struct {
    X int
    Y int
    Length int
    Vert bool
}

func NewBoat(x, y, length int, vert bool) Boat {
    return Boat{
        X: x,
        Y: y,
        Length: length,
        Vert: vert,
    }
}

func (b *Boat) inBounds() bool {
    if b.Vert {
        return b.Y + b.Length < BOARD_H
    }

    return b.X + b.Length < BOARD_W
}

type Player struct {
    Boats []Boat
    PlacedBoats int    
}

func NewPlayer() Player {
    return Player{
        Boats: make([]Boat, 5),
        PlacedBoats: 0,
    }
}

func (p *Player) AddBoat(x, y, length int, vert bool) (bool, error) {
    if p.PlacedBoats >= 5 {
        return false, errors.New("Already place all of the boats.")
    }

    if length > 5 || length < 2 {
        return false, errors.New("Invalid length of a ship.")
    }

    if x < 0 || y < 0 {
        return false, errors.New("Boat out of bounds.")
    }

    if (vert && y + length >= BOARD_H) || (!vert && x + length >= BOARD_W) {
        return false, errors.New("Boat out of bounds.")
    }

    boat_length_count := 0

    for i := 0; i < p.PlacedBoats; i++ {
        if p.Boats[i].Length == length {
            boat_length_count++;
        }
    }

    if boat_length_count > 0 || (length == 3 && boat_length_count < 2) {
        return false, errors.New("Already placed the boat.")
    }

    boat := NewBoat(x, y, length, vert)

    p.Boats[p.PlacedBoats] = boat
    p.PlacedBoats++

    return true, nil
} 

