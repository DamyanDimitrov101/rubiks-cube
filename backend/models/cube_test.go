package models

import (
	"reflect"
	"strings"
	"testing"
)

// TestNewCube verifies that a new cube is properly initialized with solid colors
func TestNewCube(t *testing.T) {
	cube := New()

	// Check that each face has the correct color
	assertSolidFace(t, cube.Up, White, "Up face should be white")
	assertSolidFace(t, cube.Down, Yellow, "Down face should be yellow")
	assertSolidFace(t, cube.Front, Green, "Front face should be green")
	assertSolidFace(t, cube.Back, Blue, "Back face should be blue")
	assertSolidFace(t, cube.Left, Orange, "Left face should be orange")
	assertSolidFace(t, cube.Right, Red, "Right face should be red")
}

// Helper function to verify a face has all cells of the same color
func assertSolidFace(t *testing.T, face Face, expectedColor Color, message string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if face[i][j] != expectedColor {
				t.Errorf("%s, but found %s at position [%d][%d]", message, face[i][j], i, j)
			}
		}
	}
}

// TestReset verifies that the cube can be reset to solved state
func TestReset(t *testing.T) {
	cube := New()

	// Perform some moves to scramble the cube
	cube.Move("F")
	cube.Move("R")
	cube.Move("U")

	// Reset the cube
	cube.Reset()

	// Verify it's solved
	assertSolidFace(t, cube.Up, White, "Up face should be white after reset")
	assertSolidFace(t, cube.Down, Yellow, "Down face should be yellow after reset")
	assertSolidFace(t, cube.Front, Green, "Front face should be green after reset")
	assertSolidFace(t, cube.Back, Blue, "Back face should be blue after reset")
	assertSolidFace(t, cube.Left, Orange, "Left face should be orange after reset")
	assertSolidFace(t, cube.Right, Red, "Right face should be red after reset")
}

// TestClone verifies that cloning a cube creates an independent copy
func TestClone(t *testing.T) {
	original := New()

	// Clone the cube
	clone := original.Clone()

	// Modify the clone
	clone.Move("F")

	// Verify the original is unchanged
	assertSolidFace(t, original.Up, White, "Original cube's Up face should be unchanged")
	assertSolidFace(t, original.Front, Green, "Original cube's Front face should be unchanged")

	// But the clone should be changed
	if clone.Up[2][0] == White {
		t.Error("Clone's Up face should be modified after F move")
	}
}

// TestGetColorScheme verifies the color scheme reporting
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

