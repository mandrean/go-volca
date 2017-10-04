package syro

import (
	"testing"
)

func TestSinValue(t *testing.T) {
	sin := Sin(0, false)
	want := int16(0)

	if sin != want {
		t.Errorf("Sinus value was incorrect, got: %d, want: %d.", sin, want)
	}
}

func TestSinValueBData(t *testing.T) {
	sin := Sin(1, true)
	want := int16(29956)

	if sin != want {
		t.Errorf("Sinus value was incorrect, got: %d, want: %d.", sin, want)
	}
}
