package models

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	// test 30 year old
	d := time.Now().AddDate(-30, 0, 0)
	result := CalculateAge(d)
	if result != 30 {
		t.Errorf("expected 30 but got %d", result)
	}

	// test newborn
	d2 := time.Now()
	result2 := CalculateAge(d2)
	if result2 != 0 {
		t.Errorf("expected 0 but got %d", result2)
	}

	// test 1 year
	d3 := time.Now().AddDate(-1, 0, 0)
	result3 := CalculateAge(d3)
	if result3 != 1 {
		t.Errorf("expected 1 but got %d", result3)
	}
}
