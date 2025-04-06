package models

import (
	"encoding/json"
	"fmt"
)

type Color string

const (
	White  Color = "white"
	Yellow Color = "yellow"
	Red    Color = "red"
	Orange Color = "orange"
	Blue   Color = "blue"
	Green  Color = "green"
)

type Face [3][3]Color

type RubiksCube struct {
	Up    Face `json:"up"`
	Down  Face `json:"down"`
	Front Face `json:"front"`
	Back  Face `json:"back"`
	Left  Face `json:"left"`
	Right Face `json:"right"`
}

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

func initFace(color Color) Face {
	var face Face
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			face[i][j] = color
		}
	}
	return face
}

func (c *RubiksCube) Reset() {
	*c = *New()
}

func rotateFaceClockwise(face Face) Face {
	var newFace Face
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newFace[i][j] = face[2-j][i]
		}
	}
	return newFace
}

func rotateFaceCounterClockwise(face Face) Face {
	var newFace Face
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newFace[i][j] = face[j][2-i]
		}
	}
	return newFace
}

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

func (c *RubiksCube) RotateFront(clockwise bool) {
	if clockwise {
		c.Front = rotateFaceClockwise(c.Front)
	} else {
		c.Front = rotateFaceCounterClockwise(c.Front)
	}

	var temp [3]Color

	for i := 0; i < 3; i++ {
		temp[i] = c.Up[2][i]
	}

	if clockwise {
		for i := 0; i < 3; i++ {
			c.Up[2][i] = c.Left[2-i][2]
		}

		for i := 0; i < 3; i++ {
			c.Left[i][2] = c.Down[0][i]
		}

		for i := 0; i < 3; i++ {
			c.Down[0][i] = c.Right[2-i][0]
		}

		for i := 0; i < 3; i++ {
			c.Right[i][0] = temp[2-i]
		}
	} else {
		for i := 0; i < 3; i++ {
			c.Up[2][i] = c.Right[i][0]
		}

		for i := 0; i < 3; i++ {
			c.Right[i][0] = c.Down[0][2-i]
		}

		for i := 0; i < 3; i++ {
			c.Down[0][i] = c.Left[i][2]
		}

		for i := 0; i < 3; i++ {
			c.Left[i][2] = temp[2-i]
		}
	}
}

func (c *RubiksCube) RotateBack(clockwise bool) {
	if clockwise {
		c.Back = rotateFaceClockwise(c.Back)
	} else {
		c.Back = rotateFaceCounterClockwise(c.Back)
	}

	var temp [3]Color

	for i := 0; i < 3; i++ {
		temp[i] = c.Up[0][i]
	}

	if clockwise {
		for i := 0; i < 3; i++ {
			c.Up[0][i] = c.Right[i][2]
		}

		for i := 0; i < 3; i++ {
			c.Right[i][2] = c.Down[2][2-i]
		}

		for i := 0; i < 3; i++ {
			c.Down[2][i] = c.Left[i][0]
		}

		for i := 0; i < 3; i++ {
			c.Left[i][0] = temp[2-i]
		}
	} else {
		for i := 0; i < 3; i++ {
			c.Up[0][i] = c.Left[2-i][0]
		}

		for i := 0; i < 3; i++ {
			c.Left[i][0] = c.Down[2][i]
		}

		for i := 0; i < 3; i++ {
			c.Down[2][i] = c.Right[2-i][2]
		}

		for i := 0; i < 3; i++ {
			c.Right[i][2] = temp[i]
		}
	}
}

func (c *RubiksCube) RotateUp(clockwise bool) {
	if clockwise {
		c.Up = rotateFaceClockwise(c.Up)
	} else {
		c.Up = rotateFaceCounterClockwise(c.Up)
	}

	var temp [3]Color

	for i := 0; i < 3; i++ {
		temp[i] = c.Front[0][i]
	}

	if clockwise {
		for i := 0; i < 3; i++ {
			c.Front[0][i] = c.Right[0][i]
		}

		for i := 0; i < 3; i++ {
			c.Right[0][i] = c.Back[0][i]
		}

		for i := 0; i < 3; i++ {
			c.Back[0][i] = c.Left[0][i]
		}

		for i := 0; i < 3; i++ {
			c.Left[0][i] = temp[i]
		}
	} else {
		for i := 0; i < 3; i++ {
			c.Front[0][i] = c.Left[0][i]
		}

		for i := 0; i < 3; i++ {
			c.Left[0][i] = c.Back[0][i]
		}

		for i := 0; i < 3; i++ {
			c.Back[0][i] = c.Right[0][i]
		}

		for i := 0; i < 3; i++ {
			c.Right[0][i] = temp[i]
		}
	}
}

