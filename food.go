package food

// CountCalories ...
func CountCalories(weight int, energyValue int) int {
	return weight * energyValue / 100
}
