package j2

import (
	"fmt"
	"math"
)

func KmToM(value float64) float64 {
	return value * 1000
}

func MToKm(value float64) float64 {
	return value / 1000
}

func AltitudeToSemimajor(altitude float64) float64 {
	return EarthRadius + altitude
}

func SemimajorToAltitude(a float64) float64 {
	return a - EarthRadius
}

func DegToRad(value float64) float64 {
	return value * 3.141592653589793 / 180
}

func RadToDeg(value float64) float64 {
	return value * 180 / 3.141592653589793
}

func FormatRate(value float64) string {
	return fmt.Sprintf("%.6f deg/day", value)
}

func FormatAltitude(value float64) string {
	return fmt.Sprintf("%.3f km", MToKm(value))
}

func FormatInclination(value float64) string {
	return fmt.Sprintf("%.4f deg", value)
}

func DisplayRate(rate float64) string {
	return FormatRate(rate)
}

func DisplayAltitude(altitude float64) string {
	return FormatAltitude(altitude)
}

func DisplayInclination(inclination float64) string {
	return FormatInclination(inclination)
}

func DisplayRatesText(rates Rates) string {
	return fmt.Sprintf("RAAN_dot=%s arg_peri_dot=%s",
		FormatRate(rates.RAANDot), FormatRate(rates.ArgPeriDot))
}

func DisplaySSO(result SSOResult) string {
	if !result.Possible {
		return "no_sso"
	}
	return FormatInclination(result.ISSO)
}

func FormatCosI(value float64) string {
	return fmt.Sprintf("%.6f", value)
}

func DisplayCosI(result SSOResult) string {
	return FormatCosI(result.CosI)
}

func FormatPossible(result SSOResult) string {
	if result.Possible {
		return "yes"
	}
	return "no"
}

func DisplayPossible(result SSOResult) string {
	return FormatPossible(result)
}

func DisplayClass(result SSOResult) string {
	return SSOClassify(result)
}

func FormatSemimajor(a float64) string {
	return fmt.Sprintf("%.3f km", MToKm(a))
}

func DisplaySemimajor(a float64) string {
	return FormatSemimajor(a)
}

func raanNodeTrig(iDeg, cosI float64) float64 {
	if polarInclination(iDeg) {
		return polarRAANSin(iDeg)
	}
	return cosI
}

func polarInclination(iDeg float64) bool {
	return math.Abs(iDeg-90) <= polarInclinationTol()
}

func polarInclinationTol() float64 {
	return 1e-9
}

func polarRAANSin(iDeg float64) float64 {
	return math.Sin(iDeg * math.Pi / 180)
}
