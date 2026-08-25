package epoch

import (
	"math"
	"testing"

	"j2-nodal/internal/j2"
)

func TestSSOResidual(t *testing.T) {
	residual, err := SSOResidual(j2.EarthRadius+600000, 0.001)
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(residual) > 1e-6 {
		t.Fatalf("residual=%g", residual)
	}
}

func TestOrbitalPeriodDays(t *testing.T) {
	if OrbitalPeriodDays(j2.EarthRadius+700000) <= 0 {
		t.Fatal("period invalid")
	}
}

func TestRAANAfterDays(t *testing.T) {
	value, err := RAANAfterDays(j2.EarthRadius+700000, 0, 30, 10)
	if err != nil {
		t.Fatal(err)
	}
	if value == 0 {
		t.Fatal("RAAN shift zero")
	}
}

func TestNodalPeriodDays(t *testing.T) {
	period, err := NodalPeriodDays(j2.EarthRadius+600000, 0.001)
	if err != nil {
		t.Fatal(err)
	}
	if period <= 0 {
		t.Fatalf("period=%g", period)
	}
}

func TestSummary(t *testing.T) {
	if Summary(j2.EarthRadius+700000, 0) == "" {
		t.Fatal("empty summary")
	}
}

func TestIsSSOOrbit(t *testing.T) {
	if !IsSSOOrbit(j2.EarthRadius+600000, 0.001) {
		t.Fatal("expected SSO orbit")
	}
}

func TestScanSemimajor(t *testing.T) {
	rates, err := ScanSemimajor(0.001, 97.8, []float64{j2.EarthRadius + 500000, j2.EarthRadius + 700000})
	if err != nil {
		t.Fatal(err)
	}
	if len(rates) != 2 {
		t.Fatalf("len=%d", len(rates))
	}
}

func TestSummaryRates(t *testing.T) {
	rates, _ := ScanSemimajor(0, 97.8, []float64{j2.EarthRadius + 500000, j2.EarthRadius + 700000})
	if SummaryRates(rates) == "" {
		t.Fatal("empty summary")
	}
}
