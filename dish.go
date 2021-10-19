package food

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
