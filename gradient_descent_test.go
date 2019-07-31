package main

import "math"
import "testing"

func Test_Regression_with_empty_arrays(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Regression with empty arrays did not panic.")
		}
	}()

	Regression([]float32{}, []float32{}, 200, 0.5)
}

func Test_Regression_with_non_empty_arrays(t *testing.T) {
	x_values := []float32{1, 2, 3, 4, 5}
	y_values := []float32{1, 3, 2, 3, 5}

	intercept, coefficient := Regression(x_values, y_values, 200, 0.5)
	assertDelta(t, 0.4, intercept, 0.001)
	assertDelta(t, 0.8, coefficient, 0.001)
}

func assertDelta(t *testing.T, x float32, y float32, delta float32) {
	abs_difference := math.Abs(float64(x - y))
	if abs_difference > float64(delta) {
		t.Error("Difference between x and y is greater than delta!")
	}
}
