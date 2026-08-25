package cli

import (
	"testing"

	"j2-nodal/internal/j2"
)

func TestValidateParams(t *testing.T) {
	if err := validateParams(j2.EarthRadius+600000, 0.001, 97.8); err != nil {
		t.Fatal(err)
	}
	if err := validateParams(j2.EarthRadius, 0, 90); err == nil {
		t.Fatal("accepted a=Re")
	}
}

func TestValidateSemimajor(t *testing.T) {
	if err := validateSemimajor(j2.EarthRadius + 100); err != nil {
		t.Fatal(err)
	}
	if err := validateSemimajor(j2.EarthRadius); err == nil {
		t.Fatal("accepted a=Re")
	}
}

func TestLoadExample(t *testing.T) {
	scenario, err := loadExample("../../example/sso-600km.json")
	if err != nil {
		t.Fatal(err)
	}
	if scenario.A != j2.EarthRadius+600000 {
		t.Fatalf("scenario=%+v", scenario)
	}
}

func TestExamplePaths(t *testing.T) {
	if len(ExamplePaths()) != 1 {
		t.Fatal("expected one example path")
	}
}

func TestEvaluateExample(t *testing.T) {
	scenario, _ := loadExample("../../example/sso-600km.json")
	result, err := j2.RunScenario(scenario)
	if err != nil {
		t.Fatal(err)
	}
	if result.RAANDot <= 0 {
		t.Fatalf("rate=%g", result.RAANDot)
	}
}
