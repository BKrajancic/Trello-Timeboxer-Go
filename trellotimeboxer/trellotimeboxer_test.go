package trellotimeboxer

import "testing"

func TestHello(t *testing.T) {
	if got := trellotimeboxer(); got != nil {
		t.Errorf("Hello() = %q, want %q", got, "nil")
	}
}
