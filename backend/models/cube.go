package models

import (
	"encoding/json"
	"fmt"
)

// Color represents the possible colors of a Rubik's Cube
type Color string

const (
	White  Color = "white"
	Yellow Color = "yellow"
	Red    Color = "red"
	Orange Color = "orange"
	Blue   Color = "blue"
	Green  Color = "green"
)

// Face represents a single face of the cube (3x3 grid)
type Face [3][3]Color

// RubiksCube represents the entire Rubik's Cube with 6 faces
type RubiksCube struct {
	Up    Face `json:"up"`
	Down  Face `json:"down"`
	Front Face `json:"front"`
	Back  Face `json:"back"`
	Left  Face `json:"left"`
	Right Face `json:"right"`
}

// New creates a new solved Rubik's Cube
func New() *RubiksCube {
	cube := &RubiksCube{
		Up:    initFace(White),
		Down:  initFace(Yellow),
		Front: initFace(Green),
		Back:  initFace(Blue),
		Left:  initFace(Orange),
		Right: initFace(Red),
	}
	return cube
}

// initFace initializes a face with a single color
func initFace(color Color) Face {
	var face Face
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			face[i][j] = color
		}
	}
	return face
}

// Reset resets the cube to its solved state
func (c *RubiksCube) Reset() {
	*c = *New()
}

// rotateFaceClockwise rotates a face 90 degrees clockwise
func rotateFaceClockwise(face Face) Face {
	var newFace Face
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newFace[i][j] = face[2-j][i]
		}
	}
	return newFace
}

// rotateFaceCounterClockwise rotates a face 90 degrees counter-clockwise
func rotateFaceCounterClockwise(face Face) Face {
	var newFace Face
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newFace[i][j] = face[j][2-i]
		}
	}
	return newFace
}

// RotateFace rotates a face and updates affected edges
func (c *RubiksCube) RotateFace(face string, clockwise bool) error {
	switch face {
	case "front":
		c.RotateFront(clockwise)
	case "back":
		c.RotateBack(clockwise)
	case "up":
		c.RotateUp(clockwise)
	case "down":
		c.RotateDown(clockwise)
	case "left":
		c.RotateLeft(clockwise)
	case "right":
		c.RotateRight(clockwise)
	default:
		return fmt.Errorf("invalid face: %s", face)
	}
	return nil
}

// RotateFront rotates the front face
func (c *RubiksCube) RotateFront(clockwise bool) {
	// 1. Rotate the front face itself
	if clockwise {
		c.Front = rotateFaceClockwise(c.Front)
	} else {
		c.Front = rotateFaceCounterClockwise(c.Front)
	}

	// 2. Update the affected edges
	var temp [3]Color

	// Save the top edge
	for i := 0; i < 3; i++ {
		temp[i] = c.Up[2][i]
	}

	if clockwise {
		// Move left to top
		for i := 0; i < 3; i++ {
			c.Up[2][i] = c.Left[2-i][2]
		}

		// Move down to left
		for i := 0; i < 3; i++ {
			c.Left[i][2] = c.Down[0][i]
		}

		// Move right to down
		for i := 0; i < 3; i++ {
			c.Down[0][i] = c.Right[2-i][0]
		}

		// Move saved top to right
		for i := 0; i < 3; i++ {
			c.Right[i][0] = temp[2-i]
		}
	} else {
		// Move right to top
		for i := 0; i < 3; i++ {
			c.Up[2][i] = c.Right[i][0]
		}

		// Move down to right
		for i := 0; i < 3; i++ {
			c.Right[i][0] = c.Down[0][2-i]
		}

		// Move left to down
		for i := 0; i < 3; i++ {
			c.Down[0][i] = c.Left[i][2]
		}

		// Move saved top to left
		for i := 0; i < 3; i++ {
			c.Left[i][2] = temp[2-i]
		}
	}
}

