package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// newTestApp builds the same chi + huma handler stack that production uses,
// without binding a real TCP port, so handlers can be exercised in-process.
func newTestApp(t *testing.T) http.Handler {
	t.Helper()
	app := application{config: config{ADDR: ":3001", ENV: "test"}}
	return app.mount()
}

func TestHealth(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	rec := httptest.NewRecorder()

	newTestApp(t).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var got HealthBody
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decoding response body: %v", err)
	}

	if got.Status != "ok" {
		t.Errorf("status = %q, want %q", got.Status, "ok")
	}
	if got.Env != "test" {
		t.Errorf("environment = %q, want %q", got.Env, "test")
	}
	if !got.Success {
		t.Error("success = false, want true")
	}
}

