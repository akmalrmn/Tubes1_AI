package server

import (
    "net/http"
    "magic-cube-solver/internal/handlers"
)

func StartServer() {
    http.HandleFunc("/api/genetic-algorithm", handlers.GeneticAlgorithmHandler)
    http.HandleFunc("/api/simulated-annealing", handlers.SimulatedAnnealingHandler)
    http.ListenAndServe(":8070", nil)
}