package food

// Dish ...
type Dish struct {
	Weight      int
	EnergyValue int
	Protein     int
	Fat         int
	Carbs       int
}

// CountCalories ...
func (dish Dish) CountCalories() int {
	return dish.Weight * dish.EnergyValue / 100
}

// CountMacros ...
func (dish Dish) CountMacros() (int, int, int) {
	protein := dish.Weight * dish.Protein / 100
	fat := dish.Weight * dish.Fat / 100
	carbs := dish.Weight * dish.Carbs / 100
	return protein, fat, carbs
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
