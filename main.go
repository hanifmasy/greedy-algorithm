package main

import (
	"fmt"
	"math"
)

type Food struct {
	Name        string
	Carbohydrate int // in grams
	CostPer100g float64
}

func main() {
	// Define the available foods
	foods := []Food{
		{"Rice", 28, 0.1},   // CostPer100g in dollars
		{"Corn", 21, 0.08},
		{"Potato", 17, 0.05},
	}

	// Calculate the cost for each food per gram of carbohydrate
	for i := range foods {
		foods[i].CostPerGramCarb = foods[i].CostPer100g / float64(foods[i].Carbohydrate)
	}

	// Calculate the minimum cost combination
	minCost := math.MaxFloat64
	var minCombination []int

	for rice := 0; rice <= 400; rice += 100 {
		for corn := 0; corn <= 400; corn += 100 {
			for potato := 0; potato <= 400; potato += 100 {
				if rice+corn+potato == 400 {
					cost := float64(rice)/100*foods[0].CostPer100g + float64(corn)/100*foods[1].CostPer100g + float64(potato)/100*foods[2].CostPer100g
					if cost < minCost {
						minCost = cost
						minCombination = []int{rice, corn, potato}
					}
				}
			}
		}
	}

	// Output the result
	fmt.Println("Minimum cost combination:")
	for i, food := range foods {
		fmt.Printf("%s: %d grams\n", food.Name, minCombination[i])
	}
	fmt.Printf("Total cost: $%.2f\n", minCost)
}
