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
