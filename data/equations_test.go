package data

import (
	"testing"
)

var testcases = []struct {
	name     string
	expected float64
	arr      []float64
}{
	{"mean", 6.0, []float64{4.0, 5.0, 9.0, 7.0, 5.0, 6.0}},
	{"minimum", 6.0, []float64{4.0, 5.0, 9.0, 7.0, 5.0, 6.0}},
	{"maximum", 6.0, []float64{4.0, 5.0, 9.0, 7.0, 5.0, 6.0}},
}

func TestMinMax(t *testing.T) {
	//arrange
	//act
	//assert

}

func TestStandardise(t *testing.T) {
}

func TestEquations(t *testing.T) {
	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			expected := test.expected
			got := mean(test.arr)
			if got != expected {
				t.Errorf("expected: %.1f, got: %.1f", expected, got)
			}
		})
	}

}

func TestStandardDeviation(t *testing.T) {
	expected := 7.905694150420948

	got := standardDeviation([]float64{5.0, 10.0, 15.0, 20.0, 25.0})

	if got != expected {
		t.Errorf("expected: %.1f, got: %.1f", expected, got)
	}
}
