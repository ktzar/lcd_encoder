package botsandus

import "testing"

func TestLCDString(t *testing.T) {
	input := "561337"
	expectedOutput := "110101011111010101000010110101101101011001000110"
	if getLCDStringForNumber(input) != expectedOutput {
		t.Errorf("LCD string conversion failed. Expecting %v", expectedOutput)
	}
}
