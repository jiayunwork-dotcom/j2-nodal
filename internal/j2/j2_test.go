package j2

import (
	"math"
	"testing"
)

func TestPolarRAANZero(t *testing.T) {
	rate, err := RAANOnly(EarthRadius+700000, 0, 90)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(rate) > 1e-9 {
		t.Fatalf("rate=%g", rate)
	}
}

func TestSemimajorTrend(t *testing.T) {
	low, _ := RAANOnly(EarthRadius+500000, 0, 30)
	high, _ := RAANOnly(EarthRadius+1000000, 0, 30)
	if math.Abs(high) >= math.Abs(low) {
		t.Fatalf("low=%g high=%g", low, high)
	}
}

func TestEccentricityTrend(t *testing.T) {
	low, _ := RAANOnly(EarthRadius+700000, 0, 30)
	high, _ := RAANOnly(EarthRadius+700000, 0.2, 30)
	if math.Abs(high) <= math.Abs(low) {
		t.Fatalf("low=%g high=%g", low, high)
	}
}

func TestSSOAbove90(t *testing.T) {
	result, err := SunSyncInclination(EarthRadius+600000, 0.001)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Possible || result.ISSO <= 90 || result.ISSO >= 100 {
		t.Fatalf("result=%+v", result)
	}
}

func TestSSORateMatchesTarget(t *testing.T) {
	sso, _ := SunSyncInclination(EarthRadius+600000, 0.001)
	rates, _ := PrecessionRates(EarthRadius+600000, 0.001, sso.ISSO)
	if math.Abs(rates.RAANDot-SSOTargetDegPerDay()) > 0.01 {
		t.Fatalf("rate=%g target=%g", rates.RAANDot, SSOTargetDegPerDay())
	}
}

func TestCriticalInclination(t *testing.T) {
	rate, err := ArgPeriOnly(EarthRadius+700000, 0, CriticalInclination())
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(rate) > 1e-6 {
		t.Fatalf("rate=%g", rate)
	}
}

func TestNoSSO(t *testing.T) {
	result, err := SunSyncInclination(EarthRadius+35786000, 0)
	if err != nil {
		t.Fatal(err)
	}
	if result.Possible {
		t.Fatalf("expected no_sso for low orbit, result=%+v", result)
	}
}

func TestValidationRejects(t *testing.T) {
	if err := ValidateParams(EarthRadius, 0, 90); err == nil {
		t.Fatal("accepted a=Re")
	}
	if err := ValidateParams(EarthRadius+1e6, 1, 90); err == nil {
		t.Fatal("accepted e=1")
	}
}

func TestExample(t *testing.T) {
	scenario, err := LoadScenario("../../example/sso-600km.json")
	if err != nil {
		t.Fatal(err)
	}
	result, err := RunScenario(scenario)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(result.RAANDot-SSOTargetDegPerDay()) > 0.05 {
		t.Fatalf("rate=%g", result.RAANDot)
	}
}
