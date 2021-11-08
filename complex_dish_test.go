package food

import (
	"math"
	"reflect"
	"testing"
)

func TestComplexDish_CountWeight_empty(test *testing.T) {
	complexDish := ComplexDish{Dishes: []Dish{}}
	receivedTotalWeight := complexDish.CountWeight()

	wantedTotalWeight := 0.0
	if receivedTotalWeight != wantedTotalWeight {
		test.Logf(
			"failed:\n  expected: %+v\n  actual: %+v",
			wantedTotalWeight,
			receivedTotalWeight,
		)
		test.Fail()
	}
}

func TestComplexDish_CountWeight_nonEmpty(test *testing.T) {
	complexDish := ComplexDish{
		Dishes: []Dish{mockDish{weight: 2.3}, mockDish{weight: 4.2}},
	}
	receivedTotalWeight := complexDish.CountWeight()

	wantedTotalWeight := 6.5
	if receivedTotalWeight != wantedTotalWeight {
		test.Logf(
			"failed:\n  expected: %+v\n  actual: %+v",
			wantedTotalWeight,
			receivedTotalWeight,
		)
		test.Fail()
	}
}

func TestComplexDish_CountCalories_empty(test *testing.T) {
	complexDish := ComplexDish{Dishes: []Dish{}, PartAsPercent: 0.5}
	receivedTotalCalories := complexDish.CountCalories()

	wantedTotalCalories := 0.0
	if receivedTotalCalories != wantedTotalCalories {
		test.Logf(
			"failed:\n  expected: %+v\n  actual: %+v",
			wantedTotalCalories,
			receivedTotalCalories,
		)
		test.Fail()
	}
}

func TestComplexDish_CountCalories_nonEmptyAndPartAsPercent(test *testing.T) {
	complexDish := ComplexDish{
		Dishes: []Dish{
			mockDish{weight: 50, calories: 2.3},
			mockDish{weight: 50, calories: 4.2},
		},
		PartAsPercent: 0.5,
	}
	receivedTotalCalories := complexDish.CountCalories()

	wantedTotalCalories := 3.25
	if receivedTotalCalories != wantedTotalCalories {
		test.Logf(
			"failed:\n  expected: %+v\n  actual: %+v",
			wantedTotalCalories,
			receivedTotalCalories,
		)
		test.Fail()
	}
}

func TestComplexDish_CountCalories_nonEmptyAndPartAsWeight(test *testing.T) {
	complexDish := ComplexDish{
		Dishes: []Dish{
			mockDish{weight: 50, calories: 2.3},
			mockDish{weight: 50, calories: 4.2},
		},
		PartAsWeight: 50,
	}
	receivedTotalCalories := complexDish.CountCalories()

	wantedTotalCalories := 3.25
	if receivedTotalCalories != wantedTotalCalories {
		test.Logf(
			"failed:\n  expected: %+v\n  actual: %+v",
			wantedTotalCalories,
			receivedTotalCalories,
		)
		test.Fail()
	}
}

func TestComplexDish_CountMacros_empty(test *testing.T) {
	complexDish := ComplexDish{Dishes: []Dish{}, PartAsPercent: 0.5}
	receivedTotalMacros := complexDish.CountMacros()

	wantedTotalMacros := Macros{}
	if !reflect.DeepEqual(receivedTotalMacros, wantedTotalMacros) {
		test.Logf(
			"failed:\n  expected: %+v\n  actual: %+v",
			wantedTotalMacros,
			receivedTotalMacros,
		)
		test.Fail()
	}
}

func TestComplexDish_CountMacros_nonEmptyAndPartAsPercent(test *testing.T) {
	complexDish := ComplexDish{
		Dishes: []Dish{
			mockDish{
				weight: 50,
				macros: Macros{Protein: 10, Fat: 20, Carbs: 30},
			},
			mockDish{
				weight: 50,
				macros: Macros{Protein: 100, Fat: 200, Carbs: 300},
			},
		},
		PartAsPercent: 0.5,
	}
	receivedTotalMacros := complexDish.CountMacros()
	receivedTotalMacros = Macros{
		Protein: math.Round(receivedTotalMacros.Protein),
		Fat:     math.Round(receivedTotalMacros.Fat),
		Carbs:   math.Round(receivedTotalMacros.Carbs),
	}

	wantedTotalMacros := Macros{Protein: 55, Fat: 110, Carbs: 165}
	if !reflect.DeepEqual(receivedTotalMacros, wantedTotalMacros) {
		test.Logf(
			"failed:\n  expected: %+v\n  actual: %+v",
			wantedTotalMacros,
			receivedTotalMacros,
		)
		test.Fail()
	}
}

func TestComplexDish_CountMacros_nonEmptyAndPartAsWeight(test *testing.T) {
	complexDish := ComplexDish{
		Dishes: []Dish{
			mockDish{
				weight: 50,
				macros: Macros{Protein: 10, Fat: 20, Carbs: 30},
			},
			mockDish{
				weight: 50,
				macros: Macros{Protein: 100, Fat: 200, Carbs: 300},
			},
		},
		PartAsWeight: 50,
	}
	receivedTotalMacros := complexDish.CountMacros()
	receivedTotalMacros = Macros{
		Protein: math.Round(receivedTotalMacros.Protein),
		Fat:     math.Round(receivedTotalMacros.Fat),
		Carbs:   math.Round(receivedTotalMacros.Carbs),
	}

	wantedTotalMacros := Macros{Protein: 55, Fat: 110, Carbs: 165}
	if !reflect.DeepEqual(receivedTotalMacros, wantedTotalMacros) {
		test.Logf(
			"failed:\n  expected: %+v\n  actual: %+v",
			wantedTotalMacros,
			receivedTotalMacros,
		)
		test.Fail()
	}
}
