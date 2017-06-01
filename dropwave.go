package main

import "fmt"
import "math"
import "github.com/hugollm/exodus"

func main() {
    search := exodus.Search{
        IndividualSize: 2,
        PopulationSize: 30,
        CrossoverRate: 0.6,
        MutationRate: 0.3,
        NewGene: newGene,
        Fitness: fitness,
        OnGeneration: onGeneration,
    }
    search.Start()
}

func newGene() float64 {
    return exodus.RandomFloat64(-10, 10)
}

func fitness(genome []float64) float64 {
    numerator := 1 + math.Cos(12 * math.Sqrt(math.Pow(genome[0], 2) + math.Pow(genome[1], 2)))
    denominator := 0.5 * (math.Pow(genome[0], 2) + math.Pow(genome[1], 2)) + 2
    return numerator / denominator
}

func onGeneration(search *exodus.Search) {
    if search.Generation % 1000 == 0 {
        fmt.Println(search.Best.Fitness, "|", search.Generation)
    }
    if search.Best.Fitness >= (1 - 0.0000001) {
        fmt.Println()
        fmt.Println("Generation: ", search.Generation)
        fmt.Println("Genome:     ", search.Best.Genome)
        fmt.Println("Fitness:    ", search.Best.Fitness)
        fmt.Println()
        search.Stop()
    }
}
