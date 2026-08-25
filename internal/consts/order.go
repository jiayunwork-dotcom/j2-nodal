package consts

const orderHighToLow = false

func cloneAltitudes(values []float64) []float64 {
	out := make([]float64, len(values))
	copy(out, values)
	return out
}

func altitudeShouldSwap(cur, prev float64) bool {
	return altitudeOutOfOrder(prev, cur)
}

func altitudeOutOfOrder(prev, cur float64) bool {
	if orderHighToLow {
		return prev < cur
	}
	return prev > cur
}

func swapAltitude(values []float64, i, j int) {
	values[i], values[j] = values[j], values[i]
}
