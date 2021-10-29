package food

// SimpleDish ...
type SimpleDish struct {
	Weight      float64
	EnergyValue float64
	Macros      Macros
}

// CountCalories ...
func (simpleDish SimpleDish) CountCalories() float64 {
	return countPortionValue(simpleDish.Weight, simpleDish.EnergyValue)
}

// CountMacros ...
func (simpleDish SimpleDish) CountMacros() Macros {
	protein := countPortionValue(simpleDish.Weight, simpleDish.Macros.Protein)
	fat := countPortionValue(simpleDish.Weight, simpleDish.Macros.Fat)
	carbs := countPortionValue(simpleDish.Weight, simpleDish.Macros.Carbs)
	return Macros{Protein: protein, Fat: fat, Carbs: carbs}
}

func countPortionValue(weight float64, value float64) float64 {
	return weight * value / 100
}
