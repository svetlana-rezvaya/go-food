package food

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
