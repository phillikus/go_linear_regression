package main

import "fmt"

func main() {
	x_values := []float32{1, 2, 3, 4, 5}
	y_values := []float32{1, 3, 2, 3, 5}

	intercept, coefficient := Regression(x_values, y_values, 200, 0.05)

	fmt.Printf("Intercept: %f\n", intercept)
	fmt.Printf("Coefficient: %f\n", coefficient)
}
