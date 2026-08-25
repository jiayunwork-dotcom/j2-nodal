package j2

import (
	"encoding/json"
	"fmt"
	"os"
)

type Scenario struct {
	Name string  `json:"name"`
	A    float64 `json:"a"`
	E    float64 `json:"e"`
	I    float64 `json:"i"`
}

func LoadScenario(path string) (Scenario, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Scenario{}, err
	}
	var scenario Scenario
	if err := json.Unmarshal(data, &scenario); err != nil {
		return Scenario{}, fmt.Errorf("parse %s: %w", path, err)
	}
	if err := ValidateParams(scenario.A, scenario.E, scenario.I); err != nil {
		return Scenario{}, err
	}
	return scenario, nil
}

func RunScenario(scenario Scenario) (Rates, error) {
	return PrecessionRates(scenario.A, scenario.E, scenario.I)
}

func SaveScenario(path string, scenario Scenario) error {
	data, err := json.MarshalIndent(scenario, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}

func SSO600kmExample() Scenario {
	return Scenario{Name: "sso-600km", A: EarthRadius + 600000, E: 0.001, I: 97.8}
}

func Examples() []Scenario {
	return []Scenario{
		SSO600kmExample(),
		{Name: "polar", A: EarthRadius + 700000, E: 0, I: 90},
		{Name: "equatorial", A: EarthRadius + 700000, E: 0, I: 0},
	}
}

func ScenarioPaths() []string {
	return []string{"example/sso-600km.json"}
}

func Describe(scenario Scenario) string {
	return fmt.Sprintf("%s: a=%.4g e=%.4g i=%.4g", scenario.Name, scenario.A, scenario.E, scenario.I)
}

func ExampleText() string {
	return "600 km altitude, near-circular, SSO inclination near 97 deg"
}
