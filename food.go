package food

// Dish ...
type Dish struct {
	Weight      int
	EnergyValue int
}

// CountCalories ...
func (dish Dish) CountCalories() int {
	return dish.Weight * dish.EnergyValue / 100
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
