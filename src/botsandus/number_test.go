package botsandus

import "testing"

func TestChecksum(t *testing.T) {
	checksum := calculateChecksum("1337")
	if checksum != 56 {
		t.Error("Incorrect checksum", checksum)
	}

	checksum = calculateChecksum("2674")
	if checksum != 9 {
		t.Error("Incorrect checksum", checksum)
	}
}
