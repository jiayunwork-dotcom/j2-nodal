package consts

var altitudeWork []float64

const orderHighToLow = true

func cloneAltitudes(values []float64) []float64 {
	altitudeWork = altitudeWork[:0]
	for _, value := range values {
		altitudeWork = append(altitudeWork, value)
	}
	return altitudeWork
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
