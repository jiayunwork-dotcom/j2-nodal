package analysis

import (
	"testing"

	"j2-nodal/internal/j2"
)

func TestScanSemimajor(t *testing.T) {
	rates, err := ScanSemimajor(0, 97.8, []float64{j2.EarthRadius + 500000, j2.EarthRadius + 700000})
	if err != nil {
		t.Fatal(err)
	}
	if len(rates) != 2 {
		t.Fatalf("len=%d", len(rates))
	}
}

func TestScanInclination(t *testing.T) {
	rates, err := ScanInclination(j2.EarthRadius+700000, 0, []float64{0, 90, 180})
	if err != nil {
		t.Fatal(err)
	}
	if len(rates) != 3 {
		t.Fatalf("len=%d", len(rates))
	}
}

func TestMaxRAAN(t *testing.T) {
	rates, _ := ScanInclination(j2.EarthRadius+700000, 0, []float64{0, 90, 180})
	if MaxRAAN(rates) <= 0 {
		t.Fatal("max invalid")
	}
}

func TestSummaryAndText(t *testing.T) {
	rates, _ := ScanInclination(j2.EarthRadius+700000, 0, []float64{0, 90})
	if Summary(rates) == "" || Text(rates) == "" {
		t.Fatal("empty output")
	}
}

func TestFirstSSOIndex(t *testing.T) {
	rates, _ := ScanSemimajor(0.001, 97.8, []float64{j2.EarthRadius + 500000, j2.EarthRadius + 600000})
	if FirstSSOIndex(rates) < 0 {
		t.Fatal("no SSO index")
	}
}

func TestIsSSORate(t *testing.T) {
	if !IsSSORate(j2.SSOTargetDegPerDay(), 1e-9) {
		t.Fatal("SSO rate check failed")
	}
}

func TestSummaryStats(t *testing.T) {
	rates, _ := ScanInclination(j2.EarthRadius+700000, 0, []float64{0, 90, 180})
	if SummaryStats(rates) == "" {
		t.Fatal("empty stats")
	}
}
