package models

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import constants "../constants"

// Orientation contains information needed to compute the position of
// a given hex's center and corners for some orientation (either a vertex
// or a side pointing North)
type Orientation struct {
	f     [2]Vector
	b     [2]Vector
	angle float64
}

// Pointy returns the offsets for an orientation where the hexagon vertex points North
func Pointy() Orientation {
	f := [2]Vector{}
	f[0] = append(f[0], constants.Sqrt3)
	f[0] = append(f[0], constants.Sqrt3/2.0)
	f[1] = append(f[1], 0.0)
	f[1] = append(f[1], 3.0/2.0)

	b := [2]Vector{}
	f[0] = append(f[0], constants.Sqrt3/3.0)
	f[0] = append(f[0], -1.0/3.0)
	f[1] = append(f[1], 0.0)
	f[1] = append(f[1], 2.0/3.0)

	return Orientation{f: f, b: b, angle: 0.5}
}

// Flat returns the offsets for an orientation where the hexagon side points North
func Flat() Orientation {
	f := [2]Vector{}
	f[0] = append(f[0], 3.0/2.0)
	f[0] = append(f[0], 0.0)
	f[1] = append(f[1], constants.Sqrt3/2.0)
	f[1] = append(f[1], constants.Sqrt3)

	b := [2]Vector{}
	f[0] = append(f[0], 2.0/3.0)
	f[0] = append(f[0], 0.0)
	f[1] = append(f[1], -1.0/3.0)
	f[1] = append(f[1], constants.Sqrt3/3.0)

	return Orientation{f: f, b: b, angle: 0.0}
}
