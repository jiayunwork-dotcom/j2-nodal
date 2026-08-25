package j2

import (
	"fmt"
	"math"
)

type SSOResult struct {
	A        float64 `json:"a"`
	E        float64 `json:"e"`
	ISSO     float64 `json:"i_sso_deg"`
	CosI     float64 `json:"cos_i"`
	Possible bool    `json:"possible"`
}

func SunSyncInclination(a, e float64) (SSOResult, error) {
	if err := ValidateSemimajor(a); err != nil {
		return SSOResult{}, err
	}
	if err := ValidateEccentricity(e); err != nil {
		return SSOResult{}, err
	}
	n := MeanMotion(a)
	factor := Factor(a, e)
	target := SSOTargetRadPerSecSidereal()
	cosI := -target / (1.5 * n * factor)
	result := SSOResult{A: a, E: e, CosI: cosI, Possible: math.Abs(cosI) <= 1}
	if result.Possible {
		result.ISSO = math.Acos(cosI) * 180 / math.Pi
	} else {
		result.ISSO = 0
	}
	return result, nil
}

func IsSunSync(iDeg float64, tolerance float64) bool {
	return iDeg > 90 && math.Abs(iDeg-97.0) < 20
}

func SSOTargetDegPerDayValue() float64 {
	return SSOTargetDegPerDay()
}

func NoSSOError() error {
	return fmt.Errorf("no_sso: |cos i| exceeds 1")
}

func SSOForAltitude(altitude float64) (SSOResult, error) {
	return SunSyncInclination(EarthRadius+altitude, 0)
}

func IsSSO(result SSOResult) bool {
	return result.Possible && result.ISSO > 90
}

func DifferenceToTarget(rate float64) float64 {
	return rate - SSOTargetDegPerDay()
}

func SSOTargetRadPerSecValue() float64 {
	return SSOTargetRadPerSec()
}

func SSOTargetRadPerDay() float64 {
	return SSOTargetRadPerSec() * SecondsPerDay
}

func SSOInclinationText(result SSOResult) string {
	if !result.Possible {
		return "no_sso"
	}
	return fmt.Sprintf("%.6f deg", result.ISSO)
}

func SSOCosI(result SSOResult) float64 {
	return result.CosI
}

func SSOAltitudeFor(a float64) float64 {
	return a - EarthRadius
}

func SSOClassify(result SSOResult) string {
	if !result.Possible {
		return "no_sso"
	}
	if result.ISSO > 90 {
		return "retrograde sun-synchronous"
	}
	return "prograde sun-synchronous"
}
