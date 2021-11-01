package food

import (
	"reflect"
	"testing"
)

func TestSimpleDish_CountWeight(test *testing.T) {
	weight := 2.3
	simpleDish := SimpleDish{Weight: weight}
	receivedWeight := simpleDish.CountWeight()

	if receivedWeight != weight {
		test.Logf("failed:\n  expected: %+v\n  actual: %+v", weight, receivedWeight)
		test.Fail()
	}
}

func TestSimpleDish_CountCalories(test *testing.T) {
	simpleDish := SimpleDish{Weight: 50, EnergyValue: 300}
	receivedCalories := simpleDish.CountCalories()

	wantedCalories := 150.0
	if receivedCalories != wantedCalories {
		test.Logf(
			"failed:\n  expected: %+v\n  actual: %+v",
			wantedCalories,
			receivedCalories,
		)
		test.Fail()
	}
}

func TestSimpleDish_CountMacros(test *testing.T) {
	simpleDish := SimpleDish{
		Weight: 50,
		Macros: Macros{Protein: 10, Fat: 20, Carbs: 30},
	}
	receivedMacros := simpleDish.CountMacros()

	wantedMacros := Macros{Protein: 5, Fat: 10, Carbs: 15}
	if !reflect.DeepEqual(receivedMacros, wantedMacros) {
		test.Logf(
			"failed:\n  expected: %+v\n  actual: %+v",
			wantedMacros,
			receivedMacros,
		)
		test.Fail()
	}
}
