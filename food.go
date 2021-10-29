package food

// CountTotalWeight ...
func CountTotalWeight(dishes []SimpleDish) float64 {
	totalWeight := 0.0
	for _, dish := range dishes {
		totalWeight = totalWeight + dish.Weight
	}

	return totalWeight
}

// CountTotalCalories ...
func CountTotalCalories(dishes []SimpleDish) float64 {
	totalCalories := 0.0
	for _, dish := range dishes {
		calories := dish.CountCalories()
		totalCalories = totalCalories + calories
	}

	return totalCalories
}

// CountTotalMacros ...
func CountTotalMacros(dishes []SimpleDish) Macros {
	totalMacros := Macros{}
	for _, dish := range dishes {
		macros := dish.CountMacros()
		totalMacros = totalMacros.Add(macros)
	}

	return totalMacros
}
