package mcbanner

// Genetic algorithm

import (
	"fmt"
	"math/rand"
)

const (
	MAXGENERATIONS = 300
	POPSIZE        = 50
)

// This is what every solution is being compared against
var target_png_bytes []byte

// A solution is up to maxPatterns Patterns
type Solution []Pattern

// A population is a collection of solutions
type Population []Solution

// For storing the population fitnesses before sorting
type PopulationFitness []float64

func NewPopulationFitness(popsize int) PopulationFitness {
	return make(PopulationFitness, popsize)
}

func NewSolution() Solution {
	return make([]Pattern, maxPatterns)
}

func NewRandomSolution() Solution {
	sol := NewSolution()
	for i := 0; i < len(sol); i++ {
		sol[i] = NewRandomPattern()
	}
	return sol
}

func (sol Solution) String() string {
	return fmt.Sprintf("%v", sol)
}

func (sol Solution) Banner() *Banner {
	b := NewBanner()
	for i := 0; i < len(sol); i++ {
		b.AddPattern(&sol[i])
	}
	return b
}

func (sol Solution) fitness() float64 {
	return Compare(sol.Banner(), target_png_bytes)
}

func NewPopulation(size int) Population {
	Seed()
	pop := make([]Solution, size)
	for i := 0; i < size; i++ {
		pop[i] = NewRandomSolution()
	}
	return pop
}

func test_solution() {
	sol := NewSolution()
	fmt.Println(sol)
	fmt.Println("fitness:", sol.fitness())
}

func crossover(a, b Solution, point uint) Solution {
	c := NewSolution()
	var i uint
	for i = 0; i < point; i++ {
		c[i] = a[i]
	}
	for i = point; i < maxPatterns; i++ {
		c[i] = b[i]
	}
	return c
}

func (sol Solution) mutate() {
	randpos := rand.Intn(maxPatterns)
	sol[randpos] = NewRandomPattern()
}

func sum(scores []float64) float64 {
	var total float64
	for _, score := range scores {
		total += score
	}
	return total
}

func FindBest(fitnessfunction func([]byte, []byte) float64, png_bytes []byte) {
	target_png_bytes = png_bytes
	var (
		bestfitnessindex     = 0
		popsize          int = POPSIZE
		pop                  = NewPopulation(popsize)
		generation       int
		average          float64
	)
	for generation = 0; generation < MAXGENERATIONS; generation++ {
		fmt.Println("Generation", generation)
		fit := NewPopulationFitness(popsize)
		for i, s := range pop {
			fit[i] = s.fitness()
		}
		//fmt.Println(fit)
		total := sum(fit)
		fmt.Println("total =", total)
		average = total / float64(popsize)
		fmt.Println("average =", average)
		var (
			bestfitness, nextbestfitness float64
		)
		nextbestfitnessindex := 0
		for i, _ := range fit {
			if fit[i] >= bestfitness {
				nextbestfitness = bestfitness
				bestfitness = fit[i]
				nextbestfitnessindex = bestfitnessindex
				bestfitnessindex = i
			}
		}
		fmt.Println("best =", bestfitness)
		fmt.Println("nextbest =", nextbestfitness)
		if bestfitness == 1.0 {
			fmt.Println("Found fitness 1")
			break
		}
		var (
			mutrate    float64 = 0.0
			crossrate  float64 = 0.1
			newpoprate float64 = 0.0
		)
		for i, _ := range pop {
			fitness := fit[i]
			if average > 0.7 && fitness < 0.5 {
				pop[i] = NewSolution()
			} else if average > 0.8 && fitness < 0.6 {
				pop[i] = NewSolution()
			} else if average > 0.9 && fitness < 0.7 {
				pop[i] = NewSolution()
			} else if fitness < (average * 0.3) {
				// 50% chance of being replaced with randomness
				if rand.Float64() <= 0.5 {
					pop[i] = NewSolution()
				}
			}
			if bestfitness > average {
				// slow down the mutation rate
				mutrate = 0.15
				crossrate = 0.07
			} else {
				mutrate = 0.4
				crossrate = 0.4
			}
			if bestfitness == nextbestfitness {
				mutrate *= 3.0
			}
			if average > 0.9 {
				newpoprate = 0.4
			} else {
				newpoprate = 0.2
			}
			// An advantage for the best ones
			if fitness > (bestfitness * 0.9) {
				if rand.Float64() <= 0.9 {
					continue
				}
			}
			// A certain chance for mutation
			if rand.Float64() <= mutrate {
				pop[rand.Intn(int(popsize))].mutate()
			}
			// A certain chance for crossover
			if rand.Float64() <= crossrate {
				crossoverpoint := uint(rand.Intn(maxPatterns))
				pop[i] = crossover(pop[bestfitnessindex], pop[nextbestfitnessindex], crossoverpoint)
			}
			// A certain chance for new random variations
			if rand.Float64() <= newpoprate {
				pop[i] = NewSolution()
			}
		}
	}
	fmt.Println("generation", generation)
	fmt.Println(pop[bestfitnessindex])
}
