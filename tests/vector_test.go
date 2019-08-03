package tests

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import (
	"testing"

	models "../models"
)

func TestVectorXY(t *testing.T) {
	vector := models.Vector{1, 2, 3}

	if vector.X() != 1 {
		t.Error("Vector X does not return the first element of the vector")
	}

	if vector.Y() != 2 {
		t.Error("Vector Y does not return the second element of the vector")
	}
}

func TestVectorIs(t *testing.T) {
	a := models.Vector{1.0, 2.0, 3.0}
	b := models.Vector{1.0, 2.0, 3.0}
	c := models.Vector{1.0, 2.0, 1.0}

	if !a.Is(b) {
		t.Error("Vector equivalence is incorrect when two vectors are equivalent")
	}

	if a.Is(c) {
		t.Error("Vector equivalence is incorrect when two vectors are not equivalent")
	}
}

func TestVectorAdd(t *testing.T) {
	a := models.Vector{1.0, 2.0, 3.0}
	b := models.Vector{3.0, 2.0, 1.0}
	expected := models.Vector{4.0, 4.0, 4.0}
	resultant := a.Add(b)

	if !expected.Is(resultant) {
		t.Error("Vector addition returns an incorrect result")
	}
}

func TestVectorMultiply(t *testing.T) {
	a := models.Vector{3.0, 2.0}
	b := models.Vector{2.0, 3.0}
	expected := models.Vector{6.0, 6.0}
	resultant := a.Multiply(b)

	if !expected.Is(resultant) {
		t.Error("Vector multiplication returns an incorrect result")
	}
}

func TestVectorDot(t *testing.T) {
	a := models.Vector{2.0, 3.0}
	b := models.Vector{3.0, 2.0}
	expected := 12.0
	result := a.Dot(b)

	if result != expected {
		t.Error("Vector dot product returns an incorrect result")
	}
}

func TestVectorToString(t *testing.T) {
	a := models.Vector{1.0, 2.0, 2.0}
	result := a.ToString()

	if result != "1,2,2" {
		t.Error("Vector to string returns an incorrect string")
	}
}