// RotateBack rotates the back face
func (c *RubiksCube) RotateBack(clockwise bool) {
	// 1. Rotate the back face itself
	if clockwise {
		c.Back = rotateFaceClockwise(c.Back)
	} else {
		c.Back = rotateFaceCounterClockwise(c.Back)
	}

	// 2. Update the affected edges
	var temp [3]Color

	// Save the top edge
	for i := 0; i < 3; i++ {
		temp[i] = c.Up[0][i]
	}

	if clockwise {
		// Move right to top
		for i := 0; i < 3; i++ {
			c.Up[0][i] = c.Right[i][2]
		}

		// Move down to right
		for i := 0; i < 3; i++ {
			c.Right[i][2] = c.Down[2][2-i]
		}

		// Move left to down
		for i := 0; i < 3; i++ {
			c.Down[2][i] = c.Left[i][0]
		}

		// Move saved top to left
		for i := 0; i < 3; i++ {
			c.Left[i][0] = temp[2-i]
		}
	} else {
		// Move left to top
		for i := 0; i < 3; i++ {
			c.Up[0][i] = c.Left[2-i][0]
		}

		// Move down to left
		for i := 0; i < 3; i++ {
			c.Left[i][0] = c.Down[2][i]
		}

		// Move right to down
		for i := 0; i < 3; i++ {
			c.Down[2][i] = c.Right[2-i][2]
		}

		// Move saved top to right
		for i := 0; i < 3; i++ {
			c.Right[i][2] = temp[i]
		}
	}
}

// RotateUp rotates the top face
func (c *RubiksCube) RotateUp(clockwise bool) {
	// 1. Rotate the up face itself
	if clockwise {
		c.Up = rotateFaceClockwise(c.Up)
	} else {
		c.Up = rotateFaceCounterClockwise(c.Up)
	}

	// 2. Update the affected edges
	var temp [3]Color

	// Save front edge
	for i := 0; i < 3; i++ {
		temp[i] = c.Front[0][i]
	}

	if clockwise {
		// Move right to front
		for i := 0; i < 3; i++ {
			c.Front[0][i] = c.Right[0][i]
		}

		// Move back to right
		for i := 0; i < 3; i++ {
			c.Right[0][i] = c.Back[0][i]
		}

		// Move left to back
		for i := 0; i < 3; i++ {
			c.Back[0][i] = c.Left[0][i]
		}

		// Move saved front to left
		for i := 0; i < 3; i++ {
			c.Left[0][i] = temp[i]
		}
	} else {
		// Move left to front
		for i := 0; i < 3; i++ {
			c.Front[0][i] = c.Left[0][i]
		}

		// Move back to left
		for i := 0; i < 3; i++ {
			c.Left[0][i] = c.Back[0][i]
		}

		// Move right to back
		for i := 0; i < 3; i++ {
			c.Back[0][i] = c.Right[0][i]
		}

		// Move saved front to right
		for i := 0; i < 3; i++ {
			c.Right[0][i] = temp[i]
		}
	}
}

// RotateDown rotates the bottom face
func (c *RubiksCube) RotateDown(clockwise bool) {
	// 1. Rotate the down face itself
	if clockwise {
		c.Down = rotateFaceClockwise(c.Down)
	} else {
		c.Down = rotateFaceCounterClockwise(c.Down)
	}

	// 2. Update the affected edges
	var temp [3]Color

	// Save front edge
	for i := 0; i < 3; i++ {
		temp[i] = c.Front[2][i]
	}

	if clockwise {
		// Move left to front
		for i := 0; i < 3; i++ {
			c.Front[2][i] = c.Left[2][i]
		}

		// Move back to left
		for i := 0; i < 3; i++ {
			c.Left[2][i] = c.Back[2][i]
		}

		// Move right to back
		for i := 0; i < 3; i++ {
			c.Back[2][i] = c.Right[2][i]
		}

		// Move saved front to right
		for i := 0; i < 3; i++ {
			c.Right[2][i] = temp[i]
		}
	} else {
		// Move right to front
		for i := 0; i < 3; i++ {
			c.Front[2][i] = c.Right[2][i]
		}

		// Move back to right
		for i := 0; i < 3; i++ {
			c.Right[2][i] = c.Back[2][i]
		}

		// Move left to back
		for i := 0; i < 3; i++ {
			c.Back[2][i] = c.Left[2][i]
		}

		// Move saved front to left
		for i := 0; i < 3; i++ {
			c.Left[2][i] = temp[i]
		}
	}
}

