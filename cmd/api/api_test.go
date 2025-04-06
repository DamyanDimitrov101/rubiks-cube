package api

import (
	"bytes"
	"encoding/json"
	"github.com/DamyanDimitrov101/rubiks-cube-simulator/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TestGetCubeHandler tests the GetCubeHandler function
func TestGetCubeHandler(t *testing.T) {
	// Create a new cube manager
	cm := NewCubeManager()

	// Test GET request
	req, err := http.NewRequest("GET", "/cube", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(cm.GetCubeHandler)

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	var response models.RubiksCube
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	// Verify the cube is solved initially
	solvedCube := models.New()
	if !reflect.DeepEqual(response, *solvedCube) {
		t.Errorf("Expected a solved cube, got %v", response)
	}

	// Test non-GET request
	req, err = http.NewRequest("POST", "/cube", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check status code (should be Method Not Allowed)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

// TestRotateHandlerValidation tests validation in the RotateHandler
func TestRotateHandlerValidation(t *testing.T) {
	cm := NewCubeManager()

	// Test cases for validation
	testCases := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
		expectedErrors []ValidationError
	}{
		{
			name: "Valid Request",
			requestBody: map[string]interface{}{
				"face":      "front",
				"clockwise": true,
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Invalid Face",
			requestBody: map[string]interface{}{
				"face":      "invalid",
				"clockwise": true,
			},
			expectedStatus: http.StatusBadRequest,
			expectedErrors: []ValidationError{
				{
					Field:   "face",
					Message: "invalid face: invalid. Valid faces are: front, back, up, down, left, right",
				},
			},
		},
		{
			name: "Empty Face",
			requestBody: map[string]interface{}{
				"face":      "",
				"clockwise": true,
			},
			expectedStatus: http.StatusBadRequest,
			expectedErrors: []ValidationError{
				{
					Field:   "face",
					Message: "face cannot be empty",
				},
			},
		},
		{
			name: "Case Sensitive Face",
			requestBody: map[string]interface{}{
				"face":      "Front", // Should be lowercase
				"clockwise": true,
			},
			expectedStatus: http.StatusBadRequest,
			expectedErrors: []ValidationError{
				{
					Field:   "face",
					Message: "invalid face: Front. Valid faces are: front, back, up, down, left, right",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert test case request body to JSON
			body, _ := json.Marshal(tc.requestBody)

			// Create a new request
			req, _ := http.NewRequest("POST", "/rotate", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Serve the request
			handler := http.HandlerFunc(cm.RotateHandler)
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.expectedStatus)
			}

			// For error cases, verify the error response
			if tc.expectedStatus == http.StatusBadRequest && tc.expectedErrors != nil {
				var response ValidationResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Failed to unmarshal error response: %v", err)
				}

				// Verify success flag is false
				if response.Success {
					t.Errorf("Expected success to be false for error response")
				}

				// Verify errors match expected
				if len(response.Errors) != len(tc.expectedErrors) {
					t.Errorf("Expected %d errors, got %d", len(tc.expectedErrors), len(response.Errors))
				} else {
					for i, expectedErr := range tc.expectedErrors {
						actualErr := response.Errors[i]
						if expectedErr.Field != actualErr.Field || expectedErr.Message != actualErr.Message {
							t.Errorf("Error mismatch:\nExpected: %+v\nActual: %+v", expectedErr, actualErr)
						}
					}
				}
			}

			// For success case, verify success response
			if tc.expectedStatus == http.StatusOK {
				var response map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Failed to unmarshal success response: %v", err)
				}

				// Verify success flag is true
				success, ok := response["success"]
				if !ok || success != true {
					t.Errorf("Expected success to be true for success response")
				}

				// Verify cube is present in response
				_, ok = response["cube"]
				if !ok {
					t.Errorf("Expected cube in success response")
				}
			}
		})
	}
}

// TestMoveHandlerValidation tests validation in the MoveHandler
func TestMoveHandlerValidation(t *testing.T) {
	cm := NewCubeManager()

	// Test cases for validation
	testCases := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
		expectedErrors []ValidationError
	}{
		{
			name: "Valid Move F",
			requestBody: map[string]interface{}{
				"notation": "F",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Valid Move R'",
			requestBody: map[string]interface{}{
				"notation": "R'",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Valid Move U2",
			requestBody: map[string]interface{}{
				"notation": "U2",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Invalid Notation",
			requestBody: map[string]interface{}{
				"notation": "X", // Invalid notation
			},
			expectedStatus: http.StatusBadRequest,
			expectedErrors: []ValidationError{
				{
					Field:   "notation",
					Message: "invalid notation: X. Valid examples: F, B, U, D, L, R, F', B', U', D', L', R', F2, B2, U2, D2, L2, R2",
				},
			},
		},
		{
			name: "Empty Notation",
			requestBody: map[string]interface{}{
				"notation": "",
			},
			expectedStatus: http.StatusBadRequest,
			expectedErrors: []ValidationError{
				{
					Field:   "notation",
					Message: "notation cannot be empty",
				},
			},
		},
		{
			name: "Lowercase Notation",
			requestBody: map[string]interface{}{
				"notation": "f", // Should be uppercase
			},
			expectedStatus: http.StatusBadRequest,
			expectedErrors: []ValidationError{
				{
					Field:   "notation",
					Message: "invalid notation: f. Valid examples: F, B, U, D, L, R, F', B', U', D', L', R', F2, B2, U2, D2, L2, R2",
				},
			},
		},
		{
			name: "Multiple Moves",
			requestBody: map[string]interface{}{
				"notation": "FF", // Should be single move
			},
			expectedStatus: http.StatusBadRequest,
			expectedErrors: []ValidationError{
				{
					Field:   "notation",
					Message: "invalid notation: FF. Valid examples: F, B, U, D, L, R, F', B', U', D', L', R', F2, B2, U2, D2, L2, R2",
				},
			},
		},
		{
			name: "Invalid Modifier",
			requestBody: map[string]interface{}{
				"notation": "F3", // Only ' and 2 are valid modifiers
			},
			expectedStatus: http.StatusBadRequest,
			expectedErrors: []ValidationError{
				{
					Field:   "notation",
					Message: "invalid notation: F3. Valid examples: F, B, U, D, L, R, F', B', U', D', L', R', F2, B2, U2, D2, L2, R2",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Convert test case request body to JSON
			body, _ := json.Marshal(tc.requestBody)

			// Create a new request
			req, _ := http.NewRequest("POST", "/move", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Serve the request
			handler := http.HandlerFunc(cm.MoveHandler)
			handler.ServeHTTP(rr, req)

			// Check the status code
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.expectedStatus)
			}

			// For error cases, verify the error response
			if tc.expectedStatus == http.StatusBadRequest && tc.expectedErrors != nil {
				var response ValidationResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Failed to unmarshal error response: %v", err)
				}

				// Verify success flag is false
				if response.Success {
					t.Errorf("Expected success to be false for error response")
				}

				// Verify errors match expected
				if len(response.Errors) != len(tc.expectedErrors) {
					t.Errorf("Expected %d errors, got %d", len(tc.expectedErrors), len(response.Errors))
				} else {
					for i, expectedErr := range tc.expectedErrors {
						actualErr := response.Errors[i]
						if expectedErr.Field != actualErr.Field || expectedErr.Message != actualErr.Message {
							t.Errorf("Error mismatch:\nExpected: %+v\nActual: %+v", expectedErr, actualErr)
						}
					}
				}
			}

			// For success case, verify success response
			if tc.expectedStatus == http.StatusOK {
				var response map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Failed to unmarshal success response: %v", err)
				}

				// Verify success flag is true
				success, ok := response["success"]
				if !ok || success != true {
					t.Errorf("Expected success to be true for success response")
				}

				// Verify cube is present in response
				_, ok = response["cube"]
				if !ok {
					t.Errorf("Expected cube in success response")
				}
			}
		})
	}
}

func TestResetHandler(t *testing.T) {
	// Create a new cube manager
	cm := NewCubeManager()

	// First make some moves to scramble the cube
	cm.cube.Move("F")
	cm.cube.Move("R")
	cm.cube.Move("U")

	// Verify cube is not in solved state
	solvedCube := models.New()
	if reflect.DeepEqual(*cm.cube, *solvedCube) {
		t.Errorf("Cube should be scrambled at this point")
	}

	// Test reset
	req, err := http.NewRequest("POST", "/reset", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(cm.ResetHandler)
	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Verify response contains success flag
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	success, ok := response["success"]
	if !ok || success != true {
		t.Errorf("Expected success flag in response, got %v", response)
	}

	message, ok := response["message"]
	if !ok || message != "Cube has been reset to solved state" {
		t.Errorf("Expected reset message in response, got %v", response)
	}

	// Verify cube is now solved
	if !reflect.DeepEqual(*cm.cube, *solvedCube) {
		t.Errorf("Cube should be solved after reset")
	}

	// Test invalid method
	req, err = http.NewRequest("GET", "/reset", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check status code (should be Method Not Allowed)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}

	// Test OPTIONS method for CORS
	req, err = http.NewRequest("OPTIONS", "/reset", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check status code (should be OK for OPTIONS)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
