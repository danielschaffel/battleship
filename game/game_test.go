package game

import (
	"testing"
)

func TestAddBoat(t *testing.T) {
    p := NewPlayer()
    var _, err = p.AddBoat(0, 0, 2, true)

    if err != nil {
        t.Error(err)
        t.Fail()
    }

    if p.PlacedBoats != 1 {
        t.Error("Boat wasn't actually placed")
        t.Fail()
    }
}

func TestAddBoatOfSameSizeNotLengthThree(t *testing.T) {
    p := NewPlayer()
    var _, _ = p.AddBoat(0, 0, 2, true)
    var _, err = p.AddBoat(0, 0, 2, true)
    

    if err.Error() != "Already placed the boat." {
        t.Error(err.Error())
        t.Fail()
    }
}
