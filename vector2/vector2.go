package vector2

import (
	"errors"
	"math"
)

type Vector2 struct {
	StartCoordinate Coordinate
	EndCoordinate   Coordinate
	Direction       int
}

// Creates and returns a vector from a set of input data
func Create(startX, startY, endX, endY float64, direction int) Vector2 {
	vector := Vector2{
		StartCoordinate: Coordinate{
			X: startX, Y: startY,
		},
		EndCoordinate: Coordinate{
			X: endX, Y: endY,
		},
		Direction: direction,
	}

	return vector
}

var ZeroVector2 = Vector2{
	StartCoordinate: ZeroCoordinate,
	EndCoordinate:   ZeroCoordinate,
	Direction:       0,
}

// Returns the vector length
func (v Vector2) Len() float64 {
	if v.StartCoordinate.X == v.EndCoordinate.X {
		return math.Abs(v.StartCoordinate.Y - v.EndCoordinate.Y)
	}

	return math.Abs(v.StartCoordinate.X - v.EndCoordinate.X)
}

// Checks the incoming array of vectors to represent a single direction and returns a motion vector that is valid for that array.
// Returns a null vector and an error if the vector array does not represent a single direction.
// Returns an incoming vector if it represents an array of one element.
func GetDisplacementVector(vectors ...Vector2) (Vector2, error) {
	for i, v := range vectors {
		if v.EndCoordinate.X != vectors[i-1].StartCoordinate.X {
			return ZeroVector2, errors.New("vectors do not represent a single direction")
		}
	}

	if len(vectors) == 1 {
		return vectors[0], nil
	}

	dispVector := Vector2{
		StartCoordinate: vectors[0].StartCoordinate,
		EndCoordinate:   vectors[len(vectors)-1].EndCoordinate,
	}
	return dispVector, nil
}

// Checks two vectors for equality.
func Equal(vector1 Vector2, vector2 Vector2) bool {
	if vector1.Len() == vector2.Len() {
		return true
	}

	return false
}

// Checks two vectors for opposing each other.
func Opposite(vector1 Vector2, vector2 Vector2) bool {
	if !Equal(vector1, vector2) {
		return false
	} else if vector2.Direction != -vector1.Direction {
		return false
	}

	return true
}
