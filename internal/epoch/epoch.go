package epoch

import (
	"fmt"
	"math"

	"j2-nodal/internal/j2"
)

func MeanAnomalyRate(a float64) float64 {
	return j2.MeanMotion(a)
}

func NodalRegression(a, e, i float64) (float64, error) {
	return j2.RAANOnly(a, e, i)
}

func ArgumentRate(a, e, i float64) (float64, error) {
	return j2.ArgPeriOnly(a, e, i)
}

func RAANAfterDays(a, e, i, days float64) (float64, error) {
	rate, err := j2.RAANOnly(a, e, i)
	if err != nil {
		return 0, err
	}
	return rate * days, nil
}

func ArgumentAfterDays(a, e, i, days float64) (float64, error) {
	rate, err := j2.ArgPeriOnly(a, e, i)
	if err != nil {
		return 0, err
	}
	return rate * days, nil
}

func SSOInclination(a, e float64) (float64, error) {
	result, err := j2.SunSyncInclination(a, e)
	if err != nil {
		return 0, err
	}
	if !result.Possible {
		return 0, j2.NoSSOError()
	}
	return result.ISSO, nil
}

func OrbitalPeriodDays(a float64) float64 {
	return 2 * math.Pi * math.Sqrt(math.Pow(a, 3)/j2.Mu) / j2.SecondsPerDay
}

func NodalPeriodDays(a, e float64) (float64, error) {
	if err := j2.ValidateSemimajor(a); err != nil {
		return 0, err
	}
	if err := j2.ValidateEccentricity(e); err != nil {
		return 0, err
	}
	return j2.SiderealPeriodDays(a), nil
}

func RevolutionsPerDay(a float64) float64 {
	return 1 / OrbitalPeriodDays(a)
}

func DescendingNodeShift(a, e, days float64) (float64, error) {
	rate, err := j2.RAANOnly(a, e, 97.8)
	if err != nil {
		return 0, err
	}
	return rate * days, nil
}

func SSOResidual(a, e float64) (float64, error) {
	i, err := SSOInclination(a, e)
	if err != nil {
		return 0, err
	}
	rate, err := j2.RAANOnly(a, e, i)
	if err != nil {
		return 0, err
	}
	return rate - j2.SSOTargetDegPerDay(), nil
}

func DriftPerYear(a, e float64) (float64, error) {
	residual, err := SSOResidual(a, e)
	if err != nil {
		return 0, err
	}
	return residual * j2.DaysPerYear, nil
}

func Summary(a, e float64) string {
	period := OrbitalPeriodDays(a)
	return fmt.Sprintf("a=%.3f km period=%.4f days rev/day=%.4f",
		a/1000, period, 1/period)
}

func Describe(a, e, i float64) string {
	raan, _ := j2.RAANOnly(a, e, i)
	arg, _ := j2.ArgPeriOnly(a, e, i)
	return fmt.Sprintf("RAAN_dot=%.6f arg_dot=%.6f deg/day", raan, arg)
}

func IsSSOOrbit(a, e float64) bool {
	i, err := SSOInclination(a, e)
	if err != nil {
		return false
	}
	rate, _ := j2.RAANOnly(a, e, i)
	return math.Abs(rate-j2.SSOTargetDegPerDay()) < 1e-6
}

func LocalTimeShift(a, e float64) (float64, error) {
	drift, err := DriftPerYear(a, e)
	if err != nil {
		return 0, err
	}
	return drift / 15, nil
}

func EpochText(epoch float64) string {
	return fmt.Sprintf("epoch %.6f days", epoch)
}

func NodeText(a, e, i float64) string {
	rate, _ := j2.RAANOnly(a, e, i)
	return fmt.Sprintf("%.6f deg/day", rate)
}

func ArgumentText(a, e, i float64) string {
	rate, _ := j2.ArgPeriOnly(a, e, i)
	return fmt.Sprintf("%.6f deg/day", rate)
}

func IsRetrogradeSSO(a, e float64) bool {
	i, _ := SSOInclination(a, e)
	return i > 90
}

func AltitudeOf(a float64) float64 {
	return a - j2.EarthRadius
}

func SemimajorForAltitude(altitude float64) float64 {
	return j2.EarthRadius + altitude
}

func DisplayAltitude(a float64) string {
	return fmt.Sprintf("%.3f km", AltitudeOf(a)/1000)
}

func DisplayPeriod(a float64) string {
	return fmt.Sprintf("%.4f days", OrbitalPeriodDays(a))
}

func DisplayRate(rate float64) string {
	return fmt.Sprintf("%.6f deg/day", rate)
}

func DisplayResidual(residual float64) string {
	return fmt.Sprintf("%.6f deg/day", residual)
}

func DisplayDrift(drift float64) string {
	return fmt.Sprintf("%.6f deg/year", drift)
}

func ScanSemimajor(e, i float64, semimajors []float64) ([]float64, error) {
	rates := make([]float64, 0, len(semimajors))
	for _, a := range semimajors {
		rate, err := j2.RAANOnly(a, e, i)
		if err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, nil
}

func ScanEccentricity(a, i float64, eccentricities []float64) ([]float64, error) {
	rates := make([]float64, 0, len(eccentricities))
	for _, e := range eccentricities {
		rate, err := j2.RAANOnly(a, e, i)
		if err != nil {
			return nil, err
		}
		rates = append(rates, rate)
	}
	return rates, nil
}

func MaxAbs(rates []float64) float64 {
	max := 0.0
	for _, rate := range rates {
		abs := math.Abs(rate)
		if abs > max {
			max = abs
		}
	}
	return max
}

func MinAbs(rates []float64) float64 {
	if len(rates) == 0 {
		return 0
	}
	min := math.Abs(rates[0])
	for _, rate := range rates {
		abs := math.Abs(rate)
		if abs < min {
			min = abs
		}
	}
	return min
}

func Average(rates []float64) float64 {
	if len(rates) == 0 {
		return 0
	}
	sum := 0.0
	for _, rate := range rates {
		sum += rate
	}
	return sum / float64(len(rates))
}

func Stddev(rates []float64) float64 {
	if len(rates) < 2 {
		return 0
	}
	mean := Average(rates)
	variance := 0.0
	for _, rate := range rates {
		delta := rate - mean
		variance += delta * delta
	}
	return math.Sqrt(variance / float64(len(rates)-1))
}

func RateRange(rates []float64) float64 {
	if len(rates) == 0 {
		return 0
	}
	max, min := rates[0], rates[0]
	for _, rate := range rates {
		if rate > max {
			max = rate
		}
		if rate < min {
			min = rate
		}
	}
	return max - min
}

func IsMonotonicAbs(rates []float64) bool {
	for i := 1; i < len(rates); i++ {
		if math.Abs(rates[i]) < math.Abs(rates[i-1]) {
			return false
		}
	}
	return true
}

func SummaryRates(rates []float64) string {
	return fmt.Sprintf("n=%d avg=%.6g std=%.6g range=%.6g",
		len(rates), Average(rates), Stddev(rates), RateRange(rates))
}

func TextRates(rates []float64) string {
	out := ""
	for _, rate := range rates {
		out += fmt.Sprintf("%.6f\n", rate)
	}
	return out
}

func AllFinite(rates []float64) bool {
	for _, rate := range rates {
		if math.IsNaN(rate) || math.IsInf(rate, 0) {
			return false
		}
	}
	return true
}
