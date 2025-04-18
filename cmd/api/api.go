package api

import (
	"encoding/json"
	"github.com/DamyanDimitrov101/rubiks-cube-simulator/models"
	"github.com/DamyanDimitrov101/rubiks-cube-simulator/validators"
	"net/http"
	"sync"
)

type CubeManager struct {
	cube  *models.RubiksCube
	mutex sync.RWMutex
}

func NewCubeManager() *CubeManager {
	return &CubeManager{
		cube: models.New(),
	}
}

func (cm *CubeManager) GetCubeHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cm.cube)
}

type rotateRequest struct {
	Face      string `json:"face"`
	Clockwise bool   `json:"clockwise"`
}

func (cm *CubeManager) RotateHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req rotateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var validationErrors []ValidationError
	if err := validators.ValidateFace(req.Face); err != nil {
		validationErrors = append(validationErrors, ValidationError{
			Field:   "face",
			Message: err.Error(),
		})
	}

	if len(validationErrors) > 0 {
		respondWithValidationError(w, validationErrors)
		return
	}

	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	err = cm.cube.RotateFace(req.Face, req.Clockwise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"cube":    cm.cube,
	})
}

type moveRequest struct {
	Notation string `json:"notation"`
}

func (cm *CubeManager) MoveHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req moveRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var validationErrors []ValidationError
	if err := validators.ValidateNotation(req.Notation); err != nil {
		validationErrors = append(validationErrors, ValidationError{
			Field:   "notation",
			Message: err.Error(),
		})
	}

	if len(validationErrors) > 0 {
		respondWithValidationError(w, validationErrors)
		return
	}

	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	err = cm.cube.Move(req.Notation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"cube":    cm.cube,
	})
}

func (cm *CubeManager) ResetHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.cube.Reset()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Cube has been reset to solved state",
		"cube":    cm.cube,
	})
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResponse struct {
	Success bool              `json:"success"`
	Errors  []ValidationError `json:"errors"`
}

func respondWithValidationError(w http.ResponseWriter, errors []ValidationError) {
	response := ValidationResponse{
		Success: false,
		Errors:  errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
