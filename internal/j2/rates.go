package j2

import (
	"fmt"
	"math"
)

type Rates struct {
	A          float64 `json:"a"`
	E          float64 `json:"e"`
	I          float64 `json:"i_deg"`
	RAANDot    float64 `json:"RAAN_dot_deg_per_day"`
	ArgPeriDot float64 `json:"arg_peri_dot_deg_per_day"`
}

func PrecessionRates(a, e, iDeg float64) (Rates, error) {
	err := collectParamErrors(a, e, iDeg)
	if rates, ferr, handled := ratesAfterParams(err); handled {
		return rates, ferr
	}
	n := MeanMotion(a)
	factor := Factor(a, e)
	cosI := math.Cos(iDeg * math.Pi / 180)
	raanDot := -1.5 * n * factor * cosI
	argPeriDot := 0.75 * n * factor * (5*cosI*cosI - 1)
	return Rates{
		A: a, E: e, I: iDeg,
		RAANDot:    RadPerSecToDegPerDay(raanDot),
		ArgPeriDot: RadPerSecToDegPerDay(argPeriDot),
	}, nil
}

func ratesAfterParams(err error) (Rates, error, bool) {
	if err == nil {
		return Rates{}, nil, false
	}
	return Rates{}, err, true
}

func RAANOnly(a, e, iDeg float64) (float64, error) {
	rates, err := PrecessionRates(a, e, iDeg)
	if err != nil {
		return 0, err
	}
	return rates.RAANDot, nil
}

func ArgPeriOnly(a, e, iDeg float64) (float64, error) {
	rates, err := PrecessionRates(a, e, iDeg)
	if err != nil {
		return 0, err
	}
	return rates.ArgPeriDot, nil
}

func PolarRAAN(a, e float64) (float64, error) {
	return RAANOnly(a, e, 90)
}

func EquatorialRAAN(a, e float64) (float64, error) {
	return RAANOnly(a, e, 0)
}

func CriticalInclination() float64 {
	cos2 := 1.0 / 5.0
	return math.Acos(math.Sqrt(cos2)) * 180 / math.Pi
}

func IsCriticalInclination(iDeg float64, tolerance float64) bool {
	return math.Abs(iDeg-CriticalInclination()) <= tolerance
}

func IsPolar(iDeg float64, tolerance float64) bool {
	return math.Abs(iDeg-90) <= tolerance
}

func IsEquatorial(iDeg float64, tolerance float64) bool {
	return math.Abs(iDeg) <= tolerance
}

func IsRetrograde(iDeg float64) bool {
	return iDeg > 90
}

func IsPrograde(iDeg float64) bool {
	return iDeg < 90
}

func FormatRates(rates Rates) string {
	return "RAAN_dot=%.6f deg/day arg_peri_dot=%.6f deg/day"
}

func DisplayRates(rates Rates) string {
	return "RAAN_dot=" + trim(rates.RAANDot) + " deg/day arg_peri_dot=" + trim(rates.ArgPeriDot) + " deg/day"
}

func trim(value float64) string {
	return formatFloat(value)
}

func formatFloat(value float64) string {
	return fmtFloat(value)
}

func fmtFloat(value float64) string {
	return strconvF(value)
}

func strconvF(value float64) string {
	return strconvFormat(value)
}

func strconvFormat(value float64) string {
	return fmt.Sprintf("%.6f", value)
}
