package models

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import (
	"math"

	constants "../constants"
)

// Hex is a hexagon/cell with a position vector and a scalar value which describes
// whether or not the cell is "frozen"
type Hex struct {
	Position Vector
	Value    float64
}

// Key returns a string representing the hex's location (for use when storing the Hex in Maps)
func (self Hex) Key() string {
	return self.Position.ToString()
}

// AxialCoordinates transforms the hex's position vector into an axial coordinate system,
// sometimes called "trapezoidal", "oblique", or "skewed" coordinate systems, by taking
// two of the three coordinates from the position vector (which contains cube coordinates).
// Since our cube coordinates satisfy the constraint x + y + z = 0, any two coordinates
// suffice and the third may always be derived.
func (self Hex) AxialCoordinates() Vector {
	return self.Position[0:2]
}

// IsFrozen is a predicate returning whether or not the cell is "frozen", i.e. if it's value exceeds
// a constant threshold
func (self Hex) IsFrozen() bool {
	return self.Value >= constants.Ice
}

// IsReceptive is a predicate returning whether or not the cell or any of its contiguous neighbors are frozen.
// In Reiter's model, receptive sites permanently store any mass that arrives, and mass at unreceptive sites
// is free to move (tending toward an average value).
func (self Hex) IsReceptive(neighbors []Hex) bool {
	return self.IsFrozen() || self.hasFrozenNeighbor(neighbors)
}

func (self Hex) hasFrozenNeighbor(neighbors []Hex) bool {
	for _, neighbor := range neighbors {
		if neighbor.IsFrozen() {
			return true
		}
	}

	return false
}

// Update returns a new hex with the same position and the given value
func (self Hex) Update(value float64) Hex {
	return Hex{
		Position: self.Position,
		Value:    value,
	}
}

// center returns a vector at the center of the hexagon (with coordinates in pixels)
func (self Hex) center(orientation Orientation, size Vector, origin Vector) Vector {
	var coordinates Vector

	for _, vector := range orientation.f {
		coordinates = append(coordinates, self.AxialCoordinates().Dot(vector))
	}

	return coordinates.Multiply(size).Add(origin)
}

// cornerOffset returns a vector containing offsets (in pixels) from the center of the hexagon
// to the given corner (specified as an index 0-5 inclusive)
func (self Hex) cornerOffset(corner int, orientation Orientation, size Vector) Vector {
	angle := constants.Tau * (orientation.angle + float64(corner)) / constants.TotalCorners

	return size.Multiply(Vector{math.Cos(angle), math.Sin(angle)})
}

// Corners returns the pixel values for each corner of the hexagon
func (self Hex) Corners(orientation Orientation, size Vector, origin Vector) []Vector {
	var collection []Vector
	center := self.center(orientation, size, origin)

	for i := 0; i < constants.TotalCorners; i++ {
		offset := self.cornerOffset(i, orientation, size)
		collection = append(collection, center.Add(offset))
	}

	return collection
}
