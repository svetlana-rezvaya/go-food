package food

// Dish ...
type Dish interface {
	CountWeight() float64
	CountCalories() float64
	CountMacros() Macros
}

// CountTotalWeight ...
func CountTotalWeight(dishes []Dish) float64 {
	totalWeight := 0.0
	for _, dish := range dishes {
		weight := dish.CountWeight()
		totalWeight = totalWeight + weight
	}

	return totalWeight
}

// CountTotalCalories ...
func CountTotalCalories(dishes []Dish) float64 {
	totalCalories := 0.0
	for _, dish := range dishes {
		calories := dish.CountCalories()
		totalCalories = totalCalories + calories
	}

	return totalCalories
}

// CountTotalMacros ...
func CountTotalMacros(dishes []Dish) Macros {
	totalMacros := Macros{}
	for _, dish := range dishes {
		macros := dish.CountMacros()
		totalMacros = totalMacros.Add(macros)
	}

	return totalMacros
}
