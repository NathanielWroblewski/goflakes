package models

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

// Cache maps a hex location to a boolean representing whether or not the hex was receptive
type Cache map[string]bool

// WasReceptive returns the cached receptive state for a given hex
func (self Cache) WasReceptive(hex Hex) bool {
	return self[hex.Key()]
}
