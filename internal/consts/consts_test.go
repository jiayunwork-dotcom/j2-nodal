package consts

import (
	"testing"

	"j2-nodal/internal/j2"
)

func TestSiderealDay(t *testing.T) {
	if SiderealDayHours() <= 23 {
		t.Fatal("sidereal day invalid")
	}
}

func TestMeanMotion(t *testing.T) {
	if MeanMotionDegPerDay(j2.EarthRadius+700000) <= 0 {
		t.Fatal("mean motion invalid")
	}
}

func TestJ2Factor(t *testing.T) {
	if J2Factor(j2.EarthRadius+700000, 0) <= 0 {
		t.Fatal("J2 factor invalid")
	}
}

func TestClamps(t *testing.T) {
	if ClampInclination(-5) != 0 || ClampEccentricity(2) >= 1 {
		t.Fatal("clamp failed")
	}
}

func TestAltitudeHelpers(t *testing.T) {
	if AltitudeOf(j2.EarthRadius+600000) != 600000 {
		t.Fatal("altitude wrong")
	}
	if RadiusAtAltitude(600000) != j2.EarthRadius+600000 {
		t.Fatal("radius wrong")
	}
}

func TestSortAltitudes(t *testing.T) {
	sorted := SortAltitudes([]float64{3, 1, 2})
	if sorted[0] != 1 || sorted[2] != 3 {
		t.Fatalf("sorted=%v", sorted)
	}
}

func TestMedian(t *testing.T) {
	if MedianAltitude([]float64{1, 2, 3}) != 2 {
		t.Fatal("median wrong")
	}
}

func TestSummary(t *testing.T) {
	if Summary(j2.EarthRadius+600000, 0.001, 97.8) == "" {
		t.Fatal("empty summary")
	}
}
