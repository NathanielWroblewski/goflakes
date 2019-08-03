package models

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import "fmt"

// Vector is a one-dimensional array of floats representing
// an element of Euclidean vector space
type Vector []float64

// X is a convenience method for accessing the first element
func (self Vector) X() float64 {
	return self[0]
}

// Y is a convenience method for accessing the second element
func (self Vector) Y() float64 {
	return self[1]
}

// Is is a predicate which tests for equivalence
func (self Vector) Is(vector Vector) bool {
	for index, element := range self {
		if element != vector[index] {
			return false
		}
	}
	return true
}

// Add performs vector addition by summing elements pair-wise
func (self Vector) Add(vector Vector) Vector {
	resultant := Vector{}

	for index, element := range self {
		resultant = append(resultant, element+vector[index])
	}

	return resultant
}

// Multiply performs vector multiplication by multiplying elements pair-wise
func (self Vector) Multiply(vector Vector) Vector {
	resultant := Vector{}

	for index, element := range self {
		resultant = append(resultant, element*vector[index])
	}

	return resultant
}

// Dot takes the dot product of two vectors (returns a scalar)
func (self Vector) Dot(vector Vector) float64 {
	sum := 0.0

	for index, element := range self {
		sum += element * vector[index]
	}

	return sum
}

// ToString returns a comma-delimited string of the vector elements coerced to integers
func (self Vector) ToString() string {
	return fmt.Sprintf("%d,%d,%d", int(self[0]), int(self[1]), int(self[2]))
}
