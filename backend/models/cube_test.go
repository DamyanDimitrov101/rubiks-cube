package models

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewCube(t *testing.T) {
	cube := New()

	assertSolidFace(t, cube.Up, White, "Up face should be white")
	assertSolidFace(t, cube.Down, Yellow, "Down face should be yellow")
	assertSolidFace(t, cube.Front, Green, "Front face should be green")
	assertSolidFace(t, cube.Back, Blue, "Back face should be blue")
	assertSolidFace(t, cube.Left, Orange, "Left face should be orange")
	assertSolidFace(t, cube.Right, Red, "Right face should be red")
}

func assertSolidFace(t *testing.T, face Face, expectedColor Color, message string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if face[i][j] != expectedColor {
				t.Errorf("%s, but found %s at position [%d][%d]", message, face[i][j], i, j)
			}
		}
	}
}

func TestReset(t *testing.T) {
	cube := New()

	cube.Move("F")
	cube.Move("R")
	cube.Move("U")

	cube.Reset()

	assertSolidFace(t, cube.Up, White, "Up face should be white after reset")
	assertSolidFace(t, cube.Down, Yellow, "Down face should be yellow after reset")
	assertSolidFace(t, cube.Front, Green, "Front face should be green after reset")
	assertSolidFace(t, cube.Back, Blue, "Back face should be blue after reset")
	assertSolidFace(t, cube.Left, Orange, "Left face should be orange after reset")
	assertSolidFace(t, cube.Right, Red, "Right face should be red after reset")
}

func TestGetColorScheme(t *testing.T) {
	cube := New()
	scheme := cube.GetColorScheme()

	expected := map[string]string{
		"up":    "white",
		"down":  "yellow",
		"front": "green",
		"back":  "blue",
		"left":  "orange",
		"right": "red",
	}

	if !reflect.DeepEqual(scheme, expected) {
		t.Errorf("Expected color scheme %v, got %v", expected, scheme)
	}
}

func TestRotateFaceClockwise(t *testing.T) {
	face := Face{
		{Red, Green, Blue},
		{Yellow, White, Orange},
		{Green, Blue, Red},
	}

	rotated := rotateFaceClockwise(face)

	expected := Face{
		{Green, Yellow, Red},
		{Blue, White, Green},
		{Red, Orange, Blue},
	}

	if !reflect.DeepEqual(rotated, expected) {
		t.Errorf("Expected rotated face to be %v, got %v", expected, rotated)
	}
}

func TestRotateFaceCounterClockwise(t *testing.T) {
	face := Face{
		{Red, Green, Blue},
		{Yellow, White, Orange},
		{Green, Blue, Red},
	}

	rotated := rotateFaceCounterClockwise(face)

	expected := Face{
		{Blue, Orange, Red},
		{Green, White, Blue},
		{Red, Yellow, Green},
	}

	if !reflect.DeepEqual(rotated, expected) {
		t.Errorf("Expected rotated face to be %v, got %v", expected, rotated)
	}
}

