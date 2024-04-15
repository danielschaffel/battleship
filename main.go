package main

import "fmt"

const BOARD_W = 10
const BOARD_H = 10
    
type Boat struct {
    X int
    Y int
    Length int
    Vert bool
}

func (b Boat) inBounds() bool {
    if b.Vert {
        return b.Y + b.Length < BOARD_H
    }

    return b.X + b.Length < BOARD_W
}

func main() {
    fmt.Printf("Width %d, Height %d\n", BOARD_H, BOARD_H)
}
