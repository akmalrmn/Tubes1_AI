package server

import (
    "net/http"
    "magic-cube-solver/internal/handlers"
)

func StartServer() {
    http.HandleFunc("/api/simulated-annealing", handlers.SimulatedAnnealingHandler)
    http.HandleFunc("/api/steepest-ascent", handlers.SteepestAscentHandler)
    http.HandleFunc("/api/genetic-algorithm", handlers.GeneticAlgorithmHandler)
    http.ListenAndServe(":8070", nil)
}