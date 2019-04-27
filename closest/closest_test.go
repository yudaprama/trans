package closest

import "testing"

var arr = []float64{-1.5, 0, 4.4, 5, 6, 7}

var testCases = []struct {
	target   float64
	expected float64
}{
	{4.5, 4.4},
	{5.5, 6},
}

func TestSearchClosest(t *testing.T) {
	for _, tc := range testCases {
		actual := SearchClosest(tc.target, arr)
		if actual != tc.expected {
			t.Errorf("%f: actual (%f), expected (%f)", tc.target, actual, tc.expected)
		}
	}
}