func TestMoveNotation(t *testing.T) {
	cube := New()

	if err := cube.Move("X"); err == nil {
		t.Error("Expected error for invalid move, got nil")
	}

	if err := cube.Move(""); err == nil {
		t.Error("Expected error for empty move, got nil")
	}

	cube.Reset()
	if err := cube.Move("F"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if cube.Up[2][0] != Orange {
		t.Errorf("After F move, Up[2][0] should be Orange, got %s", cube.Up[2][0])
	}

	cube.Reset()
	if err := cube.Move("F'"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if cube.Up[2][0] != Red {
		t.Errorf("After F' move, Up[2][0] should be Red, got %s", cube.Up[2][0])
	}

	cube.Reset()
	if err := cube.Move("F2"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if cube.Front[0][0] != Green {
		t.Errorf("After F2 move, Front[0][0] should still be Green, got %s", cube.Front[0][0])
	}
}

func TestRotateFront(t *testing.T) {
	cube := New()

	cube.RotateFront(true)

	if cube.Front[0][0] != Green {
		t.Errorf("Front face center should remain Green, got %s", cube.Front[0][0])
	}

	if cube.Up[2][0] != Orange || cube.Up[2][1] != Orange || cube.Up[2][2] != Orange {
		t.Errorf("After F, bottom row of Up face should be Orange, got %s, %s, %s",
			cube.Up[2][0], cube.Up[2][1], cube.Up[2][2])
	}

	cube.Reset()
	cube.RotateFront(false)

	if cube.Up[2][0] != Red || cube.Up[2][1] != Red || cube.Up[2][2] != Red {
		t.Errorf("After F', bottom row of Up face should be Red, got %s, %s, %s",
			cube.Up[2][0], cube.Up[2][1], cube.Up[2][2])
	}
}

func TestRotateBack(t *testing.T) {
	cube := New()

	cube.RotateBack(true)

	if cube.Back[0][0] != Blue {
		t.Errorf("Back face center should remain Blue, got %s", cube.Back[0][0])
	}

	if cube.Up[0][0] != Red || cube.Up[0][1] != Red || cube.Up[0][2] != Red {
		t.Errorf("After B, top row of Up face should be Red, got %s, %s, %s",
			cube.Up[0][0], cube.Up[0][1], cube.Up[0][2])
	}
}

func TestRotateUp(t *testing.T) {
	cube := New()

	cube.RotateUp(true)

	if cube.Up[1][1] != White {
		t.Errorf("Up face center should remain White, got %s", cube.Up[1][1])
	}

	if cube.Front[0][0] != Red || cube.Front[0][1] != Red || cube.Front[0][2] != Red {
		t.Errorf("After U, top row of Front face should be Red, got %s, %s, %s",
			cube.Front[0][0], cube.Front[0][1], cube.Front[0][2])
	}
}

func TestRotateDown(t *testing.T) {
	cube := New()

	cube.RotateDown(true)

	if cube.Down[1][1] != Yellow {
		t.Errorf("Down face center should remain Yellow, got %s", cube.Down[1][1])
	}

	if cube.Front[2][0] != Orange || cube.Front[2][1] != Orange || cube.Front[2][2] != Orange {
		t.Errorf("After D, bottom row of Front face should be Orange, got %s, %s, %s",
			cube.Front[2][0], cube.Front[2][1], cube.Front[2][2])
	}
}

func TestRotateLeft(t *testing.T) {
	cube := New()

	cube.RotateLeft(true)

	if cube.Left[1][1] != Orange {
		t.Errorf("Left face center should remain Orange, got %s", cube.Left[1][1])
	}

	if cube.Up[0][0] != Blue || cube.Up[1][0] != Blue || cube.Up[2][0] != Blue {
		t.Errorf("After L, left column of Up face should be Blue, got %s, %s, %s",
			cube.Up[0][0], cube.Up[1][0], cube.Up[2][0])
	}
}

func TestRotateRight(t *testing.T) {
	cube := New()

	cube.RotateRight(true)

	if cube.Right[1][1] != Red {
		t.Errorf("Right face center should remain Red, got %s", cube.Right[1][1])
	}

	if cube.Up[0][2] != Green || cube.Up[1][2] != Green || cube.Up[2][2] != Green {
		t.Errorf("After R, right column of Up face should be Green, got %s, %s, %s",
			cube.Up[0][2], cube.Up[1][2], cube.Up[2][2])
	}
}

func TestInvalidFace(t *testing.T) {
	cube := New()

	err := cube.RotateFace("invalid", true)
	if err == nil {
		t.Error("Expected error when rotating invalid face, got nil")
	}
}

func TestString(t *testing.T) {
	cube := New()

	str := cube.String()

	if len(str) < 10 {
		t.Error("String representation too short")
	}

	for _, face := range []string{"up", "down", "front", "back", "left", "right"} {
		if !strings.Contains(str, face) {
			t.Errorf("String representation missing %s face", face)
		}
	}
}

func TestRotateFaceEquivalent(t *testing.T) {
	cube1 := New()
	cube2 := New()

	cube1.RotateFace("front", true)

	cube2.RotateFront(true)

	if !reflect.DeepEqual(cube1, cube2) {
		t.Error("RotateFace and direct rotation methods give different results")
	}
}

func TestCubeIntegrity(t *testing.T) {
	cube := New()

	moves := []string{"F", "R", "U", "B", "L", "D", "F'", "R'", "U'", "F2", "R2"}
	for _, move := range moves {
		err := cube.Move(move)
		if err != nil {
			t.Errorf("Error executing move %s: %v", move, err)
		}
	}

	colorCounts := make(map[Color]int)

	countColors := func(face Face) {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				colorCounts[face[i][j]]++
			}
		}
	}

	countColors(cube.Up)
	countColors(cube.Down)
	countColors(cube.Front)
	countColors(cube.Back)
	countColors(cube.Left)
	countColors(cube.Right)

	for color, count := range colorCounts {
		if count != 9 {
			t.Errorf("Cube integrity issue: expected 9 of color %s, found %d", color, count)
		}
	}

	if cube.Up[1][1] != White {
		t.Errorf("Center of Up face should always be White, got %s", cube.Up[1][1])
	}
	if cube.Down[1][1] != Yellow {
		t.Errorf("Center of Down face should always be Yellow, got %s", cube.Down[1][1])
	}
	if cube.Front[1][1] != Green {
		t.Errorf("Center of Front face should always be Green, got %s", cube.Front[1][1])
	}
	if cube.Back[1][1] != Blue {
		t.Errorf("Center of Back face should always be Blue, got %s", cube.Back[1][1])
	}
	if cube.Left[1][1] != Orange {
		t.Errorf("Center of Left face should always be Orange, got %s", cube.Left[1][1])
	}
	if cube.Right[1][1] != Red {
		t.Errorf("Center of Right face should always be Red, got %s", cube.Right[1][1])
	}
}
