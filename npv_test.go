package irr

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestTotalNpvCalc(t *testing.T) {
	result := TotalNpvCalc([]float64{50000, 50000}, .05)
	if result != float64(92970.52) {
		t.Errorf("TotalNpvCalc FAIL. Expected %f, got %f\n", 92970.52, result)
	} else {
		t.Logf("TotalNpvCalc PASS. Expected %f, got %f\n", 92970.52, result)
	}
}
func TestTotalNpvCalcEmpty(t *testing.T) {
	result := TotalNpvCalc([]float64{}, .05)
	if result != float64(0) {
		t.Errorf("TotalNpvCalc FAIL. Expected %v, got %f\n", 0, result)
	} else {
		t.Logf("TotalNpvCalc PASS. Expected %v, got %f\n", 0, result)
	}
}
func TestTotalNpvCalc0(t *testing.T) {
	result := TotalNpvCalc([]float64{50000, 50000}, 0)
	if result != float64(100000) {
		t.Errorf("TotalNpvCalc FAIL. Expected %v, got %f\n", 100000, result)
	} else {
		t.Logf("TotalNpvCalc PASS. Expected %v, got %f\n", 100000, result)
	}
}
func TestTotalNpvCalcNegative(t *testing.T) {
	result := TotalNpvCalc([]float64{50000, 50000}, -.05)
	if result != float64(108033.24) {
		t.Errorf("TotalNpvCalc FAIL. Expected %v, got %f\n", 108033.24, result)
	} else {
		t.Logf("TotalNpvCalc PASS. Expected %v, got %f\n", 108033.24, result)
	}
}
