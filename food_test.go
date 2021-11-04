package food

import "testing"

func TestCountTotalWeight_empty(test *testing.T) {
	receivedTotalWeight := CountTotalWeight([]Dish{})

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

func TestCountTotalWeight_nonEmpty(test *testing.T) {
	dishes := []Dish{mockDish{weight: 2.3}, mockDish{weight: 4.2}}
	receivedTotalWeight := CountTotalWeight(dishes)

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
