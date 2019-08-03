package models

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import (
	"math"

	constants "../constants"
)

// neighborhood contains the cubic coordinate offsets to find the neighbors of
// any given hex
var neighborhood = [6]Vector{
	Vector{1.0, 0.0, -1.0}, Vector{1.0, -1.0, 0.0}, Vector{0.0, -1.0, 1.0},
	Vector{-1.0, 0.0, 1.0}, Vector{-1.0, 1.0, 0.0}, Vector{0.0, 1.0, -1.0},
}

// HexMap maps hex positions to the hex at that position.
type HexMap map[string]Hex

// NeighborsOf returns all hexes adjacent to a given hex
func (self HexMap) NeighborsOf(hex Hex) []Hex {
	var memo []Hex

	for _, offset := range neighborhood {
		coordinates := hex.Position.Add(offset)
		neighbor, ok := self[coordinates.ToString()]

		if ok {
			memo = append(memo, neighbor)
		}
	}

	return memo
}

// stepCallback is a function handler which receives a pointer to a Hex
type stepCallback func(Hex)

// classify caches whether or not each hex is receptive for the calculation of the
// the diffusion term in Reiter's model.  Without a cache, a hex would need to
// compute the receptive state of all its neighbors, which in turn would need to
// check the receptive state of their neighbors, and so on.  With the cache, we are
// able to store this value in a single O(n) pass over the collection.
func (self HexMap) classify() Cache {
	cache := Cache{}

	for _, hex := range self {
		neighbors := self.NeighborsOf(hex)
		cache[hex.Key()] = hex.IsReceptive(neighbors)
	}

	return cache
}

// Step runs the simulation for one discrete time step and returns a new HexMap containing
// the updated values.  It also invokes a function with each new hex (allowing for rendering).
func (self HexMap) Step(fn stepCallback) HexMap {
	hexmap := HexMap{}
	cache := self.classify()

	for _, hex := range self {
		neighbors := self.NeighborsOf(hex)
		updatedHex := SimulateFlow(hex, neighbors, cache)
		hexmap[hex.Key()] = updatedHex

		fn(updatedHex)
	}

	return hexmap
}

// Hexagon returns a HexMap in the shape of a hexagon.
func Hexagon() HexMap {
	hexmap := HexMap{}
	origin := Vector{0.0, 0.0, 0.0}

	for q := -constants.Radius; q <= constants.Radius; q++ {
		r1 := math.Max(-constants.Radius, -q-constants.Radius)
		r2 := math.Min(constants.Radius, -q+constants.Radius)

		for r := r1; r <= r2; r++ {
			position := Vector{q, r, -q - r}
			coordinates := position.ToString()
			value := constants.BackgroundLevel

			if position.Is(origin) {
				value = constants.Ice
			}

			hexmap[coordinates] = Hex{Position: position, Value: value}
		}
	}

	return hexmap
}
