package main

import (
    "log"
    "magic-cube-solver/internal/server"
)

func main() {
    log.Println("Starting server on http://localhost:8080")
    server.StartServer()
}