package calculator_test

import (
	"golang-unittesting/calculator"
	"testing"
)

func TestDivide(t *testing.T) {

	expected := 2.0

	got := calculator.Divide(10.0, 5.0)

	if got != expected {

		t.Errorf("expected %.1f, got %.1f", expected, got)
	}
}
