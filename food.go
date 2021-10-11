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
