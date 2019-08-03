package tests

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import (
	"testing"

	models "../models"
)

func TestHexAxialCoordinates(t *testing.T) {
	hex := models.Hex{
		Position: models.Vector{1.0, 2.0, 3.0},
		Value:    1.0,
	}
	result := hex.AxialCoordinates()
	expected := models.Vector{1.0, 2.0}

	if !expected.Is(result) {
		t.Error("Hex axial coordinates are not the first two cube coordinates from the position vector")
	}
}

func TestHexIsFrozen(t *testing.T) {
	position := models.Vector{1.0, 2.0, 3.0}
	unfrozen := models.Hex{Position: position, Value: 0.9}
	frozen := models.Hex{Position: position, Value: 1.0}

	if !frozen.IsFrozen() {
		t.Error("Frozen hex is reported as unfrozen")
	}

	if unfrozen.IsFrozen() {
		t.Error("Unfrozen hex is reported as frozen")
	}
}

func TestHexIsReceptiveWhenCellIsFrozen(t *testing.T) {
	var neighborhood []models.Hex
	position := models.Vector{1.0, 2.0, 3.0}
	frozen := models.Hex{Position: position, Value: 1.0}

	if !frozen.IsReceptive(neighborhood) {
		t.Error("Frozen hex is not receptive")
	}
}

func TestHexIsNotReceptiveWhenCellIsNotfrozen(t *testing.T) {
	var neighborhood []models.Hex
	position := models.Vector{1.0, 2.0, 3.0}
	unfrozen := models.Hex{Position: position, Value: 0.1}

	if unfrozen.IsReceptive(neighborhood) {
		t.Error("Unfrozen hex is receptive")
	}
}

func TestHexIsReceptiveWhenNeighborIsFrozen(t *testing.T) {
	var neighborhood []models.Hex
	position := models.Vector{1.0, 2.0, 3.0}
	unfrozen := models.Hex{Position: position, Value: 0.1}
	frozen := models.Hex{Position: position, Value: 1.0}
	neighborhood = append(neighborhood, frozen)

	if !unfrozen.IsReceptive(neighborhood) {
		t.Error("Unfrozen hex with frozen neighbor is not receptive")
	}
}

func TestHexIsNotReceptiveWhenNeighborIsNotFrozen(t *testing.T) {
	var neighborhood []models.Hex
	position := models.Vector{1.0, 2.0, 3.0}
	unfrozen := models.Hex{Position: position, Value: 0.1}
	neighbor := models.Hex{Position: position, Value: 0.1}
	neighborhood = append(neighborhood, neighbor)

	if unfrozen.IsReceptive(neighborhood) {
		t.Error("Unfrozen hex with unfrozen neighbor is receptive")
	}
}
