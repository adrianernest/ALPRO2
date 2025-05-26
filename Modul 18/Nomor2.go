package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tile struct {
	sideOne  int
	sideTwo  int
	value    int
	isDouble bool
}

type TileSet struct {
	tiles [28]Tile
	count int
}

func initializeTiles(ts *TileSet) {
	idx := 0
	for a := 0; a <= 6; a++ {
		for b := a; b <= 6; b++ {
			t := Tile{
				sideOne:  a,
				sideTwo:  b,
				value:    a + b,
				isDouble: a == b,
			}
			ts.tiles[idx] = t
			idx++
		}
	}
	ts.count = 28
}

func shuffleTiles(ts *TileSet) {
	rand.Seed(time.Now().UnixNano())
	for i := ts.count - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		ts.tiles[i], ts.tiles[j] = ts.tiles[j], ts.tiles[i]
	}
}

func drawTile(ts *TileSet) Tile {
	if ts.count == 0 {
		return Tile{-1, -1, -1, false}
	}
	ts.count--
	return ts.tiles[ts.count]
}

func getTileSide(t Tile, side int) int {
	if side == 1 {
		return t.sideOne
	} else if side == 2 {
		return t.sideTwo
	} else {
		return -1
	}
}

func getTileValue(t Tile) int {
	return t.value
}

func isMatch(t1, t2 Tile) bool {
	return t1.sideOne == t2.sideOne ||
		t1.sideOne == t2.sideTwo ||
		t1.sideTwo == t2.sideOne ||
		t1.sideTwo == t2.sideTwo
}

func findMatchingTile(ts *TileSet, reference Tile) Tile {
	fmt.Println("\nTile drawn: ")
	for ts.count > 0 {
		tile := drawTile(ts)
		fmt.Printf("Drawn: (%d,%d)\n", tile.sideOne, tile.sideTwo)
		if isMatch(tile, reference) {
			return tile
		}
	}
	return Tile{-1, -1, -1, false}
}

func main() {
	var tileSet TileSet

	initializeTiles(&tileSet)
	shuffleTiles(&tileSet)

	referenceTile := drawTile(&tileSet)
	fmt.Printf("Reference tile: (%d,%d)\n", referenceTile.sideOne, referenceTile.sideTwo)

	matchingTile := findMatchingTile(&tileSet, referenceTile)

	if matchingTile.sideOne == -1 {
		fmt.Println("No matching tile found.")
	} else {
		fmt.Printf("\nMatching tile: (%d,%d)\n", matchingTile.sideOne, matchingTile.sideTwo)
	}
}
