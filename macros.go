package food

// Macros ...
type Macros struct {
	Protein float64
	Fat     float64
	Carbs   float64
}

// Add ...
func (macros Macros) Add(anotherMacros Macros) Macros {
	protein := macros.Protein + anotherMacros.Protein
	fat := macros.Fat + anotherMacros.Fat
	carbs := macros.Carbs + anotherMacros.Carbs
	return Macros{Protein: protein, Fat: fat, Carbs: carbs}
}