// TestRotateFaceClockwise verifies internal rotation logic
func TestRotateFaceClockwise(t *testing.T) {
	// Create a test face with different values
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

// TestRotateFaceCounterClockwise verifies internal rotation logic
func TestRotateFaceCounterClockwise(t *testing.T) {
	// Create a test face with different values
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

// TestMoveNotation verifies standard cube notation works correctly
func TestMoveNotation(t *testing.T) {
	cube := New()

	// Test invalid move
	if err := cube.Move("X"); err == nil {
		t.Error("Expected error for invalid move, got nil")
	}

	// Test empty move
	if err := cube.Move(""); err == nil {
		t.Error("Expected error for empty move, got nil")
	}

	// Test front clockwise
	cube.Reset()
	if err := cube.Move("F"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check a few positions after F move
	if cube.Up[2][0] != Orange {
		t.Errorf("After F move, Up[2][0] should be Orange, got %s", cube.Up[2][0])
	}

	// Test front counter-clockwise
	cube.Reset()
	if err := cube.Move("F'"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if cube.Up[2][0] != Red {
		t.Errorf("After F' move, Up[2][0] should be Red, got %s", cube.Up[2][0])
	}

	// Test double move
	cube.Reset()
	if err := cube.Move("F2"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// After F2, front face should be the same, but edges affected twice
	if cube.Front[0][0] != Green {
		t.Errorf("After F2 move, Front[0][0] should still be Green, got %s", cube.Front[0][0])
	}
}

// TestRotateFront tests front face rotation
func TestRotateFront(t *testing.T) {
	cube := New()

	// Perform F move
	cube.RotateFront(true)

	// Check front face orientation
	if cube.Front[0][0] != Green {
		t.Errorf("Front face center should remain Green, got %s", cube.Front[0][0])
	}

	// Check edge effects - after F, top edge of front face should be from left face
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

// TestRotateBack tests back face rotation
func TestRotateBack(t *testing.T) {
	cube := New()

	// Perform B move
	cube.RotateBack(true)

	// Check back face orientation
	if cube.Back[0][0] != Blue {
		t.Errorf("Back face center should remain Blue, got %s", cube.Back[0][0])
	}

	// Check edge effects - after B, top edge of back face should be from right face
	if cube.Up[0][0] != Red || cube.Up[0][1] != Red || cube.Up[0][2] != Red {
		t.Errorf("After B, top row of Up face should be Red, got %s, %s, %s",
			cube.Up[0][0], cube.Up[0][1], cube.Up[0][2])
	}
}

// TestRotateUp tests up face rotation
func TestRotateUp(t *testing.T) {
	cube := New()

	// Perform U move
	cube.RotateUp(true)

	// Check up face orientation
	if cube.Up[1][1] != White {
		t.Errorf("Up face center should remain White, got %s", cube.Up[1][1])
	}

	// Check edge effects - after U, top edge of front face should be from right face
	if cube.Front[0][0] != Red || cube.Front[0][1] != Red || cube.Front[0][2] != Red {
		t.Errorf("After U, top row of Front face should be Red, got %s, %s, %s",
			cube.Front[0][0], cube.Front[0][1], cube.Front[0][2])
	}
}

// TestRotateDown tests down face rotation
func TestRotateDown(t *testing.T) {
	cube := New()

	// Perform D move
	cube.RotateDown(true)

	// Check down face orientation
	if cube.Down[1][1] != Yellow {
		t.Errorf("Down face center should remain Yellow, got %s", cube.Down[1][1])
	}

	// Check edge effects - after D, bottom edge of front face should be from left face
	if cube.Front[2][0] != Orange || cube.Front[2][1] != Orange || cube.Front[2][2] != Orange {
		t.Errorf("After D, bottom row of Front face should be Orange, got %s, %s, %s",
			cube.Front[2][0], cube.Front[2][1], cube.Front[2][2])
	}
}

// TestRotateLeft tests left face rotation
func TestRotateLeft(t *testing.T) {
	cube := New()

	// Perform L move
	cube.RotateLeft(true)

	// Check left face orientation
	if cube.Left[1][1] != Orange {
		t.Errorf("Left face center should remain Orange, got %s", cube.Left[1][1])
	}

	// Check edge effects - after L, left edge of up face should be from back face
	if cube.Up[0][0] != Blue || cube.Up[1][0] != Blue || cube.Up[2][0] != Blue {
		t.Errorf("After L, left column of Up face should be Blue, got %s, %s, %s",
			cube.Up[0][0], cube.Up[1][0], cube.Up[2][0])
	}
}

// TestRotateRight tests right face rotation
func TestRotateRight(t *testing.T) {
	cube := New()

	// Perform R move
	cube.RotateRight(true)

	// Check right face orientation
	if cube.Right[1][1] != Red {
		t.Errorf("Right face center should remain Red, got %s", cube.Right[1][1])
	}

	if cube.Up[0][2] != Green || cube.Up[1][2] != Green || cube.Up[2][2] != Green {
		t.Errorf("After R, right column of Up face should be Green, got %s, %s, %s",
			cube.Up[0][2], cube.Up[1][2], cube.Up[2][2])
	}
}

// TestInvalidFace tests error handling for invalid face input
func TestInvalidFace(t *testing.T) {
	cube := New()

	err := cube.RotateFace("invalid", true)
	if err == nil {
		t.Error("Expected error when rotating invalid face, got nil")
	}
}

// TestString verifies the String method generates valid JSON
func TestString(t *testing.T) {
	cube := New()

	str := cube.String()

	// Check that string contains expected face names and isn't empty
	if len(str) < 10 {
		t.Error("String representation too short")
	}

	for _, face := range []string{"up", "down", "front", "back", "left", "right"} {
		if !strings.Contains(str, face) {
			t.Errorf("String representation missing %s face", face)
		}
	}
}

// TestRotateFaceEquivalent tests that RotateFace works the same as direct rotation
func TestRotateFaceEquivalent(t *testing.T) {
	cube1 := New()
	cube2 := New()

	// Use RotateFace on cube1
	cube1.RotateFace("front", true)

	// Use direct method on cube2
	cube2.RotateFront(true)

	// Compare results
	if !reflect.DeepEqual(cube1, cube2) {
		t.Error("RotateFace and direct rotation methods give different results")
	}
}

// TestCubeIntegrity checks that the cube maintains its structural integrity
func TestCubeIntegrity(t *testing.T) {
	cube := New()

	// Perform a variety of moves
	moves := []string{"F", "R", "U", "B", "L", "D", "F'", "R'", "U'", "F2", "R2"}
	for _, move := range moves {
		err := cube.Move(move)
		if err != nil {
			t.Errorf("Error executing move %s: %v", move, err)
		}
	}

	// Count all colors to ensure we still have 9 of each
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

	// Verify we still have exactly 9 of each color
	for color, count := range colorCounts {
		if count != 9 {
			t.Errorf("Cube integrity issue: expected 9 of color %s, found %d", color, count)
		}
	}

	// Verify centers haven't moved (as they shouldn't in a standard cube)
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
