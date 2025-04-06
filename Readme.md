# Rubik's Cube Simulator API

A Go-based REST API that simulates a Rubik's Cube, allowing users to perform standard cube operations, track the cube state, and reset the cube.

## Features

- Get the current state of the Rubik's Cube
- Rotate cube faces (clockwise or counter-clockwise)
- Execute standard notation moves (e.g., F, R', U2, etc.)
- Reset the cube to its solved state
- Thread-safe operations
- Validation for all inputs

## Installation

### Prerequisites

- Go 1.13 or higher
- Git

### Setup

1. Clone the repository:

```bash
git clone https://github.com/DamyanDimitrov101/rubiks-cube-simulator.git
```

2. Navigate to the backend directory:

```bash
cd rubiks-cube-simulator/backend
```

3. Run the server:

**Run rubiks-cube-simulator.exe**

OR

open terminal and execute the command:
```bash
go run main.go
```

By default, the server will start on port 8080. You can modify this in the main.go file.

## API Reference

### Get Cube State

Returns the current state of the Rubik's Cube.

- **URL**: `/cube`
- **Method**: `GET`
- **Response Example**:
```json
{
  "up": [
    ["white", "white", "white"],
    ["white", "white", "white"],
    ["white", "white", "white"]
  ],
  ...
}
```

### Rotate Face

Rotates a specific face of the cube clockwise or counter-clockwise.

- **URL**: `/rotate`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "face": "front", 
    "clockwise": true
  }
  ```
    - `face`: One of "front", "back", "up", "down", "left", "right"
    - `clockwise`: Boolean indicating direction (true for clockwise, false for counter-clockwise)
- **Response Example**:
```json
{
  "success": true,
  "cube": {
    "up": [
      ["white", "white", "white"],
      ["white", "white", "white"],
      ["orange", "orange", "orange"]
    ],
    ...
    }
}
```

### Execute Move

Performs a move using standard Rubik's Cube notation.

- **URL**: `/move`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "notation": "F'"
  }
  ```
    - `notation`: Standard notation string (e.g., "F", "R", "U2")
    - Valid notations: F, B, U, D, L, R, F', B', U', D', L', R', F2, B2, U2, D2, L2, R2
- **Response Example**: 
```json
{
  "success": true,
  "cube": {
    "up": [
      ["white", "white", "white"],
      ["white", "white", "white"],
      ["white", "white", "yellow"]
    ],
    ...
  }
}
```

### Reset Cube

Resets the cube to its solved state.

- **URL**: `/reset`
- **Method**: `POST`
- **Response**: Confirmation message and the reset cube state

## Error Handling

The API provides structured error responses for validation issues:

```json
{
  "success": false,
  "errors": [
    {
      "field": "notation",
      "message": "Invalid notation: X. Valid examples: F, B, U, D, L, R, F', B', U', D', L', R', F2, B2, U2, D2, L2, R2"
    }
  ]
}
```

## Input Validation

### Face Validation
- Must be one of: "front", "back", "up", "down", "left", "right"
- Cannot be empty

### Notation Validation
- Must match the standard Rubik's Cube notation pattern
- Valid examples: F, B, U, D, L, R, F', B', U', D', L', R', F2, B2, U2, D2, L2, R2
- Cannot be empty

## Project Structure

- `api/` - HTTP handlers and routing
- `models/` - Core cube model and operations
- `validators/` - Input validation logic
- `main.go` - Application entry point

## CORS Support

The API supports Cross-Origin Resource Sharing (CORS) with all endpoints allowing requests from any origin.
****