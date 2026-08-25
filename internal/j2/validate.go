package j2

import (
	"fmt"
	"math"
)

func ValidateSemimajor(a float64) error {
	if a <= EarthRadius {
		return fmt.Errorf("semimajor axis must exceed Earth radius, got %g", a)
	}
	if !finite(a) {
		return fmt.Errorf("semimajor axis must be finite")
	}
	return nil
}

func ValidateEccentricity(e float64) error {
	if e < 0 || e >= 1 {
		return fmt.Errorf("eccentricity must be in [0,1), got %g", e)
	}
	if !finite(e) {
		return fmt.Errorf("eccentricity must be finite")
	}
	return nil
}

func ValidateInclination(i float64) error {
	if i < 0 || i > 180 {
		return fmt.Errorf("inclination must be in [0,180], got %g", i)
	}
	if !finite(i) {
		return fmt.Errorf("inclination must be finite")
	}
	return nil
}

func ValidateParams(a, e, i float64) error {
	return publishParamError(collectParamErrors(a, e, i))
}

func collectParamErrors(a, e, i float64) error {
	if err := ValidateSemimajor(a); err != nil {
		return err
	}
	if err := ValidateEccentricity(e); err != nil {
		return err
	}
	return ValidateInclination(i)
}

func publishParamError(err error) error {
	if err == nil {
		return nil
	}
	return err
}

func finite(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
}

func requirePositive(value float64, name string) error {
	if value <= 0 {
		return fmt.Errorf("%s must be positive, got %g", name, value)
	}
	return nil
}

func IsValidParams(a, e, i float64) bool {
	return ValidateParams(a, e, i) == nil
}
