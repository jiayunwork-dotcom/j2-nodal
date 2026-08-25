package verify

import (
	"fmt"
	"math"

	"j2-nodal/internal/j2"
)

type Check struct {
	Name    string `json:"name"`
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func CheckPolarZero() Check {
	rate, err := j2.RAANOnly(j2.EarthRadius+700000, 0, 90)
	if err != nil {
		return Check{Name: "polar-zero", OK: false, Message: err.Error()}
	}
	ok := math.Abs(rate) < 1e-9
	return Check{Name: "polar-zero", OK: ok, Message: fmt.Sprintf("RAAN_dot=%.6g", rate)}
}

func CheckSemimajorTrend() Check {
	low, _ := j2.RAANOnly(j2.EarthRadius+500000, 0, 30)
	high, _ := j2.RAANOnly(j2.EarthRadius+1000000, 0, 30)
	ok := math.Abs(high) < math.Abs(low)
	return Check{Name: "semimajor-trend", OK: ok, Message: fmt.Sprintf("%.6g -> %.6g", low, high)}
}

func CheckEccentricityTrend() Check {
	low, _ := j2.RAANOnly(j2.EarthRadius+700000, 0, 30)
	high, _ := j2.RAANOnly(j2.EarthRadius+700000, 0.2, 30)
	ok := math.Abs(high) > math.Abs(low)
	return Check{Name: "eccentricity-trend", OK: ok, Message: fmt.Sprintf("%.6g -> %.6g", low, high)}
}

func CheckSSOAbove90() Check {
	result, err := j2.SunSyncInclination(j2.EarthRadius+600000, 0.001)
	if err != nil {
		return Check{Name: "sso-angle", OK: false, Message: err.Error()}
	}
	ok := result.Possible && result.ISSO > 90 && result.ISSO < 100
	return Check{Name: "sso-angle", OK: ok, Message: fmt.Sprintf("i=%.4f", result.ISSO)}
}

func CheckCriticalNotSSO() Check {
	critical := j2.CriticalInclination()
	result, err := j2.SunSyncInclination(j2.EarthRadius+700000, 0)
	if err != nil {
		return Check{Name: "critical-not-sso", OK: false, Message: err.Error()}
	}
	ok := math.Abs(result.ISSO-critical) > 1
	return Check{Name: "critical-not-sso", OK: ok, Message: fmt.Sprintf("SSO=%.4f critical=%.4f", result.ISSO, critical)}
}

func CheckValidationRejects() Check {
	err := j2.ValidateParams(j2.EarthRadius, 0, 90)
	ok := err != nil
	return Check{Name: "validation", OK: ok, Message: fmt.Sprintf("err=%v", err)}
}

func RunAll() []Check {
	return []Check{
		CheckPolarZero(),
		CheckSemimajorTrend(),
		CheckEccentricityTrend(),
		CheckSSOAbove90(),
		CheckCriticalNotSSO(),
		CheckValidationRejects(),
	}
}

func AllPass(checks []Check) bool {
	for _, check := range checks {
		if !check.OK {
			return false
		}
	}
	return true
}

func FormatChecks(checks []Check) string {
	out := ""
	for _, check := range checks {
		state := "PASS"
		if !check.OK {
			state = "FAIL"
		}
		out += fmt.Sprintf("%-20s %s %s\n", check.Name, state, check.Message)
	}
	return out
}

func CheckSSOTargetRate() Check {
	sso, err := j2.SunSyncInclination(j2.EarthRadius+600000, 0.001)
	if err != nil {
		return Check{Name: "sso-rate", OK: false, Message: err.Error()}
	}
	result, err := j2.PrecessionRates(j2.EarthRadius+600000, 0.001, sso.ISSO)
	if err != nil {
		return Check{Name: "sso-rate", OK: false, Message: err.Error()}
	}
	target := j2.SSOTargetDegPerDay()
	ok := math.Abs(result.RAANDot-target) < 0.01
	return Check{Name: "sso-rate", OK: ok, Message: fmt.Sprintf("rate=%.6f target=%.6f", result.RAANDot, target)}
}

func CheckArgPeriCritical() Check {
	rate, err := j2.ArgPeriOnly(j2.EarthRadius+700000, 0, j2.CriticalInclination())
	if err != nil {
		return Check{Name: "arg-peri-critical", OK: false, Message: err.Error()}
	}
	ok := math.Abs(rate) < 1e-6
	return Check{Name: "arg-peri-critical", OK: ok, Message: fmt.Sprintf("omega_dot=%.6g", rate)}
}

func CheckConstants() Check {
	ok := j2.Mu > 0 && j2.EarthRadius > 0 && j2.J2 > 0
	return Check{Name: "constants", OK: ok, Message: j2.ConstantsText()}
}

func CheckExample() Check {
	scenario, err := j2.LoadScenario("../../example/sso-600km.json")
	if err != nil {
		return Check{Name: "example", OK: false, Message: err.Error()}
	}
	result, err := j2.RunScenario(scenario)
	if err != nil {
		return Check{Name: "example", OK: false, Message: err.Error()}
	}
	ok := math.Abs(result.RAANDot-j2.SSOTargetDegPerDay()) < 0.05
	return Check{Name: "example", OK: ok, Message: fmt.Sprintf("RAAN_dot=%.6f target=%.6f", result.RAANDot, j2.SSOTargetDegPerDay())}
}

func CheckRetrograde() Check {
	result, _ := j2.SunSyncInclination(j2.EarthRadius+600000, 0.001)
	ok := j2.IsRetrograde(result.ISSO)
	return Check{Name: "retrograde", OK: ok, Message: fmt.Sprintf("i=%.4f", result.ISSO)}
}
