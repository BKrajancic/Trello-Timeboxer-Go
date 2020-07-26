package trellotimeboxergo

import "testing"

func TestHello(t *testing.T) {
	if got := Run(); got != nil {
		t.Errorf("Hello() = %q, want %q", got, "nil")
	}
}
