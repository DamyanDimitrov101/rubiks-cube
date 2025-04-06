package validators

import (
	"fmt"
	"regexp"
)

// ValidateFace checks if the face is valid
func ValidateFace(face string) error {
	if face == "" {
		return fmt.Errorf("face cannot be empty")
	}

	validFaces := map[string]bool{
		"front": true,
		"back":  true,
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	if !validFaces[face] {
		return fmt.Errorf("invalid face: %s. Valid faces are: front, back, up, down, left, right", face)
	}

	return nil
}

// ValidateNotation checks if the move notation is valid
func ValidateNotation(notation string) error {
	if notation == "" {
		return fmt.Errorf("notation cannot be empty")
	}

	// Regular expression to match valid notations
	// F, B, U, D, L, R, F', B', U', D', L', R', F2, B2, U2, D2, L2, R2
	validPattern := regexp.MustCompile(`^[FBUDLR]('|2)?$`)

	if !validPattern.MatchString(notation) {
		return fmt.Errorf("invalid notation: %s. Valid examples: F, B, U, D, L, R, F', B', U', D', L', R', F2, B2, U2, D2, L2, R2", notation)
	}

	return nil
}
