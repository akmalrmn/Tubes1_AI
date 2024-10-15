package server

import (
    "net/http"
    "magic-cube-solver/internal/handlers"
)

func StartServer() {
    http.HandleFunc("/api/genetic-algorithm", handlers.GeneticAlgorithmHandler)
    http.ListenAndServe(":8080", nil)
}