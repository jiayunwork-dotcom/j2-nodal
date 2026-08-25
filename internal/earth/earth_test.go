package earth

import (
	"testing"

	"j2-nodal/internal/j2"
)

func TestAltitude(t *testing.T) {
	if Altitude(j2.EarthRadius+600000) != 600000 {
		t.Fatal("altitude wrong")
	}
}

func TestPeriod(t *testing.T) {
	if Period(j2.EarthRadius+700000) <= 0 {
		t.Fatal("period invalid")
	}
}

func TestRegime(t *testing.T) {
	if Regime(j2.EarthRadius+600000) != "LEO" {
		t.Fatal("regime wrong")
	}
}

func TestGravity(t *testing.T) {
	if SurfaceGravity() <= 0 {
		t.Fatal("surface gravity invalid")
	}
}

func TestSummary(t *testing.T) {
	if Summary(j2.EarthRadius+600000) == "" {
		t.Fatal("empty summary")
	}
}

func TestNodePeriod(t *testing.T) {
	period, err := NodePeriod(j2.EarthRadius+600000, 0.001)
	if err != nil {
		t.Fatal(err)
	}
	if period <= 0 {
		t.Fatalf("period=%g", period)
	}
}

func TestVelocities(t *testing.T) {
	velocity, gravity := Velocities(j2.EarthRadius + 600000)
	if velocity <= 0 || gravity <= 0 {
		t.Fatalf("velocity=%g gravity=%g", velocity, gravity)
	}
}
