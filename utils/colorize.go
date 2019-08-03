package utils

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

func Colorize(value float64) string {
	if value >= 1.0 {
		return "#349AD2"
	}

	return "#FDFFE7"
}
