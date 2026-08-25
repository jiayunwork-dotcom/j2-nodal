package verify

import "testing"

func TestRunAllChecks(t *testing.T) {
	checks := RunAll()
	if !AllPass(checks) {
		for _, check := range checks {
			if !check.OK {
				t.Logf("%s: %s", check.Name, check.Message)
			}
		}
		t.Fatal("not all checks pass")
	}
}

func TestCheckPolarZero(t *testing.T) {
	if !CheckPolarZero().OK {
		t.Fatal("polar check failed")
	}
}

func TestCheckSemimajorTrend(t *testing.T) {
	if !CheckSemimajorTrend().OK {
		t.Fatal("semimajor check failed")
	}
}

func TestCheckEccentricityTrend(t *testing.T) {
	if !CheckEccentricityTrend().OK {
		t.Fatal("eccentricity check failed")
	}
}

func TestCheckSSOAbove90(t *testing.T) {
	if !CheckSSOAbove90().OK {
		t.Fatal("SSO check failed")
	}
}

func TestCheckCriticalNotSSO(t *testing.T) {
	if !CheckCriticalNotSSO().OK {
		t.Fatal("critical check failed")
	}
}

func TestCheckSSOTargetRate(t *testing.T) {
	if !CheckSSOTargetRate().OK {
		t.Fatal("SSO rate check failed")
	}
}

func TestCheckArgPeriCritical(t *testing.T) {
	if !CheckArgPeriCritical().OK {
		t.Fatal("arg peri check failed")
	}
}

func TestCheckConstants(t *testing.T) {
	if !CheckConstants().OK {
		t.Fatal("constants check failed")
	}
}

func TestCheckExample(t *testing.T) {
	if !CheckExample().OK {
		t.Fatal("example check failed")
	}
}