// RotateLeft rotates the left face
func (c *RubiksCube) RotateLeft(clockwise bool) {
	// 1. Rotate the left face itself
	if clockwise {
		c.Left = rotateFaceClockwise(c.Left)
	} else {
		c.Left = rotateFaceCounterClockwise(c.Left)
	}

	// 2. Update the affected edges
	var temp [3]Color

	// Save up edge
	for i := 0; i < 3; i++ {
		temp[i] = c.Up[i][0]
	}

	if clockwise {
		// SWAPPED: This block was previously in the counterclockwise case
		// Move back to up (with rotation)
		for i := 0; i < 3; i++ {
			c.Up[i][0] = c.Back[2-i][2]
		}

		// Move down to back (with rotation)
		for i := 0; i < 3; i++ {
			c.Back[2-i][2] = c.Down[i][0]
		}

		// Move front to down
		for i := 0; i < 3; i++ {
			c.Down[i][0] = c.Front[i][0]
		}

		// Move saved up to front
		for i := 0; i < 3; i++ {
			c.Front[i][0] = temp[i]
		}
	} else {
		// SWAPPED: This block was previously in the clockwise case
		// Move front to up
		for i := 0; i < 3; i++ {
			c.Up[i][0] = c.Front[i][0]
		}

		// Move down to front
		for i := 0; i < 3; i++ {
			c.Front[i][0] = c.Down[i][0]
		}

		// Move back to down (with rotation)
		for i := 0; i < 3; i++ {
			c.Down[i][0] = c.Back[2-i][2]
		}

		// Move saved up to back (with rotation)
		for i := 0; i < 3; i++ {
			c.Back[2-i][2] = temp[i]
		}
	}
}

// RotateRight rotates the right face
func (c *RubiksCube) RotateRight(clockwise bool) {
	// 1. Rotate the right face itself - keeping this part unchanged
	if clockwise {
		c.Right = rotateFaceClockwise(c.Right)
	} else {
		c.Right = rotateFaceCounterClockwise(c.Right)
	}

	// 2. Update the affected edges
	var temp [3]Color

	// Save up edge
	for i := 0; i < 3; i++ {
		temp[i] = c.Up[i][2]
	}

	if clockwise {
		// SWAPPED: This block was previously in the counterclockwise case
		// Move front to up
		for i := 0; i < 3; i++ {
			c.Up[i][2] = c.Front[i][2]
		}

		// Move down to front
		for i := 0; i < 3; i++ {
			c.Front[i][2] = c.Down[i][2]
		}

		// Move back to down (with rotation)
		for i := 0; i < 3; i++ {
			c.Down[i][2] = c.Back[2-i][0]
		}

		// Move saved up to back (with rotation)
		for i := 0; i < 3; i++ {
			c.Back[2-i][0] = temp[i]
		}
	} else {
		// SWAPPED: This block was previously in the clockwise case
		// Move back to up (with rotation)
		for i := 0; i < 3; i++ {
			c.Up[i][2] = c.Back[2-i][0]
		}

		// Move down to back (with rotation)
		for i := 0; i < 3; i++ {
			c.Back[2-i][0] = c.Down[i][2]
		}

		// Move front to down
		for i := 0; i < 3; i++ {
			c.Down[i][2] = c.Front[i][2]
		}

		// Move saved up to front
		for i := 0; i < 3; i++ {
			c.Front[i][2] = temp[i]
		}
	}
}

// String returns a string representation of the cube
func (c *RubiksCube) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}

// Move implements a standard cube notation move (F, B, U, D, L, R, F', B', etc.)
func (c *RubiksCube) Move(notation string) error {
	if len(notation) == 0 {
		return fmt.Errorf("empty move notation")
	}

	// Determine the face
	var face string
	switch notation[0] {
	case 'F':
		face = "front"
	case 'B':
		face = "back"
	case 'U':
		face = "up"
	case 'D':
		face = "down"
	case 'L':
		face = "left"
	case 'R':
		face = "right"
	default:
		return fmt.Errorf("invalid move notation: %s", notation)
	}

	// Determine direction (clockwise by default)
	clockwise := true
	if len(notation) > 1 {
		if notation[1] == '\'' {
			clockwise = false
		} else if notation[1] == '2' {
			// Double move
			err := c.RotateFace(face, clockwise)
			if err != nil {
				return err
			}
			return c.RotateFace(face, clockwise)
		}
	}

	return c.RotateFace(face, clockwise)
}

// GetColorScheme returns the current color scheme
func (c *RubiksCube) GetColorScheme() map[string]string {
	return map[string]string{
		"up":    string(c.Up[1][1]),
		"down":  string(c.Down[1][1]),
		"front": string(c.Front[1][1]),
		"back":  string(c.Back[1][1]),
		"left":  string(c.Left[1][1]),
		"right": string(c.Right[1][1]),
	}
}

// Clone creates a deep copy of the cube
func (c *RubiksCube) Clone() *RubiksCube {
	clone := new(RubiksCube)

	// Copy all face data
	clone.Up = c.Up
	clone.Down = c.Down
	clone.Front = c.Front
	clone.Back = c.Back
	clone.Left = c.Left
	clone.Right = c.Right

	return clone
}
