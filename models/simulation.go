package models

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import constants "../constants"

func collectContributions(hex Hex, neighbors []Hex, cache Cache) Vector {
	var contributions Vector

	for i := 0; i < constants.TotalCorners; i++ {
		if i > len(neighbors)-1 {
			contributions = append(contributions, constants.BackgroundLevel)
		} else {
			neighbor := neighbors[i]
			contributions = append(contributions, contribution(neighbor, cache))
		}
	}

	return contributions
}

func diffusionTerm(hex Hex, neighbors []Hex, cache Cache) float64 {
	contributions := collectContributions(hex, neighbors, cache)
	weightedAverage := 0.0

	for i := 0; i < constants.TotalCorners; i++ {
		weightedAverage += (1.0 / 12.0 * contributions[i])
	}

	return 0.5*contribution(hex, cache) + weightedAverage
}

func contribution(hex Hex, cache Cache) float64 {
	if cache.WasReceptive(hex) {
		return 0.0
	}

	return hex.Value
}

func gamma(hex Hex, cache Cache) float64 {
	if cache.WasReceptive(hex) {
		return hex.Value + constants.Additive
	}

	return 0.0
}

// SimulateFlow calculates the next value for a hex according to the rules of Reiter's simulation
func SimulateFlow(hex Hex, neighbors []Hex, cache Cache) Hex {
	return hex.Update(gamma(hex, cache) + diffusionTerm(hex, neighbors, cache))
}
