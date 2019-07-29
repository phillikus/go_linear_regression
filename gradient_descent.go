package gradient_descent

func regression(x_values []float32, y_values []float32, epochs int, learning_rate float32) (float32, float32) {
	var b_current float32 = 0
	var m_current float32 = 0

	for i := 0; i < epochs; i++ {
		b_current, m_current = step(b_current, m_current, x_values, y_values, learning_rate)
	}

	return b_current, m_current
}

func step(b_current float32, m_current float32, x_values []float32, y_values []float32, learning_rate float32) (float32, float32) {
	var b_gradient float32 = 0
	var m_gradient float32 = 0

	var length = len(y_values)

	for i := 0; i < length; i++ {
		var two_over_n = float32((2 / length))
		b_gradient += -two_over_n * (y_values[i] - ((m_current * x_values[i]) + b_current))
		m_gradient += -two_over_n * x_values[i] * (y_values[i] - ((m_current * x_values[i]) + b_current))
	}

	var new_b = b_current - (learning_rate * b_gradient)
	var new_m = m_current - (learning_rate * m_gradient)

	return new_b, new_m
}
