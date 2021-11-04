package food

// ComplexDish ...
type ComplexDish struct {
	Dishes []Dish
	Part   float64
}

// CountWeight ...
func (complexDish ComplexDish) CountWeight() float64 {
	return CountTotalWeight(complexDish.Dishes)
}

// CountCalories ...
func (complexDish ComplexDish) CountCalories() float64 {
	totalWeight := CountTotalWeight(complexDish.Dishes)
	if totalWeight == 0 {
		return 0
	}

	totalCalories := CountTotalCalories(complexDish.Dishes)
	return countPartialPortionValue(totalWeight, complexDish.Part, totalCalories)
}

// CountMacros ...
func (complexDish ComplexDish) CountMacros() Macros {
	totalWeight := CountTotalWeight(complexDish.Dishes)
	if totalWeight == 0 {
		return Macros{}
	}

	totalMacros := CountTotalMacros(complexDish.Dishes)
	protein :=
		countPartialPortionValue(totalWeight, complexDish.Part, totalMacros.Protein)
	fat := countPartialPortionValue(totalWeight, complexDish.Part, totalMacros.Fat)
	carbs :=
		countPartialPortionValue(totalWeight, complexDish.Part, totalMacros.Carbs)
	return Macros{Protein: protein, Fat: fat, Carbs: carbs}
}

func countPartialPortionValue(
	weight float64,
	part float64,
	value float64,
) float64 {
	valuePerGram := value / weight
	portion := weight * part
	return valuePerGram * portion
}
