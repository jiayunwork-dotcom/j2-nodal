package cli

import (
	"fmt"
	"os"

	"j2-nodal/internal/j2"
)

func Run(args []string) int {
	if len(args) == 0 {
		return runServe([]string{})
	}
	switch args[0] {
	case "precess":
		return runPrecess(args[1:])
	case "sso":
		return runSSO(args[1:])
	case "example":
		return runExample(args[1:])
	case "serve":
		return runServe(args[1:])
	case "help", "-h", "--help":
		printHelp()
		return 0
	case "version":
		fmt.Println("j2-nodal 1.0.0")
		return 0
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", args[0])
		printHelp()
		return 2
	}
}

func printHelp() {
	fmt.Println(`j2-nodal: J2 secular orbital node and argument precession calculator

Usage:
  j2-nodal                            start HTTP server on :8080
  j2-nodal precess -a 6978137 -e 0.001 -i 97.8
  j2-nodal sso -a 6978137 -e 0.001
  j2-nodal example -file example/sso-600km.json
  j2-nodal serve -addr :8080

HTTP:
  POST /api/precess  {"a":6978137,"e":0.001,"i":97.8}
  POST /api/sso-i    {"a":6978137,"e":0.001}`)
}

func fail(err error) int {
	fmt.Fprintln(os.Stderr, "error:", err)
	return 1
}

func runPrecess(args []string) int {
	fs := flagSet("precess")
	a := fs.Float64("a", j2.EarthRadius+600000, "semimajor axis")
	e := fs.Float64("e", 0.001, "eccentricity")
	i := fs.Float64("i", 97.8, "inclination deg")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	result, err := j2.PrecessionRates(*a, *e, *i)
	if err != nil {
		return fail(err)
	}
	return printJSON(result)
}

func runSSO(args []string) int {
	fs := flagSet("sso")
	a := fs.Float64("a", j2.EarthRadius+600000, "semimajor axis")
	e := fs.Float64("e", 0.001, "eccentricity")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	result, err := j2.SunSyncInclination(*a, *e)
	if err != nil {
		return fail(err)
	}
	return printJSON(result)
}

func runExample(args []string) int {
	fs := flagSet("example")
	file := fs.String("file", "example/sso-600km.json", "scenario JSON file")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	scenario, err := j2.LoadScenario(*file)
	if err != nil {
		return fail(err)
	}
	result, err := j2.RunScenario(scenario)
	if err != nil {
		return fail(err)
	}
	return printJSON(result)
}

func loadExample(path string) (j2.Scenario, error) {
	return j2.LoadScenario(path)
}

func ExamplePaths() []string {
	return j2.ScenarioPaths()
}
