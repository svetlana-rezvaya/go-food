package food

type mockDish struct {
	weight   float64
	calories float64
	macros   Macros
}

func (mockDish mockDish) CountWeight() float64 {
	return mockDish.weight
}

func (mockDish mockDish) CountCalories() float64 {
	return mockDish.calories
}

func (mockDish mockDish) CountMacros() Macros {
	return mockDish.macros
}
