package food

import (
	"reflect"
	"testing"
)

func TestMacros_Add(test *testing.T) {
	macros := Macros{Protein: 10, Fat: 20, Carbs: 30}
	receivedMacros := macros.Add(Macros{Protein: 100, Fat: 200, Carbs: 300})

	wantedMacros := Macros{Protein: 110, Fat: 220, Carbs: 330}
	if !reflect.DeepEqual(receivedMacros, wantedMacros) {
		test.Logf(
			"failed:\n  expected: %+v\n  actual: %+v",
			wantedMacros,
			receivedMacros,
		)
		test.Fail()
	}
}
