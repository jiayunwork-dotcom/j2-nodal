package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"j2-nodal/internal/j2"
)

func postJSON(t *testing.T, path string, body interface{}) *httptest.ResponseRecorder {
	t.Helper()
	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewReader(data))
	rec := httptest.NewRecorder()
	Routes().ServeHTTP(rec, req)
	return rec
}

func TestPrecessEndpoint(t *testing.T) {
	rec := postJSON(t, "/api/precess", map[string]interface{}{
		"a": j2.EarthRadius + 600000, "e": 0.001, "i": 97.8,
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
}

func TestSSOEndpoint(t *testing.T) {
	rec := postJSON(t, "/api/sso-i", map[string]interface{}{
		"a": j2.EarthRadius + 600000, "e": 0.001,
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
}

func TestInvalidReturns400(t *testing.T) {
	rec := postJSON(t, "/api/precess", map[string]interface{}{
		"a": j2.EarthRadius, "e": 0, "i": 90,
	})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status=%d", rec.Code)
	}
}

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	Routes().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d", rec.Code)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/precess", nil)
	rec := httptest.NewRecorder()
	Routes().ServeHTTP(rec, req)
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status=%d", rec.Code)
	}
}
