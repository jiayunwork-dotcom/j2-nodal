package earth

import (
	"fmt"
	"math"

	"j2-nodal/internal/j2"
)

func Altitude(a float64) float64 {
	return a - j2.EarthRadius
}

func Semimajor(altitude float64) float64 {
	return j2.EarthRadius + altitude
}

func GravityAtAltitude(a float64) float64 {
	return j2.Mu / (a * a)
}

func CircularVelocity(a float64) float64 {
	return math.Sqrt(j2.Mu / a)
}

func Period(a float64) float64 {
	return 2 * math.Pi * math.Sqrt(math.Pow(a, 3)/j2.Mu)
}

func PeriodMinutes(a float64) float64 {
	return Period(a) / 60
}

func SurfaceGravity() float64 {
	return GravityAtAltitude(j2.EarthRadius)
}

func OrbitRate(a float64) float64 {
	return j2.MeanMotion(a)
}

func NodePeriod(a, e float64) (float64, error) {
	rate, err := j2.RAANOnly(a, e, 97.8)
	if err != nil {
		return 0, err
	}
	if rate == 0 {
		return math.Inf(1), nil
	}
	return 360 / math.Abs(rate), nil
}

func IsLEO(a float64) bool {
	alt := Altitude(a)
	return alt >= 200000 && alt <= 2000000
}

func IsMEO(a float64) bool {
	alt := Altitude(a)
	return alt > 2000000 && alt <= 35786000
}

func IsGEO(a float64) bool {
	alt := Altitude(a)
	return math.Abs(alt-35786000) < 100000
}

func Regime(a float64) string {
	switch {
	case IsLEO(a):
		return "LEO"
	case IsMEO(a):
		return "MEO"
	case IsGEO(a):
		return "GEO"
	default:
		return "high"
	}
}

func Summary(a float64) string {
	return fmt.Sprintf("a=%.3f km alt=%.3f km T=%.2f min regime=%s",
		a/1000, Altitude(a)/1000, PeriodMinutes(a), Regime(a))
}

func Velocities(a float64) (float64, float64) {
	return CircularVelocity(a), GravityAtAltitude(a)
}

func Energy(a float64) float64 {
	return -j2.Mu / (2 * a)
}

func PotentialEnergy(a float64) float64 {
	return -j2.Mu / a
}

func EccentricAnomalyApprox(e float64, times []float64) []float64 {
	out := make([]float64, len(times))
	for i, t := range times {
		m := 2 * math.Pi * t
		out[i] = m + e*math.Sin(m)
	}
	return out
}

func TrueAnomalyApprox(e float64, eccentric []float64) []float64 {
	out := make([]float64, len(eccentric))
	for i, eValue := range eccentric {
		out[i] = eValue + 2*e*math.Sin(eValue)
	}
	return out
}

func DensityRegime(a float64) string {
	return Regime(a)
}

func AltitudeKm(a float64) float64 {
	return Altitude(a) / 1000
}

func SemimajorKm(a float64) float64 {
	return a / 1000
}

func Display(a float64) string {
	return Summary(a)
}
