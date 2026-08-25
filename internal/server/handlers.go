package server

import (
	"net/http"

	"j2-nodal/internal/j2"
)

type paramsRequest struct {
	A float64 `json:"a"`
	E float64 `json:"e"`
	I float64 `json:"i"`
}

type ssoRequest struct {
	A float64 `json:"a"`
	E float64 `json:"e"`
}

func precessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w, "POST")
		return
	}
	var req paramsRequest
	if !readJSON(w, r, &req) {
		return
	}
	result, err := j2.PrecessionRates(req.A, req.E, req.I)
	if err != nil {
		badRequest(w, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func ssoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w, "POST")
		return
	}
	var req ssoRequest
	if !readJSON(w, r, &req) {
		return
	}
	result, err := j2.SunSyncInclination(req.A, req.E)
	if err != nil {
		badRequest(w, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w, "GET")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"name": "j2-nodal", "version": "1.0.0"})
}
