package practise

import (
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	if rec.Code != 200 {
		t.Fatalf("want status 200, got %d", rec.Code)
	}
	if rec.Body.String() != "Hello World\n" {
		t.Fatalf("want %q, got %q", "Hello World\n", rec.Body.String())
	}
}
