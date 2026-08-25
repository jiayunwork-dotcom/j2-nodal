package analysis

import (
	"context"
	"fmt"
	"math"

	"j2-nodal/internal/j2"
)

func ScanSemimajor(e, i float64, semimajors []float64) ([]j2.Rates, error) {
	out := make([]j2.Rates, 0, len(semimajors))
	for _, a := range semimajors {
		rate, err := j2.PrecessionRates(a, e, i)
		if err != nil {
			return nil, err
		}
		out = append(out, rate)
	}
	return out, nil
}

func ScanEccentricity(a, i float64, eccentricities []float64) ([]j2.Rates, error) {
	out := make([]j2.Rates, 0, len(eccentricities))
	for _, e := range eccentricities {
		rate, err := j2.PrecessionRates(a, e, i)
		if err != nil {
			return nil, err
		}
		out = append(out, rate)
	}
	return out, nil
}

func ScanInclination(a, e float64, inclinations []float64) ([]j2.Rates, error) {
	out := make([]j2.Rates, 0, len(inclinations))
	for _, i := range inclinations {
		rate, err := j2.PrecessionRates(a, e, i)
		if err != nil {
			return nil, err
		}
		out = append(out, rate)
	}
	return out, nil
}

func CancellableScan(ctx context.Context, e, i float64, semimajors []float64) ([]j2.Rates, error) {
	out := make([]j2.Rates, 0, len(semimajors))
	for _, a := range semimajors {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		rate, err := j2.PrecessionRates(a, e, i)
		if err != nil {
			return nil, err
		}
		out = append(out, rate)
	}
	return out, nil
}

func UniformSemimajor(startAlt, endAlt float64, n int) []float64 {
	values := make([]float64, 0, n)
	for idx := 0; idx < n; idx++ {
		alt := startAlt + float64(idx)*(endAlt-startAlt)/float64(n-1)
		values = append(values, j2.EarthRadius+alt)
	}
	return values
}

func UniformEccentricity(start, end float64, n int) []float64 {
	values := make([]float64, 0, n)
	for idx := 0; idx < n; idx++ {
		values = append(values, start+float64(idx)*(end-start)/float64(n-1))
	}
	return values
}

func UniformInclination(start, end float64, n int) []float64 {
	return UniformEccentricity(start, end, n)
}

func MaxRAAN(rates []j2.Rates) float64 {
	max := 0.0
	for _, rate := range rates {
		abs := math.Abs(rate.RAANDot)
		if abs > max {
			max = abs
		}
	}
	return max
}

func MinRAAN(rates []j2.Rates) float64 {
	if len(rates) == 0 {
		return 0
	}
	min := math.Abs(rates[0].RAANDot)
	for _, rate := range rates {
		abs := math.Abs(rate.RAANDot)
		if abs < min {
			min = abs
		}
	}
	return min
}

func Summary(rates []j2.Rates) string {
	if len(rates) == 0 {
		return "empty"
	}
	return fmt.Sprintf("RAAN_dot=%.6g..%.6g deg/day",
		MinRAAN(rates), MaxRAAN(rates))
}

func Text(rates []j2.Rates) string {
	out := ""
	for _, rate := range rates {
		out += fmt.Sprintf("%.4g %.4g %.4g %.6g\n",
			rate.A, rate.E, rate.I, rate.RAANDot)
	}
	return out
}

func IsMonotonicRAAN(rates []j2.Rates) bool {
	for i := 1; i < len(rates); i++ {
		if math.Abs(rates[i].RAANDot) < math.Abs(rates[i-1].RAANDot) {
			return false
		}
	}
	return true
}

func FirstSSOIndex(rates []j2.Rates) int {
	for i, rate := range rates {
		if math.Abs(rate.RAANDot-j2.SSOTargetDegPerDay()) < 0.01 {
			return i
		}
	}
	return -1
}

func IsSSORate(rate float64, tolerance float64) bool {
	return math.Abs(rate-j2.SSOTargetDegPerDay()) <= tolerance
}

func AverageRAAN(rates []j2.Rates) float64 {
	if len(rates) == 0 {
		return 0
	}
	sum := 0.0
	for _, rate := range rates {
		sum += rate.RAANDot
	}
	return sum / float64(len(rates))
}

func StddevRAAN(rates []j2.Rates) float64 {
	if len(rates) < 2 {
		return 0
	}
	mean := AverageRAAN(rates)
	variance := 0.0
	for _, rate := range rates {
		delta := rate.RAANDot - mean
		variance += delta * delta
	}
	return math.Sqrt(variance / float64(len(rates)-1))
}

func IsFinite(rates []j2.Rates) bool {
	for _, rate := range rates {
		if math.IsNaN(rate.RAANDot) || math.IsInf(rate.RAANDot, 0) {
			return false
		}
	}
	return true
}

func RateRange(rates []j2.Rates) float64 {
	if len(rates) == 0 {
		return 0
	}
	max, min := rates[0].RAANDot, rates[0].RAANDot
	for _, rate := range rates {
		if rate.RAANDot > max {
			max = rate.RAANDot
		}
		if rate.RAANDot < min {
			min = rate.RAANDot
		}
	}
	return max - min
}

func AllFinite(rates []j2.Rates) bool {
	return IsFinite(rates)
}

func SummaryStats(rates []j2.Rates) string {
	return fmt.Sprintf("n=%d avg=%.6g std=%.6g range=%.6g",
		len(rates), AverageRAAN(rates), StddevRAAN(rates), RateRange(rates))
}
