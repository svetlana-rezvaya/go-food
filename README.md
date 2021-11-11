# go-food

[![GoDoc](https://godoc.org/github.com/svetlana-rezvaya/go-food?status.svg)](https://godoc.org/github.com/svetlana-rezvaya/go-food)
[![Go Report Card](https://goreportcard.com/badge/github.com/svetlana-rezvaya/go-food)](https://goreportcard.com/report/github.com/svetlana-rezvaya/go-food)
[![Build Status](https://app.travis-ci.com/svetlana-rezvaya/go-food.svg?branch=master)](https://app.travis-ci.com/svetlana-rezvaya/go-food)
[![codecov](https://codecov.io/gh/svetlana-rezvaya/go-food/branch/master/graph/badge.svg)](https://codecov.io/gh/svetlana-rezvaya/go-food)

The library for calculating calories and macronutrients (protein, carbohydrates, and fat) of specified dishes.

## Problems

- [Problem 1](docs/problem_1.md) (RU)
- [Problem 2](docs/problem_2.md) (RU)

## Installation

```
$ go get github.com/svetlana-rezvaya/go-food
```

## Usage Example

```go
package main

import (
	"fmt"
	"math"

	food "github.com/svetlana-rezvaya/go-food"
)

func roundToPrecision(value float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Round(value*factor) / factor
}

func main() {
	// the calories and macronutrients are taken from https://fdc.nal.usda.gov/
	dishes := []food.Dish{
		// buckwheat grain
		food.SimpleDish{
			Weight:      50,
			EnergyValue: 343,
			Macros:      food.Macros{Protein: 13.25, Fat: 3.4, Carbs: 71.5},
		},
		// chicken breast
		food.SimpleDish{
			Weight:      150,
			EnergyValue: 165,
			Macros:      food.Macros{Protein: 31, Fat: 3.6, Carbs: 0},
		},
		// vegetable salad
		food.ComplexDish{
			Dishes: []food.Dish{
				// tomato
				food.SimpleDish{
					Weight:      500,
					EnergyValue: 18,
					Macros:      food.Macros{Protein: 0.9, Fat: 0.2, Carbs: 3.9},
				},
				// cucumber
				food.SimpleDish{
					Weight:      500,
					EnergyValue: 16,
					Macros:      food.Macros{Protein: 0.65, Fat: 0.11, Carbs: 3.63},
				},
				// garlic
				food.SimpleDish{
					Weight:      20,
					EnergyValue: 149,
					Macros:      food.Macros{Protein: 6.36, Fat: 0.5, Carbs: 33.06},
				},
				// olive oil
				food.SimpleDish{
					Weight:      100,
					EnergyValue: 884,
					Macros:      food.Macros{Protein: 0, Fat: 100, Carbs: 0},
				},
			},
			PartAsWeight: 200,
		},
		// green tea with honey
		food.ComplexDish{
			Dishes: []food.Dish{
				// green tea
				food.SimpleDish{
					Weight:      300,
					EnergyValue: 0.96,
					Macros:      food.Macros{Protein: 0.2, Fat: 0, Carbs: 0},
				},
				// honey
				food.SimpleDish{
					Weight:      15,
					EnergyValue: 304,
					Macros:      food.Macros{Protein: 0.3, Fat: 0, Carbs: 82},
				},
			},
			PartAsPercent: 1,
		},
	}

	calories := food.CountTotalCalories(dishes)
	macros := food.CountTotalMacros(dishes)

	const precision = 2
	fmt.Printf("calories: %g\n", roundToPrecision(calories, precision))
	fmt.Printf("macros: %+v\n", food.Macros{
		Protein: roundToPrecision(macros.Protein, precision),
		Fat:     roundToPrecision(macros.Fat, precision),
		Carbs:   roundToPrecision(macros.Carbs, precision),
	})

	// Output:
	// calories: 661.02
	// macros: {Protein:55.38 Fat:25.25 Carbs:55.95}
}
```

## License

The MIT License (MIT)

Copyright &copy; 2021 svetlana-rezvaya
