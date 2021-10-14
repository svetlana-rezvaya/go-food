package food

// Macros ...
type Macros struct {
	Protein float64
	Fat     float64
	Carbs   float64
}

// Add ...
func (macros Macros) Add(anotherMacros Macros) Macros {
	protein := macros.Protein + anotherMacros.Protein
	fat := macros.Fat + anotherMacros.Fat
	carbs := macros.Carbs + anotherMacros.Carbs
	return Macros{Protein: protein, Fat: fat, Carbs: carbs}
}

// Dish ...
type Dish struct {
	Weight      float64
	EnergyValue float64
	Macros      Macros
}

// CountCalories ...
func (dish Dish) CountCalories() float64 {
	return dish.Weight * dish.EnergyValue / 100
}

// CountMacros ...
func (dish Dish) CountMacros() Macros {
	protein := dish.Weight * dish.Macros.Protein / 100
	fat := dish.Weight * dish.Macros.Fat / 100
	carbs := dish.Weight * dish.Macros.Carbs / 100
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
