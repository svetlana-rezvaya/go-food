package food

// ComplexDish ...
type ComplexDish struct {
	Dishes []Dish
	Part   float64
}

// CountCalories ...
func (complexDish ComplexDish) CountCalories() float64 {
	totalWeight := 0.0
	for _, dish := range complexDish.Dishes {
		totalWeight = totalWeight + dish.Weight
	}

	totalCalories := CountTotalCalories(complexDish.Dishes)

	caloriesPerGram := totalCalories / totalWeight
	portion := totalWeight * complexDish.Part
	return caloriesPerGram * portion
}

// CountMacros ...
func (complexDish ComplexDish) CountMacros() Macros {
	totalWeight := 0.0
	for _, dish := range complexDish.Dishes {
		totalWeight = totalWeight + dish.Weight
	}

	totalMacros := CountTotalMacros(complexDish.Dishes)

	proteinPerGram := totalMacros.Protein / totalWeight
	portion := totalWeight * complexDish.Part
	protein := proteinPerGram * portion

	fatPerGram := totalMacros.Fat / totalWeight
	fat := fatPerGram * portion

	carbsPerGram := totalMacros.Carbs / totalWeight
	carbs := carbsPerGram * portion

	return Macros{Protein: protein, Fat: fat, Carbs: carbs}
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
