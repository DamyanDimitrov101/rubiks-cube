package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/DamyanDimitrov101/rubiks-cube-simulator/api"
)

func main() {
	fmt.Println("Starting Rubik's Cube Server...")

	cubeManager := api.NewCubeManager()

	http.HandleFunc("/api/cube", cubeManager.GetCubeHandler)
	http.HandleFunc("/api/cube/rotate", cubeManager.RotateHandler)
	http.HandleFunc("/api/cube/move", cubeManager.MoveHandler)
	http.HandleFunc("/api/cube/reset", cubeManager.ResetHandler)

	scriptsBuildDir := filepath.Join("..", "scripts", "build")

	if _, err := os.Stat(scriptsBuildDir); os.IsNotExist(err) {
		log.Printf("Warning: %s directory not found. Frontend will not be served.", scriptsBuildDir)
	} else {
		fs := http.FileServer(http.Dir(scriptsBuildDir))
		http.Handle("/", fs)
		fmt.Println("Serving frontend from", scriptsBuildDir)
	}

	port := "8080"
	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
