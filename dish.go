package food

// Dish ...
type Dish struct {
	Weight      float64
	EnergyValue float64
	Macros      Macros
}

// CountCalories ...
func (dish Dish) CountCalories() float64 {
	return countPortionValue(dish.Weight, dish.EnergyValue)
}

// CountMacros ...
func (dish Dish) CountMacros() Macros {
	protein := countPortionValue(dish.Weight, dish.Macros.Protein)
	fat := countPortionValue(dish.Weight, dish.Macros.Fat)
	carbs := countPortionValue(dish.Weight, dish.Macros.Carbs)
	return Macros{Protein: protein, Fat: fat, Carbs: carbs}
}

func countPortionValue(weight float64, value float64) float64 {
	return weight * value / 100
}
