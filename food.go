package food

// Dish ...
type Dish struct {
	Weight      int
	EnergyValue int
}

// CountCalories ...
func CountCalories(dish Dish) int {
	return dish.Weight * dish.EnergyValue / 100
}

// CountTotalCalories ...
func CountTotalCalories(dishes []Dish) int {
	totalCalories := 0
	for _, dish := range dishes {
		calories := CountCalories(dish)
		totalCalories = totalCalories + calories
	}

	return totalCalories
}