func (c *RubiksCube) RotateDown(clockwise bool) {
	if clockwise {
		c.Down = rotateFaceClockwise(c.Down)
	} else {
		c.Down = rotateFaceCounterClockwise(c.Down)
	}

	var temp [3]Color

	for i := 0; i < 3; i++ {
		temp[i] = c.Front[2][i]
	}

	if clockwise {
		for i := 0; i < 3; i++ {
			c.Front[2][i] = c.Left[2][i]
		}

		for i := 0; i < 3; i++ {
			c.Left[2][i] = c.Back[2][i]
		}

		for i := 0; i < 3; i++ {
			c.Back[2][i] = c.Right[2][i]
		}

		for i := 0; i < 3; i++ {
			c.Right[2][i] = temp[i]
		}
	} else {
		for i := 0; i < 3; i++ {
			c.Front[2][i] = c.Right[2][i]
		}

		for i := 0; i < 3; i++ {
			c.Right[2][i] = c.Back[2][i]
		}

		for i := 0; i < 3; i++ {
			c.Back[2][i] = c.Left[2][i]
		}

		for i := 0; i < 3; i++ {
			c.Left[2][i] = temp[i]
		}
	}
}

func (c *RubiksCube) RotateLeft(clockwise bool) {
	if clockwise {
		c.Left = rotateFaceClockwise(c.Left)
	} else {
		c.Left = rotateFaceCounterClockwise(c.Left)
	}

	var temp [3]Color

	for i := 0; i < 3; i++ {
		temp[i] = c.Up[i][0]
	}

	if clockwise {
		for i := 0; i < 3; i++ {
			c.Up[i][0] = c.Back[2-i][2]
		}

		for i := 0; i < 3; i++ {
			c.Back[2-i][2] = c.Down[i][0]
		}

		for i := 0; i < 3; i++ {
			c.Down[i][0] = c.Front[i][0]
		}

		for i := 0; i < 3; i++ {
			c.Front[i][0] = temp[i]
		}
	} else {
		for i := 0; i < 3; i++ {
			c.Up[i][0] = c.Front[i][0]
		}

		for i := 0; i < 3; i++ {
			c.Front[i][0] = c.Down[i][0]
		}

		for i := 0; i < 3; i++ {
			c.Down[i][0] = c.Back[2-i][2]
		}

		for i := 0; i < 3; i++ {
			c.Back[2-i][2] = temp[i]
		}
	}
}

func (c *RubiksCube) RotateRight(clockwise bool) {
	if clockwise {
		c.Right = rotateFaceClockwise(c.Right)
	} else {
		c.Right = rotateFaceCounterClockwise(c.Right)
	}

	var temp [3]Color

	for i := 0; i < 3; i++ {
		temp[i] = c.Up[i][2]
	}

	if clockwise {
		for i := 0; i < 3; i++ {
			c.Up[i][2] = c.Front[i][2]
		}

		for i := 0; i < 3; i++ {
			c.Front[i][2] = c.Down[i][2]
		}

		for i := 0; i < 3; i++ {
			c.Down[i][2] = c.Back[2-i][0]
		}

		for i := 0; i < 3; i++ {
			c.Back[2-i][0] = temp[i]
		}
	} else {
		for i := 0; i < 3; i++ {
			c.Up[i][2] = c.Back[2-i][0]
		}

		for i := 0; i < 3; i++ {
			c.Back[2-i][0] = c.Down[i][2]
		}

		for i := 0; i < 3; i++ {
			c.Down[i][2] = c.Front[i][2]
		}

		for i := 0; i < 3; i++ {
			c.Front[i][2] = temp[i]
		}
	}
}

func (c *RubiksCube) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}

func (c *RubiksCube) Move(notation string) error {
	if len(notation) == 0 {
		return fmt.Errorf("empty move notation")
	}

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

	clockwise := true
	if len(notation) > 1 {
		if notation[1] == '\'' {
			clockwise = false
		} else if notation[1] == '2' {
			err := c.RotateFace(face, clockwise)
			if err != nil {
				return err
			}
			return c.RotateFace(face, clockwise)
		}
	}

	return c.RotateFace(face, clockwise)
}

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
