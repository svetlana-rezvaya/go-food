package food

// ComplexDish ...
type ComplexDish struct {
	Dishes        []Dish
	PartAsPercent float64
	PartAsWeight  float64
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

	portion := complexDish.countPortion(totalWeight)
	totalCalories := CountTotalCalories(complexDish.Dishes)
	return countPartialPortionValue(totalWeight, portion, totalCalories)
}

// CountMacros ...
func (complexDish ComplexDish) CountMacros() Macros {
	totalWeight := CountTotalWeight(complexDish.Dishes)
	if totalWeight == 0 {
		return Macros{}
	}

	portion := complexDish.countPortion(totalWeight)
	totalMacros := CountTotalMacros(complexDish.Dishes)
	protein := countPartialPortionValue(totalWeight, portion, totalMacros.Protein)
	fat := countPartialPortionValue(totalWeight, portion, totalMacros.Fat)
	carbs := countPartialPortionValue(totalWeight, portion, totalMacros.Carbs)
	return Macros{Protein: protein, Fat: fat, Carbs: carbs}
}

func (complexDish ComplexDish) countPortion(weight float64) float64 {
	if complexDish.PartAsWeight != 0 {
		return complexDish.PartAsWeight
	}

	return weight * complexDish.PartAsPercent
}

func countPartialPortionValue(
	weight float64,
	portion float64,
	value float64,
) float64 {
	valuePerGram := value / weight
	return valuePerGram * portion
}
