package food

// Dish ...
type Dish struct {
	Weight      int
	EnergyValue int
}

// CountCalories ...
func CountCalories(weight int, energyValue int) int {
	return weight * energyValue / 100
}

// CountTotalCalories ...
func CountTotalCalories(dishes []Dish) int {
	totalCalories := 0
	for _, dish := range dishes {
		calories := CountCalories(dish.Weight, dish.EnergyValue)
		totalCalories = totalCalories + calories
	}

	return totalCalories
}
