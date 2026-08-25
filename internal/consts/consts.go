package consts

import (
	"fmt"
	"math"

	"j2-nodal/internal/j2"
)

const (
	AU           = 1.495978707e11
	SpeedOfLight = 2.99792458e8
	SolarDay     = 86400.0
	SiderealDay  = 86164.0905
)

func SiderealDayHours() float64 {
	return SiderealDay / 3600
}

func SolarDayHours() float64 {
	return 24
}

func MeanMotionRadPerSec(a float64) float64 {
	return j2.MeanMotion(a)
}

func MeanMotionDegPerDay(a float64) float64 {
	return MeanMotionRadPerSec(a) * SolarDay * 180 / math.Pi
}

func PeriodHours(a float64) float64 {
	return 2 * math.Pi * math.Sqrt(math.Pow(a, 3)/j2.Mu) / 3600
}

func AltitudeOf(a float64) float64 {
	return a - j2.EarthRadius
}

func RadiusAtAltitude(altitude float64) float64 {
	return j2.EarthRadius + altitude
}

func EccentricityFactor(e float64) float64 {
	return (1 - e*e) * (1 - e*e)
}

func J2Factor(a, e float64) float64 {
	return j2.J2 * math.Pow(j2.EarthRadius/a, 2) / EccentricityFactor(e)
}

func Summary(a, e, i float64) string {
	return fmt.Sprintf("a=%.3f km e=%.4g i=%.4g deg n=%.6f deg/day",
		a/1000, e, i, MeanMotionDegPerDay(a))
}

func IsValidInclination(i float64) bool {
	return i >= 0 && i <= 180
}

func ClampInclination(i float64) float64 {
	if i < 0 {
		return 0
	}
	if i > 180 {
		return 180
	}
	return i
}

func ClampEccentricity(e float64) float64 {
	if e < 0 {
		return 0
	}
	if e >= 1 {
		return 0.999
	}
	return e
}

func ClampAltitude(altitude float64) float64 {
	if altitude < 100000 {
		return 100000
	}
	return altitude
}

func IsNearPolar(i float64) bool {
	return math.Abs(i-90) < 5
}

func IsNearEquatorial(i float64) bool {
	return i < 5 || i > 175
}

func IsSunSyncRange(i float64) bool {
	return i > 90 && i < 110
}

func DescribeConstants() string {
	return j2.ConstantsText()
}

func AltitudeKm(a float64) float64 {
	return AltitudeOf(a) / 1000
}

func SemimajorKm(a float64) float64 {
	return a / 1000
}

func Display(a, e, i float64) string {
	return Summary(a, e, i)
}

func MaxAltitude(altitudes []float64) float64 {
	max := 0.0
	for _, altitude := range altitudes {
		if altitude > max {
			max = altitude
		}
	}
	return max
}

func MinAltitude(altitudes []float64) float64 {
	if len(altitudes) == 0 {
		return 0
	}
	min := altitudes[0]
	for _, altitude := range altitudes {
		if altitude < min {
			min = altitude
		}
	}
	return min
}

func SortAltitudes(altitudes []float64) []float64 {
	out := cloneAltitudes(altitudes)
	for i := 1; i < len(out); i++ {
		for j := i; j > 0 && altitudeShouldSwap(out[j], out[j-1]); j-- {
			swapAltitude(out, j, j-1)
		}
	}
	return out
}

func EqualAltitudes(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func SumAltitudes(altitudes []float64) float64 {
	sum := 0.0
	for _, altitude := range altitudes {
		sum += altitude
	}
	return sum
}

func MeanAltitude(altitudes []float64) float64 {
	if len(altitudes) == 0 {
		return 0
	}
	return SumAltitudes(altitudes) / float64(len(altitudes))
}

func MedianAltitude(altitudes []float64) float64 {
	sorted := SortAltitudes(altitudes)
	n := len(sorted)
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return sorted[n/2]
	}
	return (sorted[n/2-1] + sorted[n/2]) / 2
}

func AltitudeRange(altitudes []float64) float64 {
	return MaxAltitude(altitudes) - MinAltitude(altitudes)
}

func Text(altitudes []float64) string {
	out := ""
	for _, altitude := range SortAltitudes(altitudes) {
		out += fmt.Sprintf("%.3f km\n", altitude/1000)
	}
	return out
}

func IsFinite(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
}

func RequireFinite(value float64, name string) error {
	if !IsFinite(value) {
		return fmt.Errorf("%s must be finite", name)
	}
	return nil
}
