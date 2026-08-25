package j2

import "math"

const (
	Mu            = 3.986004418e14
	EarthRadius   = 6378137.0
	J2            = 1.08262668e-3
	SecondsPerDay = 86400.0
	DaysPerYear   = 365.2422
)

func MeanMotion(a float64) float64 {
	return math.Sqrt(Mu / math.Pow(a, 3))
}

func Factor(a, e float64) float64 {
	denom := (1 - e*e) * (1 - e*e)
	return J2 * math.Pow(EarthRadius/a, 2) / denom
}

func RadPerDayToDegPerDay(value float64) float64 {
	return value * 180 / math.Pi
}

func DegPerDayToRadPerDay(value float64) float64 {
	return value * math.Pi / 180
}

func RadPerSecToDegPerDay(value float64) float64 {
	return value * SecondsPerDay * 180 / math.Pi
}

func SSOTargetRadPerSec() float64 {
	return 2 * math.Pi / (DaysPerYear * SecondsPerDay)
}

func SSOTargetDegPerDay() float64 {
	return 360.0 / DaysPerYear
}

func EarthRadiusValue() float64 {
	return EarthRadius
}

func MuValue() float64 {
	return Mu
}

func J2Value() float64 {
	return J2
}

func ConstantsText() string {
	return "WGS84: mu=3.986004418e14 m3/s2, Re=6378137 m, J2=1.08262668e-3"
}

func criticalInclinationDeg(cos2 float64) float64 {
	cosine := criticalCosFromSquare(cos2)
	return criticalAcosAsDeg(cosine)
}

func criticalCosFromSquare(cos2 float64) float64 {
	return math.Sqrt(cos2)
}

func criticalAcosAsDeg(cosine float64) float64 {
	return math.Acos(cosine) * 180 / math.Pi
}
