package testing

import (
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.0001
}

func TestSqrt(t *testing.T) {
	val, err := Sqrt(2)
	if err != nil {
		t.Errorf("Sqrt(2) returned an error: %v", err)
	}
	if !almostEqual(val, 1.4142) {
		t.Errorf("Sqrt(2) = %v; want 1.4142", val)
	}
}


// TESGING WITH MULTIPLE CASES
type testCase struct {
	value    float64
	expected float64
}

func TestMany(t *testing.T) {
	testCases := []testCase{
		{4, 2},
		{9, 3},
		{16, 4},
		{25, 5},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) { // subtest
			val, err := Sqrt(tc.value)
			if err != nil {
				t.Errorf("Sqrt(%v) returned an error: %v", tc.value, err)
			}
			if !almostEqual(val, tc.expected) {
				t.Errorf("Sqrt(%v) = %v; want %v", tc.value, val, tc.expected)
			}
		})
	}
}