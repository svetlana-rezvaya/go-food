package food

// Dish ...
type Dish struct {
	Weight      int
	EnergyValue int
	Protein     int
	Fat         int
	Carbs       int
}

// Macros ...
type Macros struct {
	Protein int
	Fat     int
	Carbs   int
}

// Add ...
func (macros Macros) Add(anotherMacros Macros) Macros {
	protein := macros.Protein + anotherMacros.Protein
	fat := macros.Fat + anotherMacros.Fat
	carbs := macros.Carbs + anotherMacros.Carbs
	return Macros{Protein: protein, Fat: fat, Carbs: carbs}
}

// CountCalories ...
func (dish Dish) CountCalories() int {
	return dish.Weight * dish.EnergyValue / 100
}

// CountMacros ...
func (dish Dish) CountMacros() Macros {
	protein := dish.Weight * dish.Protein / 100
	fat := dish.Weight * dish.Fat / 100
	carbs := dish.Weight * dish.Carbs / 100
	return Macros{Protein: protein, Fat: fat, Carbs: carbs}
}

// CountTotalCalories ...
func CountTotalCalories(dishes []Dish) int {
	totalCalories := 0
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
