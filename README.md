# j2-nodal

j2-nodal is a Go J2 secular nodal precession calculator. It computes the
secular RAAN rate and argument-of-perigee rate from semimajor axis,
eccentricity, and inclination, and solves the sun-synchronous inclination that
makes the RAAN drift equal 360 degrees per year. WGS84/EGM constants are fixed
as `mu = 3.986004418e14`, `Re = 6378137 m`, and `J2 = 1.08262668e-3`. The
service is available through HTTP JSON endpoints and CLI subcommands with no
web page.

## Usage

Run the HTTP server:

```bash
go run . serve -addr :8080
```

Evaluate from the command line:

```bash
go run . precess -a 6978137 -e 0.001 -i 97.8
go run . sso -a 6978137 -e 0.001
```

Run the SSO example:

```bash
go run . example -file example/sso-600km.json
```

The 600 km near-circular orbit has a sun-synchronous inclination near 97.8
degrees.

## HTTP API

```text
POST /api/precess  {"a":6978137,"e":0.001,"i":97.8}
POST /api/sso-i    {"a":6978137,"e":0.001}
GET  /health
```

Invalid semimajor axes, eccentricities outside `[0,1)`, or inclinations outside
`[0,180]` return an error body with HTTP 400. Orbits that cannot be
sun-synchronous return `possible: false`.

## Formulas

- `n = sqrt(mu/a^3)`
- `RAAN_dot = -1.5*n*J2*(Re/a)^2*cos(i)/(1-e^2)^2`
- `arg_peri_dot = 0.75*n*J2*(Re/a)^2*(5*cos(i)^2-1)/(1-e^2)^2`

## Code Layout

```text
internal/j2       rates, SSO inversion, validation, scenarios
internal/epoch    secular epochs, node periods, drift
internal/earth    altitude, period, regime helpers
internal/consts   units and altitude statistics
internal/table    rates tables and CSV
internal/analysis scans and rate statistics
internal/verify   polar/SSO/critical checks
internal/server   HTTP handlers and JSON responses
internal/cli      subcommand parsing and terminal output
example/          offline scenario JSON files
```

## Build and Test

```bash
export GOTOOLCHAIN=local CGO_ENABLED=0
go build ./...
go test ./...
```

The Dockerfile builds the server binary and starts it on port 8080.
