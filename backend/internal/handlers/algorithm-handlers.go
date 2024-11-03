package handlers

import (
    "encoding/json"
    "fmt"
    "magic-cube-solver/internal/models"
    "magic-cube-solver/internal/algorithms/simulated_annealing"
    "magic-cube-solver/internal/algorithms/steepest_ascent"
    "net/http"
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
)

func plotHistory(data []float64, filename, xlabel, ylabel string) (string, error) {
    p := plot.New()

    p.Title.Text = fmt.Sprintf("%s over Iterations", ylabel)
    p.X.Label.Text = xlabel
    p.Y.Label.Text = ylabel

    pts := make(plotter.XYs, len(data))
    for i := range pts {
        pts[i].X = float64(i) // X-axis represents the iteration number
        pts[i].Y = data[i]    // Y-axis represents the values
    }

    line, err := plotter.NewLine(pts)
    if err != nil {
        return "", err // Return an error instead of panicking
    }
    p.Add(line)

    if err := p.Save(4*vg.Inch, 4*vg.Inch, filename); err != nil {
        return "", err // Return an error instead of panicking
    }

    return filename, nil // Return the filename and nil for no error
}

func SimulatedAnnealingHandler(w http.ResponseWriter, r *http.Request) {
    initialCube := models.GenerateCube()

    // Run the Simulated Annealing algorithm
    initialState, finalState, energyHistory, acceptanceProbHistory, duration, stuckCount, initialEnergy, totalIterations := simulated_annealing.SimulatedAnnealing(initialCube)

    // Generate plots
    energyPlot, err := plotHistory(energyHistory, "energy_history.png", "Iterations", "Current State Energy")
    if err != nil {
        fmt.Println("Error generating energy plot:", err)
    }

    acceptanceProbPlot, err := plotHistory(acceptanceProbHistory, "acceptance_prob_history.png", "Iterations", "e^(deltaE/T)")
    if err != nil {
        fmt.Println("Error generating acceptance probability plot:", err)
    }

    response := struct {
        InitialState            simulated_annealing.State `json:"initialState"`
        FinalState              simulated_annealing.State `json:"finalState"`
        EnergyHistory           []float64                 `json:"energyHistory"`
        TempHistory             []float64                 `json:"tempHistory"`
        AcceptanceProbHistory   []float64                 `json:"acceptanceProbHistory"`
        Duration                string            `json:"duration"`
        StuckCount              int                       `json:"stuckCount"`
        InitialEnergy           float64                   `json:"initialEnergy"`
        TotalIterations         int                       `json:"totalIterations"`
        EnergyPlot              string                    `json:"energyPlot"`
        AcceptanceProbPlot      string                    `json:"acceptanceProbPlot"`
        FinalObjectiveVal       float64                   `json:"finalObjectiveVal"`
    }{
        InitialState:            initialState,
        FinalState:              finalState,
        EnergyHistory:           energyHistory,
        AcceptanceProbHistory:   acceptanceProbHistory,
        Duration:                duration,
        StuckCount:              stuckCount,
        InitialEnergy:           initialEnergy,
        TotalIterations:         totalIterations,
        EnergyPlot:              energyPlot,
        AcceptanceProbPlot:      acceptanceProbPlot,
        FinalObjectiveVal:       finalState.Energy,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func SteepestAscentHandler(w http.ResponseWriter, r *http.Request) {
    initialState, finalState, objectiveHistory, duration, totalIterations := steepest_ascent.SteepestAscent(10000)

    // Generate plot
    objectivePlot, err := plotHistory(objectiveHistory, "objective_history.png", "Iterations", "Objective Value")
    if err != nil {
        fmt.Println("Error generating objective plot:", err)
    }

    response := struct {
        InitialState      steepest_ascent.State `json:"initialState"`
        FinalState        steepest_ascent.State `json:"finalState"`
        ObjectiveHistory  []float64             `json:"objectiveHistory"`
        Duration          string                `json:"duration"`
        TotalIterations   int                   `json:"totalIterations"`
        ObjectivePlot     string                `json:"objectivePlot"`
        FinalObjectiveVal float64               `json:"finalObjectiveVal"`
    }{
        InitialState:      initialState,
        FinalState:        finalState,
        ObjectiveHistory:  objectiveHistory,
        Duration:          fmt.Sprintf("%.3f", duration.Seconds()),
        TotalIterations:   totalIterations,
        ObjectivePlot:     objectivePlot,
        FinalObjectiveVal: finalState.ObjectiveValue,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}