package practise

import (
	"net/http/httptest"
	"testing"
)

func TestUserHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/user", nil)
	req.Header.Set("User-Agent", "GoTest/1.0")
	rec := httptest.NewRecorder()

	userHandler(rec, req)

	if rec.Code != 200 {
		t.Fatalf("want status 200, got %d", rec.Code)
	}
	if rec.Body.String() != "request received\n" {
		t.Fatalf("want %q, got %q", "request received\n", rec.Body.String())
	}
}
