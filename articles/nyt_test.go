package articles

import (
	"testing"
)

func TestConstructURL(t *testing.T) {
	base := "https://example.com"
	actual := constructURL(base)
	expected := "https://example.com?sources=the-new-york-times&apiKey="
	if (actual != expected) {
		t.Fatal("expected " + actual + " to be " + expected)
	}
}
